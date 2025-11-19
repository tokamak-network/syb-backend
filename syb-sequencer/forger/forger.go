package forger

import (
	"context"
	"log"
)

// Forger is responsible for forging batches when needed
type Forger struct {
	logger *log.Logger
}

func NewForger(logger *log.Logger) (*Forger, error) {

	return &Forger{
		logger: logger,
	}, nil
}

func (f *Forger) Start(ctx context.Context) {
	f.logger.Println("Forger started - listening for ForgeBatch events")
	f.watchEvents(ctx)
}

// watchEvents continuously listens for ForgeBatch events
func (f *Forger) watchEvents(ctx context.Context) {
	f.logger.Println("Watching for ForgeBatch events...")
	// Check if a batch needs to be forged based on batch size if full.
	// if noTxInBatch >= batchSize, forge a batch.
	// Call ForgeBatch(batchNum)
}

// ForgeBatch forges a batch of transactions
func (f *Forger) ForgeBatch(batchNum uint32) error {
	f.logger.Printf("Forging batch %d", batchNum)
	return nil
}
