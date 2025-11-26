package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/common"
	//_ "github.com/lib/pq"

	//mt "github.com/tokamak-network/syb-backend/merkletree"
	//"github.com/tokamak-network/syb-backend/graphstore"
	//"github.com/tokamak-network/syb-backend/forger"
	//"github.com/tokamak-network/syb-backend/protocol"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	// ----- DB & storage wiring -----

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed to open DB: %v", err)
	}
	defer db.Close()

	// TODO: construct concrete mt.Storage implementations using your SMT backend.
	// This depends on how you implemented merkletree.Storage for SQL.
	var graphStorage mt.Storage
	var scoreStorage mt.Storage

	// Base graph store (canonical graph in SQL).
	baseGS := graphstore.NewSQLStore(db)

	// ----- Protocol config -----

	protoCfg := protocol.Config{
		NumLevels: 24,
		NumLeaves: 1 << 24,
		MaxDegree: 30,
	}

	// ----- Forger config -----

	cfg := forger.Config{
		RPCURL:          os.Getenv("RPC_URL"),
		ContractAddress: common.HexToAddress(os.Getenv("CONTRACT_ADDR")),

		ProtoConfig: protoCfg,

		GraphStorage:   graphStorage,
		ScoreStorage:   scoreStorage,
		BaseGraphStore: baseGS,

		ForgerBatchSize: 32,
		PollInterval:    15 * time.Second,
	}

	f, err := forger.New(ctx, cfg)
	if err != nil {
		log.Fatalf("failed to create forger: %v", err)
	}

	if err := f.Run(ctx); err != nil && err != context.Canceled {
		log.Fatalf("forger exited with error: %v", err)
	}
}
