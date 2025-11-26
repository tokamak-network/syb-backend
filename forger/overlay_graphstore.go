package forger

import (
	"context"
	"sort"
	"sync"

	"syb-backend/protocol"
)

// OverlayGraphStore is a copy-on-write layer over a base GraphStore.
//
// Reads:
//   - Neighbors(v) returns the union of base.Neighbors(v) and the
//     overlay's own in-memory edges, deduplicated and sorted.
//
// Writes:
//   - AddEdge(u, v) only touches the overlay; it NEVER writes to the base.
//
// This is exactly what the forger needs for speculative batch building:
// you can apply edges, compute new roots via protocol.State, and then
// throw the overlay away without mutating the canonical DB-backed graph.
type OverlayGraphStore struct {
	base protocol.GraphStore

	mu  sync.RWMutex
	adj map[uint64]map[uint64]struct{} // overlay-only adjacency
}

// NewOverlayGraphStore wraps a base GraphStore with an in-memory overlay.
func NewOverlayGraphStore(base protocol.GraphStore) *OverlayGraphStore {
	return &OverlayGraphStore{
		base: base,
		adj:  make(map[uint64]map[uint64]struct{}),
	}
}

// AddEdge records an undirected edge {u, v} in the overlay only.
// Self-edges (u == v) are treated as no-ops.
func (o *OverlayGraphStore) AddEdge(ctx context.Context, u, v uint64) error {
	_ = ctx // currently unused; kept for interface symmetry

	if u == v {
		// Match the usual semantics: ignore self edges.
		return nil
	}
	if u > v {
		u, v = v, u
	}

	o.mu.Lock()
	defer o.mu.Unlock()

	if o.adj[u] == nil {
		o.adj[u] = make(map[uint64]struct{})
	}
	if o.adj[v] == nil {
		o.adj[v] = make(map[uint64]struct{})
	}
	o.adj[u][v] = struct{}{}
	o.adj[v][u] = struct{}{}

	return nil
}

// Neighbors returns the sorted list of neighbors of v, combining:
//
//   - neighbors from the base GraphStore, and
//   - neighbors from the overlay's in-memory adjacency.
//
// Duplicates are removed; the result is sorted ascending.
func (o *OverlayGraphStore) Neighbors(ctx context.Context, v uint64) ([]uint64, error) {
	// Start with base neighbors.
	baseNbrs, err := o.base.Neighbors(ctx, v)
	if err != nil {
		return nil, err
	}

	o.mu.RLock()
	defer o.mu.RUnlock()

	// Merge base + overlay into a set.
	m := make(map[uint64]struct{}, len(baseNbrs))
	for _, n := range baseNbrs {
		m[n] = struct{}{}
	}
	if overlaySet, ok := o.adj[v]; ok {
		for n := range overlaySet {
			m[n] = struct{}{}
		}
	}

	// Convert to sorted slice.
	out := make([]uint64, 0, len(m))
	for n := range m {
		out = append(out, n)
	}
	sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })

	return out, nil
}
