package conf

import (
	"fmt"
	"net/http"

	"github.com/urfave/cli/v2"
)

func get(ctx *cli.Context, presenter Presenter) error {
	id := ctx.String("id")
	all := ctx.Bool("all")
	addr := ctx.String("addr")

	if !all && id == "" {
		return fmt.Errorf("either -all or -id is required")
	}

	switch id {
	case "":
		var configurations []configuration
		uri := fmt.Sprintf("http://%s/api/v1/configurations", addr)
		err := apiRequestWithResponse(http.MethodGet, uri, nil, &configurations)
		if err != nil {
			return err
		}
		presenter.Show(configurations)
	default:
		var configuration configuration
		uri := fmt.Sprintf("http://%s/api/v1/configuration/%s", addr, id)
		err := apiRequestWithResponse(http.MethodGet, uri, nil, &configuration)
		if err != nil {
			return err
		}
		presenter.Show(configuration)
	}
	return nil
}
