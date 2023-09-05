package conf

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
)

func idFlag() *cli.StringFlag {
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

func idRequiredFlag() *cli.StringFlag {
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

func allFlag() *cli.BoolFlag {
	return &cli.BoolFlag{
		Name: "all",
	}
}
