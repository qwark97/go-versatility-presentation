package main

import (
	"github.com/urfave/cli/v2"
)

const (
	defaultAddr = "localhost:9090"
)

func globalFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "addr",
			Category:    "network",
			DefaultText: defaultAddr,
			Value:       defaultAddr,
			Usage:       "allows to change default `ADDR` of the HUB",
			Aliases:     []string{"a"},
		},
	}
}
