package main

import (
	"github.com/qwark97/go-versatility-presentation/cli/flags"
	"github.com/qwark97/go-versatility-presentation/cli/presentation"
	"github.com/qwark97/go-versatility-presentation/cli/sensors/conf"
	"github.com/qwark97/go-versatility-presentation/cli/sensors/read"
	"github.com/qwark97/go-versatility-presentation/cli/sensors/reload"
	"github.com/urfave/cli/v2"
)

func newCLI() *cli.App {
	presenter := presentation.NewStdout()
	app := &cli.App{
		Name:        "sensors",
		Description: "CLI to manage configuration of the sensors used in Smart Home",
		Flags: []cli.Flag{
			flags.AddrFlag(),
		},
		Commands: []*cli.Command{
			conf.Command(presenter),
			reload.Command(presenter),
			read.Command(presenter),
		},
	}
	return app
}
