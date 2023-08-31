package scheduler

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/qwark97/go-versatility-presentation/hub/logger"
	"github.com/qwark97/go-versatility-presentation/hub/peripherals/storage"
)

type Storage interface {
	SaveLastReading(id uuid.UUID, data any) error
	ReadLastReading(id uuid.UUID) (any, error)
}

type Scheduler struct {
	ctx     context.Context
	storage Storage
	log     logger.Logger

	state      map[uuid.UUID]entry
	httpClient *http.Client
}

type entry struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
	conf       storage.Configuration
}

func New(ctx context.Context, storage Storage, log logger.Logger) *Scheduler {
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	return &Scheduler{
		ctx:        ctx,
		storage:    storage,
		log:        log,
		httpClient: client,
	}
}

func (s *Scheduler) Add(configuration storage.Configuration) error {
	if s.state == nil {
		s.state = make(map[uuid.UUID]entry)
	}

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

func (s *Scheduler) Remove(id uuid.UUID) {
	entry, present := s.state[id]
	if !present {
		return
	}
	entry.cancelFunc()
	delete(s.state, id)
}

func (s *Scheduler) worker(e entry) {
	d, err := time.ParseDuration(e.conf.Frequency)
	if err != nil {
		s.log.Warning("invalid frequency: %s", err)
		d = time.Second
	}
	s.log.Info("using %s frequency", d)

	ticker := time.NewTicker(d)

	process := func() {
		ctx, cancel := context.WithTimeout(e.ctx, d)
		defer cancel()
		err := s.processData(ctx, e)
		if err != nil {
			s.log.Warning("failed to process data for entry with id %s: %v", e.conf.ID, err)
		}
	}
	process()

	for {
		select {
		case <-e.ctx.Done():
			s.log.Info("scheduler for entry with id: %s has been stopped", e.conf.ID)
			return
		case <-ticker.C:
			process()
		}
	}
}

func (s *Scheduler) processData(ctx context.Context, e entry) error {
	request, err := http.NewRequest(e.conf.Method, e.conf.Addr, nil)
	if err != nil {
		return err
	}

	response, err := s.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	var (
		data      string
		container []byte
	)
	container, err = io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if len(container) == 0 {
		return fmt.Errorf("endpoint have not returned any data")
	}

	data = fmt.Sprintf("%s%s (%s)", string(container), e.conf.Unit, e.conf.Description)
	err = s.storage.SaveLastReading(e.conf.ID, data)
	if err != nil {
		return err
	}

	return nil
}

func (s *Scheduler) LastReading(id uuid.UUID) (any, error) {
	return s.storage.ReadLastReading(id)
}
