package scheduler

import (
	"context"

	"log/slog"
)

type Scheduler struct {
}

func New(ctx context.Context, log *slog.Logger) Scheduler {
	return Scheduler{}
}

func (s Scheduler) Start() error {
	return nil
}
