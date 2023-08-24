package main

import (
	"context"

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

	p := peripherals.New(ctx, log)
	s := scheduler.New(ctx, log)

	httpServer := server.New(p, s, flagsConf, log)
	if err := httpServer.Start(); err != nil {
		arg := slog.Any("error", err.Error())
		log.Error("server failed", arg)
	}
}
