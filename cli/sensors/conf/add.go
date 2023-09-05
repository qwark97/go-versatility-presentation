package conf

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/urfave/cli/v2"
)

func add(ctx *cli.Context, presenter Presenter) error {
	addr := ctx.String("addr")

	newConfiguration := gatherConfigurationFromInput(presenter)
	data, err := json.Marshal(newConfiguration)
	if err != nil {
		return err
	}
	body := bytes.NewReader(data)
	uri := fmt.Sprintf("http://%s/api/v1/configuration", addr)
	err = apiRequest(http.MethodPost, uri, body)
	if err != nil {
		return err
	}
	return nil
}

func gatherConfigurationFromInput(presenter Presenter) configuration {
	var newConfiguration = configuration{}
	presenter.Show("Procedure of adding new configuration:")
	presenter.Show("HTTP method and URL space-separated:")
	fmt.Scan(&newConfiguration.Method, &newConfiguration.Addr)
	presenter.Show("Frequency of requests (eg. 1s):")
	fmt.Scan(&newConfiguration.Frequency)
	presenter.Show("Unit of received result:")
	fmt.Scan(&newConfiguration.Unit)
	presenter.Show("Description of the configuration:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	newConfiguration.Description = scanner.Text()
	return newConfiguration
}
