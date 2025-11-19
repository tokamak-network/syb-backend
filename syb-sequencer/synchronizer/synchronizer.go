package synchronizer

import (
	"context"
	"log"
)

// Synchronizer listens to ForgeBatch events from the contract
type Synchronizer struct {
	logger *log.Logger
}

func NewSynchronizer(logger *log.Logger) (*Synchronizer, error) {
	return &Synchronizer{
		logger: logger,
	}, nil
}

func (s *Synchronizer) Start(ctx context.Context) {
	s.logger.Println("Synchronizer started - listening for ForgeBatch events")
	s.watchEvents(ctx)
}

// watchEvents continuously listens for ForgeBatch events
func (s *Synchronizer) watchEvents(ctx context.Context) {
	s.logger.Println("Watching for ForgeBatch events...")
}
