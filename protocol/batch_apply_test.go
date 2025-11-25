package protocol

import (
	"context"
	"math/big"
	"sort"
	"sync"
	"testing"

	mt "github.com/iden3/go-merkletree-sql/v2"
	mtmem "github.com/iden3/go-merkletree-sql/v2/db/memory"
)

// ---- test helpers ----

// fakeGraphStore is a simple in-memory GraphStore implementation used only in tests.
type fakeGraphStore struct {
	mu  sync.RWMutex
	adj map[uint64]map[uint64]struct{}
}

func newFakeGraphStore() *fakeGraphStore {
	return &fakeGraphStore{
		adj: make(map[uint64]map[uint64]struct{}),
	}
}

func (f *fakeGraphStore) AddEdge(ctx context.Context, u, v uint64) error {
	_ = ctx
	if u == v {
		// match your production semantics: disallow self edges
		return nil
	}
	if f.adj[u] == nil {
		f.adj[u] = make(map[uint64]struct{})
	}
	if f.adj[v] == nil {
		f.adj[v] = make(map[uint64]struct{})
	}
	f.adj[u][v] = struct{}{}
	f.adj[v][u] = struct{}{}
	return nil
}

func (f *fakeGraphStore) Neighbors(ctx context.Context, v uint64) ([]uint64, error) {
	_ = ctx
	f.mu.RLock()
	defer f.mu.RUnlock()

	m := f.adj[v]
	if m == nil {
		return nil, nil
	}
	out := make([]uint64, 0, len(m))
	for n := range m {
		out = append(out, n)
	}
	sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })
	return out, nil
}

func newTestMerkleTree(t *testing.T, numLevels int) *mt.MerkleTree {
	t.Helper()
	st := mtmem.NewMemoryStorage()
	mtree, err := mt.NewMerkleTree(context.Background(), st, numLevels)
	if err != nil {
		t.Fatalf("NewMerkleTree: %v", err)
	}
	return mtree
}

// ---- tests ----

func TestApplyEdge_UpdatesGraphStoreAndGraphTree(t *testing.T) {
	ctx := context.Background()

	const (
		numLevels = 8   // enough depth; see previous tests about maxLevels vs leaves
		maxDegree = 30
		numLeaves = 256
	)

	graphMT := newTestMerkleTree(t, numLevels)
	scoreMT := newTestMerkleTree(t, numLevels)

	gs := newFakeGraphStore()

	st := &State{
		Graph: graphMT,
		Score: scoreMT,
		gs:    gs,
	}

	// Dense-init the GraphTree with zero neighbor arrays.
	if err := st.InitGraphTree(ctx, maxDegree, numLeaves); err != nil {
		t.Fatalf("InitGraphTree: %v", err)
	}

	// Apply a single edge {1,2}.
	if err := st.ApplyEdge(ctx, 1, 2); err != nil {
		t.Fatalf("ApplyEdge(1,2): %v", err)
	}

	// Check fakeGraphStore adjacency first.
	n1, err := gs.Neighbors(ctx, 1)
	if err != nil {
		t.Fatalf("gs.Neighbors(1): %v", err)
	}
	n2, err := gs.Neighbors(ctx, 2)
	if err != nil {
		t.Fatalf("gs.Neighbors(2): %v", err)
	}
	if len(n1) != 1 || n1[0] != 2 {
		t.Fatalf("neighbors(1) = %v, want [2]", n1)
	}
	if len(n2) != 1 || n2[0] != 1 {
		t.Fatalf("neighbors(2) = %v, want [1]", n2)
	}

	// Now verify that the GraphTree leaf values for 1 and 2 match the
	// expected NbrHash_G(v) computed via nbrArrayHasher(buildNbrDataCompact(...)).
	checkLeaf := func(v uint64, expectedNbrs []uint64) {
		t.Helper()
		key := new(big.Int).SetUint64(v)
		_, val, _, err := st.Graph.Get(ctx, key)
		if err != nil {
			t.Fatalf("Graph.Get(%d): %v", v, err)
		}

		nbrData := buildNbrDataCompact(expectedNbrs)
		wantHash := nbrArrayHasher(nbrData)

		if val.Cmp(wantHash) != 0 {
			t.Fatalf("Graph leaf value for %d mismatch:\n got  %s\n want %s",
				v, val.Text(16), wantHash.Text(16))
		}
	}

	checkLeaf(1, []uint64{2})
	checkLeaf(2, []uint64{1})
}

func TestApplyBatch_UsesApplyEdge(t *testing.T) {
	ctx := context.Background()

	const (
		numLevels = 8
		maxDegree = 30
		numLeaves = 256
	)

	graphMT := newTestMerkleTree(t, numLevels)
	scoreMT := newTestMerkleTree(t, numLevels)

	gs := newFakeGraphStore()

	st := &State{
		Graph: graphMT,
		Score: scoreMT,
		gs:    gs,
	}

	if err := st.InitGraphTree(ctx, maxDegree, numLeaves); err != nil {
		t.Fatalf("InitGraphTree: %v", err)
	}

	// Construct a Batch with edges: {1,2}, {2,3}.
	b := &Batch{
		BatchID: 1,
		Count:   2,
		Edges: []Edge{
			{Ilo: 1, Ihi: 2},
			{Ilo: 2, Ihi: 3},
		},
	}

	if err := st.ApplyBatch(ctx, b); err != nil {
		t.Fatalf("ApplyBatch: %v", err)
	}

	// gs adjacency expectations:
	// 1: {2}
	// 2: {1,3}
	// 3: {2}
	type wantNbrs struct {
		v    uint64
		want []uint64
	}
	for _, tc := range []wantNbrs{
		{1, []uint64{2}},
		{2, []uint64{1, 3}},
		{3, []uint64{2}},
	} {
		n, err := gs.Neighbors(ctx, tc.v)
		if err != nil {
			t.Fatalf("gs.Neighbors(%d): %v", tc.v, err)
		}
		if len(n) != len(tc.want) {
			t.Fatalf("neighbors(%d) len=%d, want %d (%v)",
				tc.v, len(n), len(tc.want), n)
		}
		for i := range n {
			if n[i] != tc.want[i] {
				t.Fatalf("neighbors(%d) = %v, want %v", tc.v, n, tc.want)
			}
		}
	}

	// Spot-check GraphTree leaves for 2 and 3.
	checkLeaf := func(v uint64, expectedNbrs []uint64) {
		t.Helper()
		key := new(big.Int).SetUint64(v)
		_, val, _, err := st.Graph.Get(ctx, key)
		if err != nil {
			t.Fatalf("Graph.Get(%d): %v", v, err)
		}

		nbrData := buildNbrDataCompact(expectedNbrs)
		wantHash := nbrArrayHasher(nbrData)

		if val.Cmp(wantHash) != 0 {
			t.Fatalf("Graph leaf value for %d mismatch:\n got  %s\n want %s",
				v, val.Text(16), wantHash.Text(16))
		}
	}

	checkLeaf(2, []uint64{1, 3})
	checkLeaf(3, []uint64{2})
}
