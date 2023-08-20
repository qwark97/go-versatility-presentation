package main

import (
	"context"

	"log/slog"

	"github.com/qwark97/go-versatility-presentation/hub/configuration"
	"github.com/qwark97/go-versatility-presentation/hub/flags"
	"github.com/qwark97/go-versatility-presentation/hub/scheduler"
	"github.com/qwark97/go-versatility-presentation/hub/server"
)

func main() {
	flagsConf := flags.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log := slog.Default()

	confService := configuration.New(ctx, log)
	schedService := scheduler.New(ctx, log)

	httpServer := server.New(ctx, confService, schedService, flagsConf, log)
	if err := httpServer.Start(); err != nil {
		arg := slog.Any("error", err.Error())
		log.Error("server failed", arg)
	}
}
