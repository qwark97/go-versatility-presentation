package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/qwark97/go-versatility-presentation/hub/logger"
)

type Conf interface {
	Path() string
}

type Storage struct {
	conf Conf
	log  logger.Logger
}

func New(conf Conf, log logger.Logger) Storage {
	return Storage{
		conf: conf,
		log:  log,
	}
}

func (s Storage) AssureDir() error {
	absPath, err := s.absPath()
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
		s.log.Error("failed to create dirs tree")
		return err
	}
	return nil
}

func (s Storage) ReadConfigurations() ([]Configuration, error) {
	absPath, err := s.absPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		return nil, err
	}

	var container []Configuration
	err = json.Unmarshal(data, &container)
	if err != nil {
		s.log.Error(fmt.Sprintf("failed to unmarshal data: %v", err))
		return nil, err
	}

	return container, nil
}

func (s Storage) SaveConfigurations(configurations []Configuration) error {
	absPath, err := s.absPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(configurations, "", "\t")
	if err != nil {
		s.log.Error(fmt.Sprintf("failed to marshal data: %v", err))
		return err
	}

	return os.WriteFile(absPath, data, 0755)
}

func (s Storage) absPath() (string, error) {
	absDir, err := filepath.Abs(s.conf.Path())
	if err != nil {
		s.log.Error("failed to get absolute path")
	}
	return absDir, err
}
