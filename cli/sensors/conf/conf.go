package conf

import (
	"fmt"

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
				Name: "get",
				Flags: []cli.Flag{
					idFlag(),
					allFlag(),
				},
				Action: func(ctx *cli.Context) error {
					stringID := ctx.String("id")
					if !ctx.Bool("all") && stringID == "" {
						return fmt.Errorf("either -all or -id is required")
					}

					switch stringID {
					case "":
						uri := fmt.Sprintf("http://%s/api/v1/configurations", ctx.String("addr"))
						configurations, err := getConfigurations(uri)
						if err != nil {
							return err
						}
						presenter.Show(configurations)
					default:
						uri := fmt.Sprintf("http://%s/api/v1/configuration/%s", ctx.String("addr"), ctx.String("id"))
						configuration, err := getConfiguration(uri)
						if err != nil {
							return err
						}
						presenter.Show(configuration)
					}
					return nil
				},
			},
			{
				Name: "add",
				Action: func(ctx *cli.Context) error {
					return nil
				},
			},
			{
				Name: "rm",
				Flags: []cli.Flag{
					idFlag(),
					allFlag(),
				},
				Action: func(ctx *cli.Context) error {
					id := ctx.String("id")
					if !ctx.Bool("all") && id == "" {
						return fmt.Errorf("either -all or -id is required")
					}

					if id != "" {
						uri := fmt.Sprintf("http://%s/api/v1/configuration/%s", ctx.String("addr"), ctx.String("id"))
						return removeConfiguration(uri)
					} else {
						uri := fmt.Sprintf("http://%s/api/v1/configurations", ctx.String("addr"))
						configurations, err := getConfigurations(uri)
						if err != nil {
							return err
						}
						for _, id := range configurations {
							uri := fmt.Sprintf("http://%s/api/v1/configuration/%s", ctx.String("addr"), id)
							return removeConfiguration(uri)
						}
					}
					return nil
				},
			},
			{
				Name: "verify",
				Flags: []cli.Flag{
					idRequiredFlag(),
				},
				Action: func(ctx *cli.Context) error {
					id := ctx.String("id")
					uri := fmt.Sprintf("http://%s/api/v1/configurations/%s/verify", ctx.String("addr"), id)
					return verifyConfigurations(uri)
				},
			},
		},
	}
}

func getConfiguration(uri string) (any, error) {
	fmt.Println("showConfiguration ID", uri)
	return nil, nil
}

func getConfigurations(uri string) ([]any, error) {
	fmt.Println("showConfigurations")
	return []any{nil}, nil
}

func removeConfiguration(uri string) error {
	fmt.Println("removeConfiguration ID", uri)
	return nil
}

func verifyConfigurations(uri string) error {
	fmt.Println("verifyConfigurations ID", uri)
	return nil
}
