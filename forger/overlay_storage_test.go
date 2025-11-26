package forger

import (
	"context"
	"math/big"
	"testing"

	mt "github.com/iden3/go-merkletree-sql/v2"
	mtmem "github.com/iden3/go-merkletree-sql/v2/db/memory"
)

// helper to compare two *mt.Hash by bytes.
func hashesEqual(a, b *mt.Hash) bool {
	if a == nil || b == nil {
		return false
	}
	if len(a[:]) != len(b[:]) {
		return false
	}
	for i := range a[:] {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestOverlayStorage_InitialRootMatchesBase(t *testing.T) {
	ctx := context.Background()

	const depth = 4

	// Base storage + tree.
	baseStorage := mtmem.NewMemoryStorage()
	baseTree, err := mt.NewMerkleTree(ctx, baseStorage, depth)
	if err != nil {
		t.Fatalf("NewMerkleTree(base): %v", err)
	}

	// Add one leaf to base tree so the root is non-zero.
	if err := baseTree.Add(ctx, big.NewInt(1), big.NewInt(2)); err != nil {
		t.Fatalf("baseTree.Add: %v", err)
	}
	baseRoot := baseTree.Root()

	// Overlay storage on top of the same base storage.
	overlayStorage := NewOverlayStorage(baseStorage)

	// New MerkleTree using the overlay storage.
	overlayTree, err := mt.NewMerkleTree(ctx, overlayStorage, depth)
	if err != nil {
		t.Fatalf("NewMerkleTree(overlay): %v", err)
	}
	overlayRoot := overlayTree.Root()

	if !hashesEqual(baseRoot, overlayRoot) {
		t.Fatalf("overlayRoot != baseRoot:\n base    = %s\n overlay = %s",
			baseRoot.BigInt().Text(16), overlayRoot.BigInt().Text(16))
	}
}

func TestOverlayStorage_MutationsDontAffectBase(t *testing.T) {
	ctx := context.Background()

	const depth = 4

	// Base storage + tree, with an initial leaf.
	baseStorage := mtmem.NewMemoryStorage()
	baseTree, err := mt.NewMerkleTree(ctx, baseStorage, depth)
	if err != nil {
		t.Fatalf("NewMerkleTree(base): %v", err)
	}
	if err := baseTree.Add(ctx, big.NewInt(1), big.NewInt(2)); err != nil {
		t.Fatalf("baseTree.Add: %v", err)
	}
	baseRootBefore := baseTree.Root()

	// Overlay storage & tree starting from baseRootBefore.
	overlayStorage := NewOverlayStorage(baseStorage)
	overlayTree, err := mt.NewMerkleTree(ctx, overlayStorage, depth)
	if err != nil {
		t.Fatalf("NewMerkleTree(overlay): %v", err)
	}

	// Add a leaf via the overlay tree only.
	if err := overlayTree.Add(ctx, big.NewInt(3), big.NewInt(5)); err != nil {
		t.Fatalf("overlayTree.Add: %v", err)
	}
	overlayRoot := overlayTree.Root()

	// Base root must remain unchanged.
	baseRootAfter := baseTree.Root()
	if !hashesEqual(baseRootBefore, baseRootAfter) {
		t.Fatalf("base root changed after overlay mutation:\n before = %s\n after  = %s",
			baseRootBefore.BigInt().Text(16), baseRootAfter.BigInt().Text(16))
	}

	// Overlay root should differ from the original base root (we added a leaf).
	if hashesEqual(baseRootBefore, overlayRoot) {
		t.Fatalf("overlay root did not change after Add:\n base    = %s\n overlay = %s",
			baseRootBefore.BigInt().Text(16), overlayRoot.BigInt().Text(16))
	}
}
