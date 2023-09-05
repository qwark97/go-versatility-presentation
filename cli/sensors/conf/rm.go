package conf

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func rm(ctx *cli.Context, presenter Presenter) error {
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
}

func removeConfiguration(uri string) error {
	fmt.Println("removeConfiguration ID", uri)
	return nil
}
