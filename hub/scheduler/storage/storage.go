package storage

import (
	"github.com/qwark97/go-versatility-presentation/hub/logger"
)

type Conf interface {
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
