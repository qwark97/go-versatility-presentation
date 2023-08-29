package scheduler

import (
	"context"

	"github.com/qwark97/go-versatility-presentation/hub/logger"
)

type Scheduler struct {
}

func New(ctx context.Context, log logger.Logger) Scheduler {
	return Scheduler{}
}

func (s Scheduler) Start() error {
	return nil
}
