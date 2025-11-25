package protocol

import "context"

// GraphStore is a minimal adjacency-graph abstraction.
//
// It keeps protocol logic decoupled from how we actually store edges
// (SQL, in-memory, etc.).
//
// - AddEdge must treat the graph as undirected: AddEdge(u,v) implies AddEdge(v,u)
//   at the logical level (either via implementation or via calling convention).
// - Neighbors(v) must return a sorted, deduplicated slice of neighbor IDs.

type GraphStore interface {
	// AddEdge inserts an undirected edge {u,v}.
	// Implementations should enforce u != v and idempotency (no duplicate neighbors).
	AddEdge(ctx context.Context, u, v uint64) error

	// Neighbors returns the sorted, deduplicated neighbor list for vertex v.
	Neighbors(ctx context.Context, v uint64) ([]uint64, error)
}

