package conf

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func get(ctx *cli.Context, presenter Presenter) error {
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
}

func getConfiguration(uri string) (any, error) {
	fmt.Println("showConfiguration ID", uri)
	return nil, nil
}

func getConfigurations(uri string) ([]any, error) {
	fmt.Println("showConfigurations")
	return []any{nil}, nil
}
