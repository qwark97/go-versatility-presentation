package peripherals

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path"
	"path/filepath"
	"slices"

	"github.com/google/uuid"
	"github.com/qwark97/go-versatility-presentation/hub/peripherals/model"
)

type Scheduler interface{}

type Conf interface {
	Path() string
}

type Peripherals struct {
	ctx       context.Context
	scheduler Scheduler
	conf      Conf
	log       *slog.Logger
}

func New(ctx context.Context, scheduler Scheduler, conf Conf, log *slog.Logger) Peripherals {
	return Peripherals{
		ctx:       ctx,
		scheduler: scheduler,
		conf:      conf,
		log:       log,
	}
}

func (p Peripherals) assureDir() error {
	absPath, err := p.absPath()
	if err != nil {
		return err
	}
	absDir := path.Dir(absPath)

	_, err = os.Stat(absDir)
	if !errors.Is(err, os.ErrNotExist) {
		return nil
	}

	err = os.MkdirAll(absDir, 0755)
	if err != nil {
		p.log.Error("failed to create dirs tree")
		return err
	}
	return nil
}

func (p Peripherals) absPath() (string, error) {
	absDir, err := filepath.Abs(p.conf.Path())
	if err != nil {
		p.log.Error("failed to get absolute path")
	}
	return absDir, err
}

func (p Peripherals) All(ctx context.Context) ([]model.Configuration, error) {
	configurations, err := p.readConfigurations()
	if err != nil {
		p.log.Error(fmt.Sprintf("failed to read configurations: %v", err))
		return nil, err
	}

	return configurations, nil
}

func (p Peripherals) Add(ctx context.Context, configuration model.Configuration) error {
	err := p.assureDir()
	if err != nil {
		return err
	}

	configurations, err := p.readConfigurations()
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
	}

	configuration.ID = uuid.New()
	configurations = append(configurations, configuration)

	return p.saveConfigurations(configurations)
}

func (p Peripherals) ByID(ctx context.Context, id uuid.UUID) (model.Configuration, error) {
	var notFound model.Configuration

	configurations, err := p.readConfigurations()
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
	configurations, err := p.readConfigurations()
	if err != nil {
		p.log.Error(fmt.Sprintf("failed to read configurations: %v", err))
		return err
	}

	configurations = slices.DeleteFunc(configurations, func(c model.Configuration) bool { return c.ID == id })

	return p.saveConfigurations(configurations)
}

func (p Peripherals) Verify(ctx context.Context, id uuid.UUID) (bool, error) {
	return true, nil
}

func (p Peripherals) Reload(ctx context.Context) error {
	return nil
}

func (p Peripherals) readConfigurations() ([]model.Configuration, error) {
	absPath, err := p.absPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		return nil, err
	}

	var container []model.Configuration
	err = json.Unmarshal(data, &container)
	if err != nil {
		p.log.Error(fmt.Sprintf("failed to unmarshal data: %v", err))
		return nil, err
	}

	return container, nil
}

func (p Peripherals) saveConfigurations(configurations []model.Configuration) error {
	absPath, err := p.absPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(configurations, "", "\t")
	if err != nil {
		p.log.Error(fmt.Sprintf("failed to marshal data: %v", err))
		return err
	}

	return os.WriteFile(absPath, data, 0755)
}
