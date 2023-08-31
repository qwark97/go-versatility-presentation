package peripherals

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/qwark97/go-versatility-presentation/hub/logger"
	"github.com/qwark97/go-versatility-presentation/hub/peripherals/storage"
	"golang.org/x/exp/slices"
)

type Scheduler interface {
	Add(configuration storage.Configuration) error
	Remove(id uuid.UUID)
	LastReading(id uuid.UUID) (any, error)
}

type Storage interface {
	AssureDir() error
	SaveConfigurations(configurations []storage.Configuration) error
	ReadConfigurations() ([]storage.Configuration, error)
}

type Peripherals struct {
	ctx       context.Context
	scheduler Scheduler
	storage   Storage
	log       logger.Logger
}

func New(ctx context.Context, scheduler Scheduler, storage Storage, log logger.Logger) Peripherals {
	return Peripherals{
		ctx:       ctx,
		scheduler: scheduler,
		storage:   storage,
		log:       log,
	}
}

func (p Peripherals) All(ctx context.Context) ([]storage.Configuration, error) {
	configurations, err := p.storage.ReadConfigurations()
	if err != nil {
		p.log.Error(fmt.Sprintf("failed to read configurations: %v", err))
		return nil, err
	}

	return configurations, nil
}

func (p Peripherals) Add(ctx context.Context, configuration storage.Configuration) error {
	err := p.storage.AssureDir()
	if err != nil {
		return err
	}

	configurations, err := p.storage.ReadConfigurations()
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
	}

	foundIdx := slices.IndexFunc(configurations, func(c storage.Configuration) bool { return c.ID == configuration.ID })
	if foundIdx >= 0 {
		return fmt.Errorf("configuration already exists")
	}

	configuration.ID = uuid.New()

	err = p.scheduler.Add(configuration)
	if err != nil {
		return err
	}

	configurations = append(configurations, configuration)

	return p.storage.SaveConfigurations(configurations)
}

func (p Peripherals) ByID(ctx context.Context, id uuid.UUID) (storage.Configuration, error) {
	var notFound storage.Configuration

	configurations, err := p.storage.ReadConfigurations()
	if err != nil {
		p.log.Error(fmt.Sprintf("failed to read configurations: %v", err))
		return notFound, err
	}

	for _, configuration := range configurations {
		if configuration.ID == id {
			return configuration, nil
		}
	}

	return notFound, nil
}
func (p Peripherals) DeleteOne(ctx context.Context, id uuid.UUID) error {
	configurations, err := p.storage.ReadConfigurations()
	if err != nil {
		p.log.Error(fmt.Sprintf("failed to read configurations: %v", err))
		return err
	}

	configurations = slices.DeleteFunc(configurations, func(c storage.Configuration) bool { return id == c.ID })

	p.storage.SaveConfigurations(configurations)
	if err != nil {
		return err
	}
	p.scheduler.Remove(id)
	return nil
}

func (p Peripherals) Verify(ctx context.Context, id uuid.UUID) (bool, error) {
	configuration, err := p.ByID(ctx, id)
	if err != nil || configuration.ID != id {
		return false, err
	}

	client := http.DefaultClient
	req, err := http.NewRequest(configuration.Method, configuration.Addr, nil)
	if err != nil {
		return false, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			p.log.Error("failed to close body: %v", err)
		}
	}()

	validStatus := resp.StatusCode-http.StatusOK < 100
	hasLength := resp.ContentLength > 0

	return validStatus && hasLength, nil
}

func (p Peripherals) Reload(ctx context.Context) error {
	configurations, err := p.storage.ReadConfigurations()
	if err != nil {
		p.log.Error(fmt.Sprintf("failed to read configurations: %v", err))
		return err
	}

	for _, configuration := range configurations {
		p.scheduler.Remove(configuration.ID)
		err := p.scheduler.Add(configuration)
		if err != nil {
			p.log.Warning(fmt.Sprintf("failed to add again configurations: %v", err))
		}
	}

	return nil
}

func (p Peripherals) LastReading(ctx context.Context, id uuid.UUID) (any, error) {
	data, err := p.scheduler.LastReading(id)
	return data, err
}
