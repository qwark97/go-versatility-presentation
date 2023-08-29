package scheduler

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/qwark97/go-versatility-presentation/hub/logger"
	"github.com/qwark97/go-versatility-presentation/hub/peripherals/storage"
)

type Storage interface {
}

type Scheduler struct {
	ctx     context.Context
	storage Storage
	log     logger.Logger

	state map[uuid.UUID]entry
}

type entry struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
	conf       storage.Configuration
}

func New(ctx context.Context, storage Storage, log logger.Logger) Scheduler {
	return Scheduler{
		ctx:     ctx,
		storage: storage,
		log:     log,
	}
}

func (s Scheduler) Add(configuration storage.Configuration) error {
	ctx, cancel := context.WithCancel(s.ctx)
	e := entry{
		ctx:        ctx,
		cancelFunc: cancel,
		conf:       configuration,
	}
	if _, present := s.state[e.conf.ID]; present {
		return fmt.Errorf("entry with id: %s already exists", e.conf.ID)
	} else {
		s.state[e.conf.ID] = e
	}

	go s.worker(e)
	return nil
}

func (s Scheduler) Remove(id uuid.UUID) {
	entry, present := s.state[id]
	if !present {
		return
	}
	entry.cancelFunc()
	delete(s.state, id)
}

func (s Scheduler) worker(e entry) {
	d, err := time.ParseDuration(e.conf.Frequency)
	if err != nil {
		s.log.Warning("invalid frequency: %s", err)
		d = time.Second
	}
	s.log.Info("using %s frequency", d)

	ticker := time.NewTicker(d)
	for {
		select {
		case <-e.ctx.Done():
			s.log.Info("scheduler for entry with id: %s has been stopped", e.conf.ID)
			return
		case <-ticker.C:
			err := s.processData(e)
			if err != nil {
				s.log.Warning("failed to process data for entry with id: %s", e.conf.ID)
			}
		}
	}
}

func (s Scheduler) processData(e entry) error {
	s.log.Info("PUK (%s)", e.conf.ID)
	return nil
}
