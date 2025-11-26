package forger

import (
	"context"
	"reflect"
	"testing"

	"syb-backend/graphstore"
)

// TestOverlayGraphStore_BasicUnion checks that:
//   - overlay sees base edges
//   - overlay adds its own edges in-memory
//   - overlay.Neighbors returns the union of base + overlay edges (deduped, sorted)
//   - base store is NOT mutated by overlay.AddEdge.
func TestOverlayGraphStore_BasicUnion(t *testing.T) {
	ctx := context.Background()

	// Base graph: 1--2, 4--5
	base := graphstore.NewMemoryStore()
	if err := base.AddEdge(ctx, 1, 2); err != nil {
		t.Fatalf("base.AddEdge(1,2): %v", err)
	}
	if err := base.AddEdge(ctx, 4, 5); err != nil {
		t.Fatalf("base.AddEdge(4,5): %v", err)
	}

	overlay := NewOverlayGraphStore(base)

	// Before adding speculative edges, overlay should see base edges.
	n1, err := overlay.Neighbors(ctx, 1)
	if err != nil {
		t.Fatalf("overlay.Neighbors(1): %v", err)
	}
	if !reflect.DeepEqual(n1, []uint64{2}) {
		t.Fatalf("overlay.Neighbors(1) = %v, want [2]", n1)
	}

	// Add speculative edges ONLY to overlay:
	// 1--3 and 3--4 (forming 1--2,1--3,3--4,4--5 in the overlay view).
	if err := overlay.AddEdge(ctx, 1, 3); err != nil {
		t.Fatalf("overlay.AddEdge(1,3): %v", err)
	}
	if err := overlay.AddEdge(ctx, 3, 4); err != nil {
		t.Fatalf("overlay.AddEdge(3,4): %v", err)
	}

	// Overlay view:
	//   1: {2,3}
	//   2: {1}
	//   3: {1,4}
	//   4: {3,5}
	//   5: {4}
	type want struct {
		v    uint64
		nbrs []uint64
	}
	for _, tc := range []want{
		{1, []uint64{2, 3}},
		{2, []uint64{1}},
		{3, []uint64{1, 4}},
		{4, []uint64{3, 5}},
		{5, []uint64{4}},
	} {
		got, err := overlay.Neighbors(ctx, tc.v)
		if err != nil {
			t.Fatalf("overlay.Neighbors(%d): %v", tc.v, err)
		}
		if !reflect.DeepEqual(got, tc.nbrs) {
			t.Fatalf("overlay.Neighbors(%d) = %v, want %v", tc.v, got, tc.nbrs)
		}
	}

	// Base must NOT see speculative edges.
	base1, err := base.Neighbors(ctx, 1)
	if err != nil {
		t.Fatalf("base.Neighbors(1): %v", err)
	}
	base3, err := base.Neighbors(ctx, 3)
	if err != nil {
		t.Fatalf("base.Neighbors(3): %v", err)
	}

	if !reflect.DeepEqual(base1, []uint64{2}) {
		t.Fatalf("base.Neighbors(1) = %v, want [2]", base1)
	}
	if base3 != nil && len(base3) != 0 {
		t.Fatalf("base.Neighbors(3) = %v, want []", base3)
	}
}

// TestOverlayGraphStore_DedupAndSelfEdge checks:
//   - adding the same edge twice doesn't create duplicate neighbors
//   - self-edges are ignored (per the usual semantics).
func TestOverlayGraphStore_DedupAndSelfEdge(t *testing.T) {
	ctx := context.Background()

	base := graphstore.NewMemoryStore()
	overlay := NewOverlayGraphStore(base)

	// Add edge 1--2 twice.
	if err := overlay.AddEdge(ctx, 1, 2); err != nil {
		t.Fatalf("overlay.AddEdge(1,2): %v", err)
	}
	if err := overlay.AddEdge(ctx, 1, 2); err != nil {
		t.Fatalf("overlay.AddEdge(1,2) again: %v", err)
	}

	// Add a self-edge; should be a no-op.
	if err := overlay.AddEdge(ctx, 5, 5); err != nil {
		t.Fatalf("overlay.AddEdge(5,5): %v", err)
	}

	n1, err := overlay.Neighbors(ctx, 1)
	if err != nil {
		t.Fatalf("overlay.Neighbors(1): %v", err)
	}
	n2, err := overlay.Neighbors(ctx, 2)
	if err != nil {
		t.Fatalf("overlay.Neighbors(2): %v", err)
	}
	n5, err := overlay.Neighbors(ctx, 5)
	if err != nil {
		t.Fatalf("overlay.Neighbors(5): %v", err)
	}

	if !reflect.DeepEqual(n1, []uint64{2}) {
		t.Fatalf("overlay.Neighbors(1) = %v, want [2]", n1)
	}
	if !reflect.DeepEqual(n2, []uint64{1}) {
		t.Fatalf("overlay.Neighbors(2) = %v, want [1]", n2)
	}
	if len(n5) != 0 {
		t.Fatalf("overlay.Neighbors(5) = %v, want []", n5)
	}
}
