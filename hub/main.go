package main

import (
	"context"
	"fmt"

	"log/slog"

	"github.com/qwark97/go-versatility-presentation/hub/flags"
	"github.com/qwark97/go-versatility-presentation/hub/peripherals"
	"github.com/qwark97/go-versatility-presentation/hub/scheduler"
	"github.com/qwark97/go-versatility-presentation/hub/server"
)

func main() {
	flagsConf := flags.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log := slog.Default()

	s := scheduler.New(ctx, log)
	go func() {
		if err := s.Start(); err != nil {
			log.Error(fmt.Sprintf("scheduler failed: %v", err))
		}
	}()

	p := peripherals.New(ctx, s, log)

	httpServer := server.New(p, flagsConf, log)
	if err := httpServer.Start(); err != nil {
		arg := slog.Any("error", err.Error())
		log.Error("server failed", arg)
	}
}
