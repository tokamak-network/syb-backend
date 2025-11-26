package forger

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	mt "github.com/iden3/go-merkletree-sql/v2"
	"syb-backend/protocol"
)

// Config holds configuration for the forger daemon.
type Config struct {
	RPCURL          string
	ContractAddress common.Address

	ProtoConfig protocol.Config

	GraphStorage mt.Storage // Merkle storage for GraphTree
	ScoreStorage mt.Storage // Merkle storage for ScoreTree

	// BaseGraphStore is the canonical graph backend (usually SQLStore).
	// forger.New will wrap this in an OverlayGraphStore so that
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
	state *protocol.State // State built on top of OverlayGraphStore
}

// New constructs a Forger and loads protocol.State using an OverlayGraphStore:
//   - BaseGraphStore (e.g. SQLStore) is wrapped in NewOverlayGraphStore
//   - Speculative edges live only in the overlay; base store is unchanged.
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
	overlayGS := NewOverlayGraphStore(baseGS)

	// Build protocol.State using the overlay GraphStore.
	// NOTE: this still uses the same underlying Merkle storages (GraphStorage
	// and ScoreStorage), so Merkle updates are canonical. A Merkle overlay
	// would be a separate step (overlay mt.Storage) if/when you want that.
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
//
// CURRENTLY: stub implementation. It just logs and sleeps until ctx is cancelled.
// Later you will:
//   - read pending unforged edges from the contract
//   - construct a *fresh* overlay (GraphStore + possibly Merkle storage)
//   - apply edges via protocol.State (ApplyEdge/ApplyBatch)
//   - generate zk proofs and call submitBatch.
func (f *Forger) Run(ctx context.Context) error {
	log.Println("[forger] starting forger loop (stub)")

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// TODO:
		//   1. Query unforged queue (nextEdgeId, lastForgedId) from contract.
		//   2. If >= f.cfg.ForgerBatchSize, build ONE speculative overlay state:
		//        - newOverlayGS := NewOverlayGraphStore(f.cfg.BaseGraphStore)
		//        - newOverlayMtStorage for graph/score if you add Merkle overlays
		//        - overlayState := protocol.NewState(...) using the overlays,
		//          but *reusing* the current roots as starting point.
		//   3. Apply edges to overlayState via ApplyEdge / ApplyBatch.
		//   4. Read overlayState.Graph.Root() / overlayState.Score.Root()
		//      as new roots for submitBatch.
		//   5. Generate proof, send tx.
		//
		// IMPORTANT: do NOT use f.state directly if you want each batch to
		// have its own fresh overlay; build a new one per batch and discard it.

		time.Sleep(f.cfg.PollInterval)
	}
}
