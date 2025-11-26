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
// GraphStore is taken from s.GraphStore (must be non-nil).
func (s *State) ApplyBatch(ctx context.Context, b *Batch) error {
	if b == nil {
		return fmt.Errorf("protocol: ApplyBatch: nil Batch")
	}
	if s.GraphStore == nil {
		return fmt.Errorf("protocol: ApplyBatch: nil State.GraphStore")
	}

	for _, e := range b.Edges {
		if err := s.ApplyEdge(ctx, e.Ilo, e.Ihi); err != nil {
			return err
		}
	}

	// TODO (optional): compare s.Graph.Root() / s.Score.Root() with
	// b.NewGraphRoot / b.NewScoreRoot for sanity checks.

	return nil
}

// ApplyEdge applies a single undirected edge {u,v} to the state:
//
//   - record the edge in GraphStore
//   - recompute NbrHash_G(u) and NbrHash_G(v)
//   - update the corresponding leaves in the GraphTree
func (s *State) ApplyEdge(ctx context.Context, u, v uint64) error {
	return s.applyGraphEdge(ctx, u, v)
}

// internal implementation used by ApplyBatch / ApplyEdge.
func (s *State) applyGraphEdge(ctx context.Context, u, v uint64) error {
	if s.GraphStore == nil {
		return fmt.Errorf("protocol: applyGraphEdge: nil State.GraphStore")
	}
	if err := s.GraphStore.AddEdge(ctx, u, v); err != nil {
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
	if s.GraphStore == nil {
		return fmt.Errorf("protocol: updateGraphLeaf: nil State.GraphStore")
	}

	// Get sorted neighbors of v from the GraphStore.
	nbrs, err := s.GraphStore.Neighbors(ctx, v)
	if err != nil {
		return err
	}

	// Build compact [deg, u0, u1, ...] then pad to the fixed length
	// implied by s.maxDegree, and hash with Poseidon.
	compact := buildNbrDataCompact(nbrs)
	padded := padNbrData(compact, s.maxDegree)
	val := nbrArrayHasher(padded)

	key := new(big.Int).SetUint64(v)


	// We dense-initialized the graph tree via NewState/InitGraphTree, so Update
	// should always find the key. If it ever returns ErrKeyNotFound, that's
	// a protocol bug and we just surface it.
	if _, err := s.Graph.Update(ctx, key, val); err != nil {
		return err
	}
	return nil
}
