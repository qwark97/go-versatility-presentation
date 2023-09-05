package conf

import (
	"fmt"
	"net/http"

	"github.com/urfave/cli/v2"
)

func rm(ctx *cli.Context, presenter Presenter) error {
	id := ctx.String("id")
	all := ctx.Bool("all")
	addr := ctx.String("addr")

	if !all && id == "" {
		return fmt.Errorf("either -all or -id is required")
	}

	if id != "" {
		uri := fmt.Sprintf("http://%s/api/v1/configuration/%s", addr, id)
		return apiRequest(http.MethodDelete, uri, nil)
	} else {
		var configurations []configuration
		uri := fmt.Sprintf("http://%s/api/v1/configurations", addr)
		err := apiRequestWithResponse(http.MethodGet, uri, nil, &configurations)
		if err != nil {
			return err
		}
		for _, id := range configurations {
			uri := fmt.Sprintf("http://%s/api/v1/configuration/%s", addr, id.ID)
			err = apiRequest(http.MethodDelete, uri, nil)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
