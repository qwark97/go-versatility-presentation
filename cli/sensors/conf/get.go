package conf

import (
	"fmt"
	"net/http"

	"github.com/urfave/cli/v2"
)

func get(ctx *cli.Context, presenter Presenter) error {
	stringID := ctx.String("id")
	if !ctx.Bool("all") && stringID == "" {
		return fmt.Errorf("either -all or -id is required")
	}

	switch stringID {
	case "":
		var configurations []Configuration
		uri := fmt.Sprintf("http://%s/api/v1/configurations", ctx.String("addr"))
		err := apiRequest(http.MethodGet, uri, nil, &configurations)
		if err != nil {
			return err
		}
		presenter.Show(configurations)
	default:
		var configuration Configuration
		uri := fmt.Sprintf("http://%s/api/v1/configuration/%s", ctx.String("addr"), ctx.String("id"))
		err := apiRequest(http.MethodGet, uri, nil, &configuration)
		if err != nil {
			return err
		}
		presenter.Show(configuration)
	}
	return nil
}
