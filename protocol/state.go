package protocol

import (
	"context"
	"errors"
	"log"

	mt "github.com/iden3/go-merkletree-sql/v2" 
)

// Config controls Merkle layout & dense init.
type Config struct {
	NumLevels int    // Merkle tree depth (D)
	NumLeaves uint64 // how many dense leaves to pre-init (N = 2^D)
	MaxDegree uint64 // max neighbors per node (maxDeg)
}

// State wraps the two Merkle trees: GraphTree + ScoreTree,
// plus a GraphStore for the underlying graph adjacency.
type State struct {
	Graph      *mt.MerkleTree
	Score      *mt.MerkleTree
	GraphStore GraphStore

	maxDegree uint64
	padLen    int // padLen = 1 + 15 * numRounds
	numRounds int // numRounds = (padLen-1)/15
}

// NewState constructs graphTree + scoreTree on top of the given storages
// and GraphStore. If Graph is "fresh" (root == HashZero), it runs dense init.
func NewState(
	ctx context.Context,
	graphStorage, scoreStorage mt.Storage,
	gs GraphStore,
	cfg Config,
) (*State, error) {
	if graphStorage == nil || scoreStorage == nil {
		return nil, errors.New("graphStorage and scoreStorage must be non-nil")
	}
	if gs == nil {
		return nil, errors.New("GraphStore must be non-nil")
	}
	if cfg.NumLevels <= 0 {
		return nil, errors.New("NumLevels must be > 0")
	}
	if cfg.NumLeaves == 0 {
		return nil, errors.New("NumLeaves must be > 0")
	}
	if cfg.MaxDegree == 0 {
		return nil, errors.New("MaxDegree must be > 0")
	}

	scoreTree, err := mt.NewMerkleTree(ctx, scoreStorage, cfg.NumLevels)
	if err != nil {
		return nil, err
	}

	graphTree, err := mt.NewMerkleTree(ctx, graphStorage, cfg.NumLevels)
	if err != nil {
		return nil, err
	}

	padLen := padLenFromMaxDegree(cfg.MaxDegree)
	numRounds := (padLen - 1) / 15

	st := &State{
		Graph:      graphTree,
		Score:      scoreTree,
		GraphStore: gs,

		maxDegree: cfg.MaxDegree,
		padLen:    padLen,
		numRounds: numRounds,
	}

	// If graphTree is fresh, run dense init.
	if hashesEqual(graphTree.Root(), &mt.HashZero) {
		log.Printf(
			"[protocol] graphTree has zero root, running dense init (NumLeaves=%d, MaxDegree=%d, padLen=%d)\n",
			cfg.NumLeaves, cfg.MaxDegree, padLen,
		)

		if err := st.InitGraphTree(ctx, cfg.NumLeaves); err != nil {
			return nil, err
		}
		log.Printf("[protocol] graphTree dense init done. New root: %s\n",
			graphTree.Root().BigInt().Text(16))
	} else {
		log.Printf("[protocol] graphTree existing root: %s\n",
			graphTree.Root().BigInt().Text(16))
	}

	return st, nil
}

// InitGraphTree inserts NbrHash of an all-zero neighbour array at each leaf
// index up to numLeaves. Index i is simply 0..numLeaves-1 (dense semantics).
func (s *State) InitGraphTree(ctx context.Context, numLeaves uint64) error {
	// Padded, all-zero neighbour array for this maxDegree.
	zeroNbrArray := zeroArray(s.maxDegree)
	val := nbrArrayHasher(zeroNbrArray)

	for i := uint64(0); i < numLeaves; i++ {
		k := bigFromUint64(i)
		if err := s.Graph.Add(ctx, k, val); err != nil {
			if errors.Is(err, mt.ErrEntryIndexAlreadyExists) {
				// If re-running init on a partially-initialized tree, ignore duplicates.
				continue
			}
			return err
		}
	}
	return nil
}

// hashesEqual compares two mt.Hash values.
func hashesEqual(a, b *mt.Hash) bool {
	if a == nil || b == nil {
		return false
	}
	if len(a[:]) != len(b[:]) {
		return false
	}
	for i := range a[:] {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
