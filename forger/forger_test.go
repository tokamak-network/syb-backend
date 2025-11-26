package forger

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
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
