package forger

import (
	"context"
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	mt "github.com/iden3/go-merkletree-sql/v2"
	mtmem "github.com/iden3/go-merkletree-sql/v2/db/memory"

	"syb-backend/graphstore"
	"syb-backend/protocol"
)

// newTestRPCServer creates a tiny HTTP JSON-RPC server that always returns a
// valid JSON-RPC response. It's enough for ethclient.DialContext to succeed.
func newTestRPCServer(t *testing.T) *httptest.Server {
	t.Helper()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var req struct {
			ID any `json:"id"`
		}
		_ = json.NewDecoder(r.Body).Decode(&req)

		resp := map[string]any{
			"jsonrpc": "2.0",
			"id":      req.ID,
			"result":  nil,
		}
		_ = json.NewEncoder(w).Encode(resp)
	})

	return httptest.NewServer(handler)
}

func newTestConfig(t *testing.T) (Config, mt.Storage, mt.Storage, *graphstore.MemoryStore) {
	t.Helper()

	graphStorage := mtmem.NewMemoryStorage()
	scoreStorage := mtmem.NewMemoryStorage()
	baseGS := graphstore.NewMemoryStore()

	protoCfg := protocol.Config{
		NumLevels: 5,
		NumLeaves: 16,
		MaxDegree: 30,
	}

	return Config{
		ProtoConfig:    protoCfg,
		GraphStorage:   graphStorage,
		ScoreStorage:   scoreStorage,
		BaseGraphStore: baseGS,

		ForgerBatchSize: 4,
		PollInterval:    10 * time.Millisecond,
	}, graphStorage, scoreStorage, baseGS
}

func TestNew_ValidConfig(t *testing.T) {
	ctx := context.Background()
	cfg, _, _, _ := newTestConfig(t)

	// Start a test RPC server and point the Forger at it.
	srv := newTestRPCServer(t)
	defer srv.Close()
	cfg.RPCURL = srv.URL
	cfg.ContractAddress = [20]byte{} // zero address is fine for test

	f, err := New(ctx, cfg)
	if err != nil {
		t.Fatalf("New returned error: %v", err)
	}
	if f == nil {
		t.Fatalf("New returned nil Forger")
	}
	if f.state == nil {
		t.Fatalf("New: forger.state is nil")
	}
}

func TestNew_MissingDeps(t *testing.T) {
	ctx := context.Background()
	cfg, graphStorage, scoreStorage, baseGS := newTestConfig(t)

	srv := newTestRPCServer(t)
	defer srv.Close()
	cfg.RPCURL = srv.URL
	cfg.ContractAddress = [20]byte{}

	t.Run("missing GraphStorage", func(t *testing.T) {
		cfg2 := cfg
		cfg2.GraphStorage = nil
		_, err := New(ctx, cfg2)
		if err == nil {
			t.Fatalf("expected error when GraphStorage is nil, got nil")
		}
	})

	t.Run("missing ScoreStorage", func(t *testing.T) {
		cfg2 := cfg
		cfg2.GraphStorage = graphStorage
		cfg2.ScoreStorage = nil
		_, err := New(ctx, cfg2)
		if err == nil {
			t.Fatalf("expected error when ScoreStorage is nil, got nil")
		}
	})

	t.Run("missing BaseGraphStore", func(t *testing.T) {
		cfg2 := cfg
		cfg2.GraphStorage = graphStorage
		cfg2.ScoreStorage = scoreStorage
		cfg2.BaseGraphStore = nil
		_, err := New(ctx, cfg2)
		if err == nil {
			t.Fatalf("expected error when BaseGraphStore is nil, got nil")
		}
	})

	// Sanity: original config still works.
	cfg.GraphStorage = graphStorage
	cfg.ScoreStorage = scoreStorage
	cfg.BaseGraphStore = baseGS
	if _, err := New(ctx, cfg); err != nil {
		t.Fatalf("New with full deps failed unexpectedly: %v", err)
	}
}

