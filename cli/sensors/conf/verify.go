package conf

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func verify(ctx *cli.Context, presenter Presenter) error {
	id := ctx.String("id")
	uri := fmt.Sprintf("http://%s/api/v1/configurations/%s/verify", ctx.String("addr"), id)
	return verifyConfigurations(uri)
}

func verifyConfigurations(uri string) error {
	fmt.Println("verifyConfigurations ID", uri)
	return nil
}
