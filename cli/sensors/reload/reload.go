package reload

import (
	"fmt"
	"net/http"

	"github.com/urfave/cli/v2"
)

type Presenter interface {
	Show(data any)
}

func Command(presenter Presenter) *cli.Command {
	return &cli.Command{
		Name: "reload",
		Action: func(ctx *cli.Context) error {
			uri := fmt.Sprintf("http://%s/api/v1/configurations/reload", ctx.String("addr"))
			err := reloadConfigurations(uri)
			if err != nil {
				return err
			}
			presenter.Show("configuration reloaded")
			return nil
		},
	}
}

func reloadConfigurations(uri string) error {
	client := http.DefaultClient
	resp, err := client.Post(uri, "application/json", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("invalid response: %s", resp.Status)
	}
	return nil
}
