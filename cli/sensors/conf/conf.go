package conf

import (
	"github.com/qwark97/go-versatility-presentation/cli/flags"
	"github.com/urfave/cli/v2"
)

type Presenter interface {
	Show(data any)
}

func Command(presenter Presenter) *cli.Command {
	return &cli.Command{
		Name: "conf",
		Subcommands: []*cli.Command{
			{
				Name:  "get",
				Usage: "allows to read configuration by given ID",
				Flags: []cli.Flag{
					flags.IDFlag(),
					flags.AllFlag(),
				},
				Action: func(ctx *cli.Context) error {
					return get(ctx, presenter)
				},
			},
			{
				Name:  "add",
				Usage: "allows to add new configuration entity",
				Action: func(ctx *cli.Context) error {
					return add(ctx, presenter)
				},
			},
			{
				Name:  "rm",
				Usage: "allows to remove configuration by given ID",
				Flags: []cli.Flag{
					flags.IDFlag(),
					flags.AllFlag(),
				},
				Action: func(ctx *cli.Context) error {
					return rm(ctx, presenter)
				},
			},
			{
				Name:  "verify",
				Usage: "allows to verify if all defined configurations are valid",
				Flags: []cli.Flag{
					flags.IDRequiredFlag(),
				},
				Action: func(ctx *cli.Context) error {
					return verify(ctx, presenter)
				},
			},
		},
	}
}
