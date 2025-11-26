package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/common"
	//_ "github.com/lib/pq" // or your DB driver

	//mt "github.com/you/syb-backend/merkletree"
	//"github.com/you/syb-backend/graphstore"
	//"github.com/you/syb-backend/protocol"
	//"github.com/you/syb-backend/syncer"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	// Open DB (example with Postgres DSN in env).
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed to open DB: %v", err)
	}
	defer db.Close()

	// TODO: init merkletree.Storage for graph & score.
	// This depends on your Storage implementation; here we assume you have
	// something like NewSQLStorage(db, "graph_tree") etc.
	var graphStorage mt.Storage
	var scoreStorage mt.Storage

	// GraphStore backed by edges table.
	gs := graphstore.NewSQLStore(db)

	protoCfg := protocol.Config{
		NumLevels: 24,
		NumLeaves: 1 << 24,
		MaxDegree: 30,
	}

	cfg := syncer.Config{
		RPCURL:          os.Getenv("RPC_URL"),
		ContractAddress: common.HexToAddress(os.Getenv("CONTRACT_ADDR")),

		StartBlock:    0,    // set appropriately
		FinalityDepth: 12,   // example
		ChunkSize:     5000, // example
		StateFile:     "sync_state.json",

		ProtoConfig: protoCfg,

		GraphStorage: graphStorage,
		ScoreStorage: scoreStorage,
		GraphStore:   gs,

		BatchSubmittedTopic: common.HexToHash(os.Getenv("BATCH_TOPIC")),
	}

	s, err := syncer.New(ctx, cfg)
	if err != nil {
		log.Fatalf("failed to create syncer: %v", err)
	}

	if err := s.Run(ctx); err != nil && err != context.Canceled {
		log.Fatalf("syncer exited with error: %v", err)
	}
}
