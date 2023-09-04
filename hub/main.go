package main

import (
	"context"

	"github.com/qwark97/go-versatility-presentation/hub/flags"
	"github.com/qwark97/go-versatility-presentation/hub/logger"
	"github.com/qwark97/go-versatility-presentation/hub/peripherals"
	fileStorage "github.com/qwark97/go-versatility-presentation/hub/peripherals/storage"
	"github.com/qwark97/go-versatility-presentation/hub/scheduler"
	inMemStorage "github.com/qwark97/go-versatility-presentation/hub/scheduler/storage"
	"github.com/qwark97/go-versatility-presentation/hub/server"
)

func main() {
	flagsConf := flags.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log := logger.New()

	inMemS := inMemStorage.New(flagsConf, log)
	s := scheduler.New(ctx, inMemS, log)
	fileS := fileStorage.New(flagsConf, log)
	p := peripherals.New(ctx, s, fileS, log)
	err := p.Reload(ctx)
	if err != nil {
		log.Info("did not find initial peripherals")
	}

	httpServer := server.New(p, flagsConf, log)
	if err := httpServer.Start(); err != nil {
		log.Error("server failed: %v", err)
	}
}
