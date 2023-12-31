package conf

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func apiRequestWithResponse[T any](method, uri string, requestBody io.Reader, responseContainer *T) error {
	client := http.DefaultClient

	request, err := http.NewRequest(method, uri, requestBody)
	if err != nil {
		return err
	}

	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode%200 >= 100 {
		return fmt.Errorf("invalid response status code: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, responseContainer)
	if err != nil {
		return err
	}
	return nil
}

func apiRequest(method, uri string, requestBody io.Reader) error {
	client := http.DefaultClient

	request, err := http.NewRequest(method, uri, requestBody)
	if err != nil {
		return err
	}

	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode%200 >= 100 {
		return fmt.Errorf("invalid response status code: %s", resp.Status)
	}

	return nil
}