func TestRun_StopsOnContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, _, _, _ := newTestConfig(t)
	srv := newTestRPCServer(t)
	defer srv.Close()
	cfg.RPCURL = srv.URL
	cfg.ContractAddress = [20]byte{}

	f, err := New(ctx, cfg)
	if err != nil {
		t.Fatalf("New returned error: %v", err)
	}

	errCh := make(chan error, 1)
	go func() {
		errCh <- f.Run(ctx)
	}()

	// Cancel almost immediately; Run should observe ctx.Done and return.
	cancel()

	select {
	case err := <-errCh:
		if err != context.Canceled {
			t.Fatalf("Run returned error %v, want context.Canceled", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatalf("Run did not return after context cancellation")
	}
}

func TestOverlay_EndToEnd_StateAndStorage(t *testing.T) {
	ctx := context.Background()

	// --- 1. Base state setup ---

	const depth = 5
	const numLeaves = 16

	baseGraphStorage := mtmem.NewMemoryStorage()
	baseScoreStorage := mtmem.NewMemoryStorage()
	baseGS := graphstore.NewMemoryStore()

	protoCfg := protocol.Config{
		NumLevels: depth,
		NumLeaves: numLeaves,
		MaxDegree: 8,
	}

	// Base canonical state.
	baseState, err := protocol.NewState(ctx, baseGraphStorage, baseScoreStorage, baseGS, protoCfg)
	if err != nil {
		t.Fatalf("protocol.NewState (base): %v", err)
	}

	// Apply some base edges: 1--2, 2--3
	if err := baseState.ApplyEdge(ctx, 1, 2); err != nil {
		t.Fatalf("baseState.ApplyEdge(1,2): %v", err)
	}
	if err := baseState.ApplyEdge(ctx, 2, 3); err != nil {
		t.Fatalf("baseState.ApplyEdge(2,3): %v", err)
	}

	baseGraphRootBefore := baseState.Graph.Root()
	baseScoreRootBefore := baseState.Score.Root()

	// Sanity: base neighbors
	baseNbr2, err := baseGS.Neighbors(ctx, 2)
	if err != nil {
		t.Fatalf("baseGS.Neighbors(2): %v", err)
	}
	if !reflect.DeepEqual(baseNbr2, []uint64{1, 3}) {
		t.Fatalf("base neighbors(2) = %v, want [1 3]", baseNbr2)
	}

	// --- 2. Overlay state using overlay storages + overlay graphstore ---

	overlayGraphStorage := NewOverlayStorage(baseGraphStorage)
	overlayScoreStorage := NewOverlayStorage(baseScoreStorage)
	overlayGS := NewOverlayGraphStore(baseGS)

	overlayState, err := protocol.NewState(ctx, overlayGraphStorage, overlayScoreStorage, overlayGS, protoCfg)
	if err != nil {
		t.Fatalf("protocol.NewState (overlay): %v", err)
	}

	// Initial roots must match.
	if !hashesEqual(baseGraphRootBefore, overlayState.Graph.Root()) {
		t.Fatalf("overlay graph root mismatch before edges:\n base    = %s\n overlay = %s",
			baseGraphRootBefore.BigInt().Text(16),
			overlayState.Graph.Root().BigInt().Text(16),
		)
	}
	if !hashesEqual(baseScoreRootBefore, overlayState.Score.Root()) {
		t.Fatalf("overlay score root mismatch before edges:\n base    = %s\n overlay = %s",
			baseScoreRootBefore.BigInt().Text(16),
			overlayState.Score.Root().BigInt().Text(16),
		)
	}

	// --- 3. Apply speculative edges ONLY to overlay state ---

	// Add edges 3--4 and 4--5 speculatively.
	if err := overlayState.ApplyEdge(ctx, 3, 4); err != nil {
		t.Fatalf("overlayState.ApplyEdge(3,4): %v", err)
	}
	if err := overlayState.ApplyEdge(ctx, 4, 5); err != nil {
		t.Fatalf("overlayState.ApplyEdge(4,5): %v", err)
	}

	overlayGraphRootAfter := overlayState.Graph.Root()
	overlayScoreRootAfter := overlayState.Score.Root()
	baseGraphRootAfter := baseState.Graph.Root()
	baseScoreRootAfter := baseState.Score.Root()

	// --- 4. Check GraphStore adjacency: base unchanged, overlay extended ---

	// Base neighbors of 3: only {2}
	baseNbr3, err := baseGS.Neighbors(ctx, 3)
	if err != nil {
		t.Fatalf("baseGS.Neighbors(3): %v", err)
	}
	if !reflect.DeepEqual(baseNbr3, []uint64{2}) {
		t.Fatalf("base neighbors(3) = %v, want [2]", baseNbr3)
	}

	// Overlay neighbors of 3: {2,4}
	overlayNbr3, err := overlayGS.Neighbors(ctx, 3)
	if err != nil {
		t.Fatalf("overlayGS.Neighbors(3): %v", err)
	}
	if !reflect.DeepEqual(overlayNbr3, []uint64{2, 4}) {
		t.Fatalf("overlay neighbors(3) = %v, want [2 4]", overlayNbr3)
	}

	// Base neighbors of 4: only {3} in overlay, but empty in base.
	baseNbr4, err := baseGS.Neighbors(ctx, 4)
	if err != nil {
		t.Fatalf("baseGS.Neighbors(4): %v", err)
	}
	if len(baseNbr4) != 0 {
		t.Fatalf("base neighbors(4) = %v, want []", baseNbr4)
	}
	overlayNbr4, err := overlayGS.Neighbors(ctx, 4)
	if err != nil {
		t.Fatalf("overlayGS.Neighbors(4): %v", err)
	}
	if !reflect.DeepEqual(overlayNbr4, []uint64{3, 5}) {
		t.Fatalf("overlay neighbors(4) = %v, want [3 5]", overlayNbr4)
	}

	// --- 5. Check Merkle roots: base unchanged, overlay changed ---

	if !hashesEqual(baseGraphRootBefore, baseGraphRootAfter) {
		t.Fatalf("base graph root changed unexpectedly:\n before = %s\n after  = %s",
			baseGraphRootBefore.BigInt().Text(16),
			baseGraphRootAfter.BigInt().Text(16),
		)
	}
	if hashesEqual(baseGraphRootBefore, overlayGraphRootAfter) {
		t.Fatalf("overlay graph root did not change:\n base    = %s\n overlay = %s",
			baseGraphRootBefore.BigInt().Text(16),
			overlayGraphRootAfter.BigInt().Text(16),
		)
	}

	// Score roots: depending on whether scoreUpdate is implemented,
	// they may or may not change. For now, just assert base != overlay
	// or at least that base didn't change.
	if !hashesEqual(baseScoreRootBefore, baseScoreRootAfter) {
		t.Fatalf("base score root changed unexpectedly:\n before = %s\n after  = %s",
			baseScoreRootBefore.BigInt().Text(16),
			baseScoreRootAfter.BigInt().Text(16),
		)
	}
	// It's okay if overlayScoreRootAfter == baseScoreRootBefore for now.

	// --- 6. Check leaf + internal structure differ via proof for vertex 3 ---

	idx := big.NewInt(3)

	baseProof, baseVal, err := baseState.Graph.GenerateProof(ctx, idx, baseGraphRootBefore)
	if err != nil {
		t.Fatalf("baseState.Graph.GenerateProof: %v", err)
	}
	overlayProof, overlayVal, err := overlayState.Graph.GenerateProof(ctx, idx, overlayGraphRootAfter)
	if err != nil {
		t.Fatalf("overlayState.Graph.GenerateProof: %v", err)
	}

	// Leaf value (NbrHash) for vertex 3 should differ:
	if baseVal.Cmp(overlayVal) == 0 {
		t.Fatalf("leaf value for vertex 3 identical in base and overlay, want different:\n base    = %s\n overlay = %s",
			baseVal.Text(16), overlayVal.Text(16),
		)
	}

	// Internal path (siblings) should differ as well, indicating intermediate
	// hashes have changed.
	if reflect.DeepEqual(baseProof.Siblings, overlayProof.Siblings) {
		t.Fatalf("siblings for vertex 3 proof are identical, want different")
	}
}
