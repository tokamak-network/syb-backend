package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/common"
	mt "github.com/iden3/go-merkletree-sql/v2"
	mtmem "github.com/iden3/go-merkletree-sql/v2/db/memory"
	_ "github.com/lib/pq" // Postgres driver

	"syb-backend/forger"
	"syb-backend/graphstore"
	"syb-backend/protocol"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	// ----- Env / config -----

	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		log.Fatal("RPC_URL env var must be set")
	}

	contractAddrHex := os.Getenv("CONTRACT_ADDR")
	if contractAddrHex == "" {
		log.Fatal("CONTRACT_ADDR env var must be set")
	}
	contractAddr := common.HexToAddress(contractAddrHex)

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL env var must be set for Postgres")
	}

	forgerBatchSize := uint64(32)
	if s := os.Getenv("FORGER_BATCH_SIZE"); s != "" {
		if v, err := strconv.ParseUint(s, 10, 64); err == nil && v > 0 {
			forgerBatchSize = v
		}
	}

	pollInterval := 15 * time.Second
	if s := os.Getenv("FORGER_POLL_MS"); s != "" {
		if ms, err := strconv.Atoi(s); err == nil && ms > 0 {
			pollInterval = time.Duration(ms) * time.Millisecond
		}
	}

	// ----- DB & GraphStore wiring -----

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to open DB: %v", err)
	}
	defer db.Close()

	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("failed to ping DB: %v", err)
	}

	// Base graph store (canonical graph in SQL via edges table).
	baseGS := graphstore.NewSQLStore(db)

	// ----- Merkle storage (GraphTree + ScoreTree) -----
	//
	// For now, we use in-memory merkletree storage. In production you will
	// likely want to use SQL-backed storages (e.g. go-merkletree-sql/v2/db/pgx)
	// so the canonical trees are persisted, and the forger’s OverlayStorage
	// wraps those.
	var graphStorage mt.Storage = mtmem.NewMemoryStorage()
	var scoreStorage mt.Storage = mtmem.NewMemoryStorage()

	// ----- Protocol config -----

	protoCfg := protocol.Config{
		NumLevels: 24,        // depth D
		NumLeaves: 1 << 24,   // N = 2^D
		MaxDegree: 30,        // maxDeg from your spec
	}

	// ----- Forger config -----

	cfg := forger.Config{
		RPCURL:          rpcURL,
		ContractAddress: contractAddr,

		ProtoConfig: protoCfg,

		GraphStorage:   graphStorage,
		ScoreStorage:   scoreStorage,
		BaseGraphStore: baseGS,

		ForgerBatchSize: forgerBatchSize,
		PollInterval:    pollInterval,
	}

	f, err := forger.New(ctx, cfg)
	if err != nil {
		log.Fatalf("failed to create forger: %v", err)
	}

	log.Println("[forger-main] starting forger loop")
	if err := f.Run(ctx); err != nil && err != context.Canceled {
		log.Fatalf("forger exited with error: %v", err)
	}
	log.Println("[forger-main] forger exited cleanly")
}
