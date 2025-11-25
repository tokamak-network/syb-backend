package protocol

import (
	"context"
	"testing"

	coredb "github.com/iden3/go-iden3-core/v2/db"
	mt "github.com/iden3/go-merkletree-sql/v2"
)

// fakeGraphStore is a minimal GraphStore implementation for tests.
// State.InitGraphTree doesn't actually use GraphStore, so this can be trivial.
type fakeGraphStore struct{}

func (f *fakeGraphStore) AddEdge(ctx context.Context, u, v uint64) error {
	return nil
}

func (f *fakeGraphStore) Neighbors(ctx context.Context, v uint64) ([]uint64, error) {
	return nil, nil
}

// newTestStorages creates in-memory SMT storages for graph & score trees.
func newTestStorages(t *testing.T) (mt.Storage, mt.Storage) {
	t.Helper()

	// Using iden3-core in-memory DB. Adjust if you use a different storage.
	graphStore := coredb.NewMemoryStorage()
	scoreStore := coredb.NewMemoryStorage()

	return graphStore, scoreStore
}

// TestNewState_BadConfig verifies NewState fails on obviously bad configs.
func TestNewState_BadConfig(t *testing.T) {
	ctx := context.Background()
	graphStore, scoreStore := newTestStorages(t)
	gs := &fakeGraphStore{}

	_, err := NewState(ctx, nil, scoreStore, gs, Config{
		NumLevels: 24,
		NumLeaves: 1 << 24,
		MaxDegree: 30,
	})
	if err == nil {
		t.Fatalf("expected error when graphStorage is nil, got nil")
	}

	_, err = NewState(ctx, graphStore, nil, gs, Config{
		NumLevels: 24,
		NumLeaves: 1 << 24,
		MaxDegree: 30,
	})
	if err == nil {
		t.Fatalf("expected error when scoreStorage is nil, got nil")
	}

	_, err = NewState(ctx, graphStore, scoreStore, nil, Config{
		NumLevels: 24,
		NumLeaves: 1 << 24,
		MaxDegree: 30,
	})
	if err == nil {
		t.Fatalf("expected error when GraphStore is nil, got nil")
	}

	_, err = NewState(ctx, graphStore, scoreStore, gs, Config{
		NumLevels: 0,
		NumLeaves: 1 << 24,
		MaxDegree: 30,
	})
	if err == nil {
		t.Fatalf("expected error when NumLevels <= 0, got nil")
	}

	_, err = NewState(ctx, graphStore, scoreStore, gs, Config{
		NumLevels: 24,
		NumLeaves: 0,
		MaxDegree: 30,
	})
	if err == nil {
		t.Fatalf("expected error when NumLeaves == 0, got nil")
	}

	_, err = NewState(ctx, graphStore, scoreStore, gs, Config{
		NumLevels: 24,
		NumLeaves: 1 << 24,
		MaxDegree: 0,
	})
	if err == nil {
		t.Fatalf("expected error when MaxDegree == 0, got nil")
	}
}

// TestNewState_DenseInit ensures that for a fresh graph tree (zero root),
// NewState runs InitGraphTree and changes the graph root, while score root
// remains zero.
func TestNewState_DenseInit(t *testing.T) {
	ctx := context.Background()
	graphStore, scoreStore := newTestStorages(t)
	gs := &fakeGraphStore{}

	cfg := Config{
		NumLevels: 4,      // small tree for test
		NumLeaves: 1 << 4, // 16 leaves
		MaxDegree: 30,
	}

	// Sanity: a fresh tree created directly should have zero root.
	graphTree, err := mt.NewMerkleTree(ctx, graphStore, cfg.NumLevels)
	if err != nil {
		t.Fatalf("NewMerkleTree(graph) failed: %v", err)
	}
	if !hashesEqual(graphTree.Root(), &mt.HashZero) {
		t.Fatalf("expected fresh graphTree root to be HashZero")
	}

	scoreTree, err := mt.NewMerkleTree(ctx, scoreStore, cfg.NumLevels)
	if err != nil {
		t.Fatalf("NewMerkleTree(score) failed: %v", err)
	}
	if !hashesEqual(scoreTree.Root(), &mt.HashZero) {
		t.Fatalf("expected fresh scoreTree root to be HashZero")
	}

	// Now build protocol.State on top of fresh storages. This should:
	//  - Detect zero graph root,
	//  - Run InitGraphTree (dense init),
	//  - Commit a non-zero graph root.
	state, err := NewState(ctx, graphStore, scoreStore, gs, cfg)
	if err != nil {
		t.Fatalf("NewState failed: %v", err)
	}

	if state.Graph == nil || state.Score == nil {
		t.Fatalf("state.Graph or state.Score is nil")
	}

	if hashesEqual(state.Graph.Root(), &mt.HashZero) {
		t.Fatalf("expected graph root to be non-zero after dense init")
	}

	// We didn't touch the score tree during NewState, so its root
	// should still be zero.
	if !hashesEqual(state.Score.Root(), &mt.HashZero) {
		t.Fatalf("expected score root to remain HashZero")
	}
}

// TestInitGraphTree_Idempotent checks that re-running InitGraphTree on the
// same State and range doesn't change the root (thanks to duplicate handling).
func TestInitGraphTree_Idempotent(t *testing.T) {
	ctx := context.Background()
	graphStore, scoreStore := newTestStorages(t)
	gs := &fakeGraphStore{}

	cfg := Config{
		NumLevels: 4,
		NumLeaves: 8,
		MaxDegree: 30,
	}

	state, err := NewState(ctx, graphStore, scoreStore, gs, cfg)
	if err != nil {
		t.Fatalf("NewState failed: %v", err)
	}

	// Capture the current graph root.
	root1 := state.Graph.Root()

	// Re-run InitGraphTree with the same numLeaves. This should:
	//  - Try to add entries that already exist,
	//  - Hit ErrEntryIndexAlreadyExists,
	//  - Ignore those and leave the tree unchanged.
	if err := state.InitGraphTree(ctx, cfg.NumLeaves); err != nil {
		t.Fatalf("InitGraphTree second run failed: %v", err)
	}

	root2 := state.Graph.Root()
	if !hashesEqual(root1, root2) {
		t.Fatalf("expected graph root to remain unchanged on second InitGraphTree run")
	}
}
