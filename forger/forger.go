// forger/forger.go
package forger

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	mt "github.com/tokamak-network/syb-backend/merkletree"
	"github.com/tokamak-network/syb-backend/protocol"
)

// Config holds configuration for the forger daemon.
type Config struct {
	RPCURL          string
	ContractAddress common.Address

	ProtoConfig protocol.Config

	GraphStorage mt.Storage // Merkle storage for GraphTree
	ScoreStorage mt.Storage // Merkle storage for ScoreTree

	// BaseGraphStore is the canonical graph backend (usually SQLStore).
	// forger.New will wrap this in a protocol.OverlayStore so that
	// speculative edges live only in-memory.
	BaseGraphStore protocol.GraphStore

	ForgerBatchSize uint64        // when to forge
	PollInterval    time.Duration // how often to check unforged queue
}

// Forger coordinates:
//   - protocol.State (graph & score trees)
//   - RPC to read unforged queue and send submitBatch txs
//   - zk proof generation (TODO).
type Forger struct {
	cfg   Config
	cli   *ethclient.Client
	state *protocol.State // State built on top of OverlayStore
}

// New constructs a Forger and loads protocol.State using an OverlayStore:
//   - BaseGraphStore (e.g. SQLStore) is wrapped in protocol.NewOverlayStore
//   - Overlay is an in-memory MemoryStore for speculative edges
func New(ctx context.Context, cfg Config) (*Forger, error) {
	if cfg.GraphStorage == nil || cfg.ScoreStorage == nil {
		return nil, errors.New("GraphStorage and ScoreStorage must be provided")
	}
	if cfg.BaseGraphStore == nil {
		return nil, errors.New("BaseGraphStore must be provided")
	}
	if cfg.PollInterval == 0 {
		cfg.PollInterval = 15 * time.Second
	}

	cli, err := ethclient.DialContext(ctx, cfg.RPCURL)
	if err != nil {
		return nil, err
	}

	// Wrap the base graph store with an overlay for speculation.
	baseGS := cfg.BaseGraphStore
	overlayGS := protocol.NewOverlayStore(baseGS)

	// Build protocol.State using the overlay GraphStore.
	state, err := protocol.NewState(
		ctx,
		cfg.GraphStorage,
		cfg.ScoreStorage,
		overlayGS,
		cfg.ProtoConfig,
	)
	if err != nil {
		return nil, err
	}

	return &Forger{
		cfg:   cfg,
		cli:   cli,
		state: state,
	}, nil
}

// Run is the main loop of the forger.
// TODO: implement unforged queue scanning, batch construction, proofs, submitBatch.
func (f *Forger) Run(ctx context.Context) error {
	log.Println("[forger] starting forger loop (stub)")

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// TODO:
		//   - read pending unforged edges from contract (nextEdgeId, lastForgedId)
		//   - if >= ForgerBatchSize, build a candidate batch:
		//       * use f.state (which uses OverlayStore) to apply edges
		//         speculatively via protocol.ApplyBatch-like helpers
		//       * build zk inputs, generate proof
		//       * send submitBatch tx
		//
		// For now just sleep.
		time.Sleep(f.cfg.PollInterval)
	}
}
