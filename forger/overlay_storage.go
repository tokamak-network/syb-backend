package forger

import (
	"context"
	"sync"

	mt "github.com/iden3/go-merkletree-sql/v2"
)

// OverlayStorage is a copy-on-write wrapper around a base mt.Storage.
//
// Reads:
//   - GetRoot: first time, pulls root from base and caches it;
//              afterwards, returns the cached overlay root.
//   - Get:     returns overlay node if present, otherwise falls back to base.
//
// Writes:
//   - SetRoot: only updates the overlay's cached root; base root is untouched.
//   - Put:     only records the node in the overlay map; base storage is untouched.
//
// This lets you build a speculative mt.MerkleTree for forging without
// mutating the canonical tree stored in the base storage.
type OverlayStorage struct {
	base mt.Storage

	mu       sync.RWMutex
	nodes    map[string]*mt.Node // keyed by string(keyBytes)
	root     *mt.Hash
	rootInit bool
}

// NewOverlayStorage creates a new OverlayStorage over the given base Storage.
func NewOverlayStorage(base mt.Storage) *OverlayStorage {
	return &OverlayStorage{
		base:  base,
		nodes: make(map[string]*mt.Node),
	}
}

// keyToStr converts a node key (byte slice) into a map key.
func keyToStr(key []byte) string {
	// Using string is fine here; keys are fixed-length hash bytes.
	return string(key)
}

// GetRoot returns the overlay root if set; otherwise it lazily fetches the
// base root, caches it, and returns that. It never writes back to the base.
func (o *OverlayStorage) GetRoot(ctx context.Context) (*mt.Hash, error) {
	o.mu.RLock()
	if o.rootInit {
		defer o.mu.RUnlock()
		return o.root, nil
	}
	o.mu.RUnlock()

	// First call: consult base.
	baseRoot, err := o.base.GetRoot(ctx)
	if err != nil {
		return nil, err
	}

	o.mu.Lock()
	defer o.mu.Unlock()
	if !o.rootInit {
		o.root = baseRoot
		o.rootInit = true
	}
	return o.root, nil
}

// SetRoot updates only the overlay's notion of the root. It does not call
// base.SetRoot, so canonical storage is untouched.
func (o *OverlayStorage) SetRoot(ctx context.Context, h *mt.Hash) error {
	_ = ctx // currently unused, kept for interface symmetry

	o.mu.Lock()
	defer o.mu.Unlock()
	o.root = h
	o.rootInit = true
	return nil
}

// Get returns a node from the overlay if present, otherwise falls back to base.
func (o *OverlayStorage) Get(ctx context.Context, key []byte) (*mt.Node, error) {
	k := keyToStr(key)

	o.mu.RLock()
	n, ok := o.nodes[k]
	o.mu.RUnlock()
	if ok {
		return n, nil
	}

	// Not in overlay; delegate to base.
	return o.base.Get(ctx, key)
}

// Put records the node only in the overlay map. The base storage is never
// modified, so you can discard the overlay safely after forging.
func (o *OverlayStorage) Put(ctx context.Context, key []byte, n *mt.Node) error {
	_ = ctx // currently unused

	k := keyToStr(key)

	o.mu.Lock()
	defer o.mu.Unlock()
	o.nodes[k] = n
	return nil
}
