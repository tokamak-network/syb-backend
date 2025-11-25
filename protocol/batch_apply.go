package protocol

import (
	"context"
	"fmt"
	"math/big"
)

// ApplyBatch applies a decoded Batch to the off-chain state:
//   - For each edge, update GraphStore + GraphTree.
//   - (ScoreTree updates are left as a TODO hook for your scoreUpdate logic.)
//
// This is called by the syncer (following on-chain batches).
func (s *State) ApplyBatch(ctx context.Context, b *Batch) error {
	if b == nil {
		return fmt.Errorf("protocol: ApplyBatch: nil Batch")
	}

	for _, e := range b.Edges {
		if err := s.ApplyEdge(ctx, e.Ilo, e.Ihi); err != nil {
			return err
		}
	}

	// TODO (optional): once you're confident in the off-chain logic, you can
	// compare s.Graph.Root() and s.Score.Root() against b.NewGraphRoot /
	// b.NewScoreRoot (converted to big.Int) and log/alert if they disagree.

	return nil
}

// ApplyEdge applies a single undirected edge {u,v} to the state:
//
//   - record the edge in GraphStore
//   - recompute NbrHash_G(u) and NbrHash_G(v)
//   - update the corresponding leaves in the GraphTree
//
// Forger can call this repeatedly on an overlay State to compute
// speculative newGraphRoot / newScoreRoot before submitting a batch.
func (s *State) ApplyEdge(ctx context.Context, u, v uint64) error {
	return s.applyGraphEdge(ctx, u, v)
}

// applyGraphEdge is the internal implementation used by both ApplyBatch and
// ApplyEdge. It only touches the *graph* state; ScoreTree updates will be
// wired in later once scoreUpdate(...) is plugged in.
func (s *State) applyGraphEdge(ctx context.Context, u, v uint64) error {
	if s.gs == nil {
		return fmt.Errorf("protocol: State.gs (GraphStore) is nil")
	}
	if err := s.gs.AddEdge(ctx, u, v); err != nil {
		return fmt.Errorf("AddEdge(%d,%d): %w", u, v, err)
	}
	if err := s.updateGraphLeaf(ctx, u); err != nil {
		return fmt.Errorf("updateGraphLeaf(%d): %w", u, err)
	}
	if err := s.updateGraphLeaf(ctx, v); err != nil {
		return fmt.Errorf("updateGraphLeaf(%d): %w", v, err)
	}
	return nil
}

// updateGraphLeaf recomputes NbrHash_G(v) from the current GraphStore
// and updates the corresponding leaf in the GraphTree.
func (s *State) updateGraphLeaf(ctx context.Context, v uint64) error {
	if s.gs == nil {
		return fmt.Errorf("protocol: State.gs (GraphStore) is nil")
	}

	// Get sorted neighbors of v from the GraphStore.
	nbrs, err := s.gs.Neighbors(ctx, v)
	if err != nil {
		return err
	}

	// Build neighbor data [deg, u0, u1, ...] and hash it using the Poseidon-based
	// NbrHash algorithm defined in hash.go.
	nbrData := buildNbrDataCompact(nbrs)
	val := nbrArrayHasher(nbrData) // *big.Int

	key := new(big.Int).SetUint64(v) // or bigFromUint64(v) if you prefer

	// We dense-initialized the graph tree over [0..NumLeaves-1], so Update
	// should always find the key. If it ever returns ErrKeyNotFound, that's
	// a protocol bug and we just surface it.
	if _, err := s.Graph.Update(ctx, key, val); err != nil {
		return err
	}
	return nil
}

