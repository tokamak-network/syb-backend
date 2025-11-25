package protocol

import (
	"context"
	"math/big"
	"testing"

	mtmem "github.com/iden3/go-merkletree-sql/v2/db/memory"
)

// NOTE: fakeGraphStore & newFakeGraphStore are defined in state_test.go.
// We just reuse newFakeGraphStore() here.

// TestApplyEdge_UpdatesGraphStoreAndGraphTree checks that:
//   - ApplyEdge updates GraphStore adjacency
//   - ApplyEdge updates GraphTree leaves for both endpoints
func TestApplyEdge_UpdatesGraphStoreAndGraphTree(t *testing.T) {
	ctx := context.Background()

	graphStorage := mtmem.NewMemoryStorage()
	scoreStorage := mtmem.NewMemoryStorage()
	gs := newFakeGraphStore()

	cfg := Config{
		NumLevels: 8,
		NumLeaves: 256,
		MaxDegree: 30,
	}

	st, err := NewState(ctx, graphStorage, scoreStorage, gs, cfg)
	if err != nil {
		t.Fatalf("NewState: %v", err)
	}

	// Apply a single edge {1,2}.
	if err := st.ApplyEdge(ctx, 1, 2); err != nil {
		t.Fatalf("ApplyEdge(1,2): %v", err)
	}

	// Check adjacency in GraphStore.
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

	// Verify GraphTree leaf values for 1 and 2.
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

// TestApplyBatch_UsesApplyEdge ensures ApplyBatch iterates over all edges
// and has the same effects as calling ApplyEdge for each.
func TestApplyBatch_UsesApplyEdge(t *testing.T) {
	ctx := context.Background()

	graphStorage := mtmem.NewMemoryStorage()
	scoreStorage := mtmem.NewMemoryStorage()
	gs := newFakeGraphStore()

	cfg := Config{
		NumLevels: 8,
		NumLeaves: 256,
		MaxDegree: 30,
	}

	st, err := NewState(ctx, graphStorage, scoreStorage, gs, cfg)
	if err != nil {
		t.Fatalf("NewState: %v", err)
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

	// Expected adjacency:
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

	// Spot-check GraphTree for vertex 2 and 3.
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
