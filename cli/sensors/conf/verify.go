package conf

import (
	"fmt"
	"net/http"

	"github.com/urfave/cli/v2"
)

func verify(ctx *cli.Context, presenter Presenter) error {
	id := ctx.String("id")
	addr := ctx.String("addr")
	var status verification
	uri := fmt.Sprintf("http://%s/api/v1/configuration/%s/verify", addr, id)
	err := apiRequestWithResponse(http.MethodPost, uri, nil, &status)
	if err != nil {
		return err
	}
	msg := fmt.Sprintf("configuration with id: %s ", id)
	if status.Success {
		msg += "is valid"
	} else {
		msg += "is invalid"
	}
	presenter.Show(msg)
	return nil
}
