package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/ethereum/go-ethereum/common"
	mt "github.com/iden3/go-merkletree-sql/v2"
	mtmem "github.com/iden3/go-merkletree-sql/v2/db/memory"
	_ "github.com/lib/pq" // Postgres driver

	"syb-backend/graphstore"
	"syb-backend/protocol"
	"syb-backend/syncer"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	// --- Env / config ---

	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		log.Fatal("RPC_URL env var must be set")
	}

	contractAddrHex := os.Getenv("CONTRACT_ADDR")
	if contractAddrHex == "" {
		log.Fatal("CONTRACT_ADDR env var must be set")
	}
	contractAddr := common.HexToAddress(contractAddrHex)

	batchTopicHex := os.Getenv("BATCH_TOPIC")
	if batchTopicHex == "" {
		log.Fatal("BATCH_TOPIC env var must be set (keccak256 of BatchSubmitted signature)")
	}
	batchTopic := common.HexToHash(batchTopicHex)

	startBlockStr := os.Getenv("START_BLOCK")
	if startBlockStr == "" {
		log.Fatal("START_BLOCK env var must be set (uint64)")
	}
	startBlock, err := strconv.ParseUint(startBlockStr, 10, 64)
	if err != nil || startBlock == 0 {
		log.Fatalf("invalid START_BLOCK: %q (must be >0 uint64)", startBlockStr)
	}

	// --- DB for GraphStore (edges table) ---

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL env var must be set for Postgres")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to open DB: %v", err)
	}
	defer db.Close()

	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("failed to ping DB: %v", err)
	}

	// GraphStore backed by edges(ilo, ihi) table.
	gs := graphstore.NewSQLStore(db)

	// --- Merkle storage (GraphTree + ScoreTree) ---

	// For now, use in-memory merkletree storage. In prod you can swap these
	// to SQL-backed storages (e.g. go-merkletree-sql/db/pgx).
	var graphStorage mt.Storage = mtmem.NewMemoryStorage()
	var scoreStorage mt.Storage = mtmem.NewMemoryStorage()

	protoCfg := protocol.Config{
		NumLevels: 24,      // depth D
		NumLeaves: 1 << 24, // N = 2^D (your logical "max nodes")
		MaxDegree: 30,      // maxDeg from your spec
	}

	cfg := syncer.Config{
		RPCURL:          rpcURL,
		ContractAddress: contractAddr,

		StartBlock:    startBlock,
		FinalityDepth: 12,   // example
		ChunkSize:     5000, // example
		StateFile:     "sync_state.json",

		ProtoConfig: protoCfg,

		GraphStorage: graphStorage,
		ScoreStorage: scoreStorage,
		GraphStore:   gs,

		BatchSubmittedTopic: batchTopic,
	}

	s, err := syncer.New(ctx, cfg)
	if err != nil {
		log.Fatalf("failed to create syncer: %v", err)
	}

	log.Println("[syncer-main] starting syncer loop")
	if err := s.Run(ctx); err != nil && err != context.Canceled {
		log.Fatalf("syncer exited with error: %v", err)
	}

	log.Println("[syncer-main] syncer exited cleanly")
}
