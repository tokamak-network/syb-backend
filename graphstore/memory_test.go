package graphstore

import (
	"context"
	"reflect"
	"testing"
)

func TestMemoryStore_AddEdgeAndNeighbors(t *testing.T) {
	ctx := context.Background()
	s := NewMemoryStore()

	// Add a small triangle: 1-2, 2-3, 1-3.
	if err := s.AddEdge(ctx, 1, 2); err != nil {
		t.Fatalf("AddEdge(1,2) err: %v", err)
	}
	if err := s.AddEdge(ctx, 2, 3); err != nil {
		t.Fatalf("AddEdge(2,3) err: %v", err)
	}
	if err := s.AddEdge(ctx, 3, 1); err != nil {
		t.Fatalf("AddEdge(3,1) err: %v", err)
	}

	// Add duplicates + self-edge and make sure nothing weird happens.
	if err := s.AddEdge(ctx, 2, 1); err != nil { // duplicate in reverse order
		t.Fatalf("AddEdge(2,1) err: %v", err)
	}
	if err := s.AddEdge(ctx, 1, 1); err != nil { // self edge -> no-op
		t.Fatalf("AddEdge(1,1) err: %v", err)
	}

	cases := []struct {
		v    uint64
		want []uint64
	}{
		{v: 1, want: []uint64{2, 3}},
		{v: 2, want: []uint64{1, 3}},
		{v: 3, want: []uint64{1, 2}},
		{v: 4, want: nil}, // no neighbors
	}

	for _, tc := range cases {
		got, err := s.Neighbors(ctx, tc.v)
		if err != nil {
			t.Fatalf("Neighbors(%d) err: %v", tc.v, err)
		}
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("Neighbors(%d) = %v, want %v", tc.v, got, tc.want)
		}
	}
}
