package storage

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/qwark97/go-versatility-presentation/hub/logger"
)

type Conf interface {
}

type InMemStorage struct {
	conf Conf
	log  logger.Logger

	state *sync.Map
}

func New(conf Conf, log logger.Logger) *InMemStorage {
	return &InMemStorage{
		conf:  conf,
		log:   log,
		state: &sync.Map{},
	}
}

func (ims *InMemStorage) SaveLastReading(id uuid.UUID, data any) error {
	ims.state.Store(id, data)
	return nil
}
func (ims *InMemStorage) ReadLastReading(id uuid.UUID) (any, error) {
	val, present := ims.state.Load(id)
	if !present {
		return nil, fmt.Errorf("missing value with id: %s", id)
	}
	return val, nil
}
