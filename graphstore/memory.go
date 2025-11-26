package graphstore

import (
	"context"
	"sort"
	"sync"
)

// MemoryStore is an in-memory implementation of the GraphStore interface,
// useful for tests or ephemeral simulations.
type MemoryStore struct {
	mu  sync.RWMutex
	adj map[uint64]map[uint64]struct{}
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		adj: make(map[uint64]map[uint64]struct{}),
	}
}

func (m *MemoryStore) AddEdge(ctx context.Context, u, v uint64) error {
	if u == v {
		// Treat self-edge as no-op.
		return nil
	}
	if u > v {
		u, v = v, u
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	if m.adj[u] == nil {
		m.adj[u] = make(map[uint64]struct{})
	}
	if m.adj[v] == nil {
		m.adj[v] = make(map[uint64]struct{})
	}
	m.adj[u][v] = struct{}{}
	m.adj[v][u] = struct{}{}
	return nil
}

func (m *MemoryStore) Neighbors(ctx context.Context, v uint64) ([]uint64, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	mset := m.adj[v]
	if mset == nil {
		return nil, nil
	}
	out := make([]uint64, 0, len(mset))
	for nbr := range mset {
		out = append(out, nbr)
	}
	sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })
	return out, nil
}
