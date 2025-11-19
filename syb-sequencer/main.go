package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"syb-sequencer/forger"
	"syb-sequencer/synchronizer"
)

func main() {
	logger := log.New(os.Stdout, "SEQUENCER: ", log.LstdFlags|log.Lshortfile)
	logger.Println("Starting Sybil Sequencer...")

	// Create Forger
	forger, err := forger.NewForger(logger)
	if err != nil {
		logger.Fatalf("Failed to create forger: %v", err)
	}

	// Create Synchronizer
	synchronizer, err := synchronizer.NewSynchronizer(logger)
	if err != nil {
		logger.Fatalf("Failed to create synchronizer: %v", err)
	}

	// Create context that can be cancelled
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start synchronizer (listens to ForgeBatch events)
	go func() {
		synchronizer.Start(ctx)
	}()

	// Start forger (Checks when to forge a batch)
	go func() {
		forger.Start(ctx)
	}()

	// Wait for interrupt signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	logger.Println("Shutting down...")
	logger.Println("Shutdown complete.")
}
