package flags

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
)

const (
	defaultAddr = "localhost:9090"
)

func AddrFlag() *cli.StringFlag {
	return &cli.StringFlag{
		Name:        "addr",
		Category:    "network",
		DefaultText: defaultAddr,
		Value:       defaultAddr,
		Usage:       "allows to change default `ADDR` of the HUB",
		Aliases:     []string{"a"},
	}
}

func IDFlag() *cli.StringFlag {
	return &cli.StringFlag{
		Name: "id",
		Action: func(ctx *cli.Context, s string) error {
			if _, err := uuid.Parse(s); err != nil {
				return fmt.Errorf("invalid -id flag: %s", err)
			}
			return nil
		},
	}
}

func IDRequiredFlag() *cli.StringFlag {
	return &cli.StringFlag{
		Name:     "id",
		Required: true,
		Action: func(ctx *cli.Context, s string) error {
			if _, err := uuid.Parse(s); err != nil {
				return fmt.Errorf("invalid -id flag: %s", err)
			}
			return nil
		},
	}
}

func AllFlag() *cli.BoolFlag {
	return &cli.BoolFlag{
		Name: "all",
	}
}
