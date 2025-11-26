package syncer

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"math/big"
	"os"
	"sort"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	mt "github.com/iden3/go-merkletree-sql/v2"

	"syb-backend/protocol"
)

// Config holds runtime configuration for the synchronizer daemon.
type Config struct {
	RPCURL          string
	ContractAddress common.Address

	StartBlock    uint64
	FinalityDepth uint64 // e.g. 12
	ChunkSize     uint64 // e.g. 5000
	StateFile     string // e.g. "sync_state.json"

	ProtoConfig protocol.Config // NumLevels, NumLeaves, MaxDegree

	GraphStorage mt.Storage
	ScoreStorage mt.Storage
	GraphStore   protocol.GraphStore

	// Topic hash of the BatchSubmitted event (keccak256("BatchSubmitted(...)")).
	BatchSubmittedTopic common.Hash
}

// SyncState is stored in sync_state.json.
type SyncState struct {
	LastProcessedBlock uint64 `json:"lastProcessedBlock"`
}

// Syncer coordinates:
//   - Ethereum RPC (follow blocks / logs)
//   - Merkle state (via protocol.State)
//   - Local JSON sync state.
type Syncer struct {
	cfg   Config
	cli   *ethclient.Client
	state SyncState

	trees *protocol.State
}

// New constructs a Syncer and performs startup:
//   - dials RPC
//   - loads protocol.State (graph + score trees)
//   - initializes graphTree if needed (dense init happens inside NewState)
//   - loads or creates sync_state.json
func New(ctx context.Context, cfg Config) (*Syncer, error) {
	if cfg.GraphStorage == nil || cfg.ScoreStorage == nil {
		return nil, errors.New("GraphStorage and ScoreStorage must be provided")
	}
	if cfg.GraphStore == nil {
		return nil, errors.New("GraphStore must be provided")
	}
	if cfg.ChunkSize == 0 {
		cfg.ChunkSize = 5000
	}
	if cfg.FinalityDepth == 0 {
		cfg.FinalityDepth = 12
	}
	if cfg.StartBlock == 0 {
		return nil, errors.New("StartBlock must be > 0")
	}
	if (cfg.BatchSubmittedTopic == common.Hash{}) {
		return nil, errors.New("BatchSubmittedTopic must be set")
	}

	cli, err := ethclient.DialContext(ctx, cfg.RPCURL)
	if err != nil {
		return nil, err
	}

	trees, err := protocol.NewState(
		ctx,
		cfg.GraphStorage,
		cfg.ScoreStorage,
		cfg.GraphStore,
		cfg.ProtoConfig,
	)
	if err != nil {
		return nil, err
	}

	state, err := loadState(cfg.StateFile)
	if err != nil {
		return nil, err
	}
	if state == nil {
		// First run: set lastProcessedBlock = START_BLOCK - 1 and create file.
		state = &SyncState{
			LastProcessedBlock: cfg.StartBlock - 1,
		}
		if err := saveState(cfg.StateFile, state); err != nil {
			return nil, err
		}
		log.Printf("[syncer] no state file found; initializing lastProcessedBlock=%d\n",
			state.LastProcessedBlock)
	} else {
		log.Printf("[syncer] loaded state: lastProcessedBlock=%d\n", state.LastProcessedBlock)
	}

	return &Syncer{
		cfg:   cfg,
		cli:   cli,
		state: *state,
		trees: trees,
	}, nil
}

// Run runs the main sync loop until ctx is cancelled or a fatal error occurs.
func (s *Syncer) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		head, err := s.cli.BlockNumber(ctx)
		if err != nil {
			log.Printf("[syncer] error fetching block number: %v\n", err)
			time.Sleep(10 * time.Second)
			continue
		}

		// Don't process blocks too close to the tip (reorg-safe window).
		if head <= s.cfg.FinalityDepth {
			time.Sleep(10 * time.Second)
			continue
		}
		target := head - s.cfg.FinalityDepth

		if target <= s.state.LastProcessedBlock {
			// Already synced up to our finality depth.
			time.Sleep(10 * time.Second)
			continue;
		}

		from := s.state.LastProcessedBlock + 1
		for from <= target {
			to := from + s.cfg.ChunkSize - 1
			if to > target {
				to = target
			}

			if err := s.processRange(ctx, from, to); err != nil {
				log.Printf("[syncer] error processing [%d,%d]: %v\n", from, to, err)
				// backoff and let the outer loop retry later
				time.Sleep(10 * time.Second)
				break
			}

			// Successfully processed [from..to], update state & persist JSON.
			s.state.LastProcessedBlock = to
			if err := saveState(s.cfg.StateFile, &s.state); err != nil {
				log.Printf("[syncer] error saving state: %v\n", err)
			}

			from = to + 1
		}
	}
}

// processRange fetches and applies BatchSubmitted logs in [from, to].
func (s *Syncer) processRange(ctx context.Context, from, to uint64) error {
	log.Printf("[syncer] processing range [%d,%d]\n", from, to)

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(from)),
		ToBlock:   big.NewInt(int64(to)),
		Addresses: []common.Address{s.cfg.ContractAddress},
		Topics:    [][]common.Hash{{s.cfg.BatchSubmittedTopic}},
	}

	logs, err := s.cli.FilterLogs(ctx, query)
	if err != nil {
		return err
	}

	// Ensure canonical order (blockNumber, logIndex).
	sort.Slice(logs, func(i, j int) bool {
		if logs[i].BlockNumber == logs[j].BlockNumber {
			return logs[i].Index < logs[j].Index
		}
		return logs[i].BlockNumber < logs[j].BlockNumber
	})

	for _, lg := range logs {
		// Decode on-chain log -> protocol.Batch.
		batch, err := protocol.DecodeBatchFromLog(lg)
		if err != nil {
			return err
		}
		// Apply to local state trees.
		if err := s.trees.ApplyBatch(ctx, batch); err != nil {
			return err
		}
	}

	return nil
}

// ---- JSON state helpers ----

func loadState(path string) (*SyncState, error) {
	if path == "" {
		// No persistence requested.
		return &SyncState{LastProcessedBlock: 0}, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	var st SyncState
	if err := json.Unmarshal(data, &st); err != nil {
		return nil, err
	}
	return &st, nil
}

func saveState(path string, st *SyncState) error {
	if path == "" {
		// No persistence requested.
		return nil
	}

	tmp := path + ".tmp"
	data, err := json.MarshalIndent(st, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(tmp, data, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}
