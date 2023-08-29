package main

import (
	"context"

	"github.com/qwark97/go-versatility-presentation/hub/flags"
	"github.com/qwark97/go-versatility-presentation/hub/logger"
	"github.com/qwark97/go-versatility-presentation/hub/peripherals"
	"github.com/qwark97/go-versatility-presentation/hub/peripherals/storage"
	"github.com/qwark97/go-versatility-presentation/hub/scheduler"
	"github.com/qwark97/go-versatility-presentation/hub/server"
)

func main() {
	flagsConf := flags.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log := logger.New()

	s := scheduler.New(ctx, log)
	go func() {
		if err := s.Start(); err != nil {
			log.Error("scheduler failed: %v", err)
			panic(err)
		}
	}()

	fs := storage.New(flagsConf, log)
	p := peripherals.New(ctx, s, fs, log)
	httpServer := server.New(p, flagsConf, log)
	if err := httpServer.Start(); err != nil {
		log.Error("server failed: %v", err)
	}
}
