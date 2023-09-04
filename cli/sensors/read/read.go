package read

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
)

type Presenter interface {
	Show(data any)
}

func Command(presenter Presenter) *cli.Command {
	return &cli.Command{
		Name: "read",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "id",
				Required: true,
				Action: func(ctx *cli.Context, s string) error {
					if _, err := uuid.Parse(ctx.String("id")); err != nil {
						return fmt.Errorf("invalid -id flag: %s", err)
					}
					return nil
				},
			},
		},
		Action: func(ctx *cli.Context) error {
			id := ctx.String("id")
			uri := fmt.Sprintf("http://%s/api/v1/last-reading/%s", ctx.String("addr"), id)

			newCtx, cancel := context.WithCancel(ctx.Context)
			defer cancel()

			dataCh, errCh := readLastData(newCtx, uri)
			for res := range dataCh {
				presenter.Show(res)
			}
			errs := []error{}
			for err := range errCh {
				errs = append(errs, err)
			}
			err := errors.Join(errs...)
			if err != nil {
				return err
			}
			return nil
		},
	}
}

func readLastData(ctx context.Context, uri string) (<-chan string, <-chan error) {
	var (
		dataCh = make(chan string)
		errCh  = make(chan error)
	)

	go func(ctx context.Context) {
		client := http.DefaultClient
		client.Timeout = time.Second

		defer close(dataCh)
		defer close(errCh)

		for {
			select {
			case <-ctx.Done():
				errCh <- ctx.Err()
				return
			default:
				func() {
					ctx, cancel := context.WithTimeout(ctx, time.Second)
					defer cancel()

					req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
					if err != nil {
						errCh <- err
						return
					}

					resp, err := client.Do(req)
					if err != nil {
						errCh <- err
						return
					}
					defer resp.Body.Close()

					if resp.StatusCode != http.StatusOK {
						errCh <- fmt.Errorf("invalid http status: %s", resp.Status)
						return
					}
					data, err := io.ReadAll(resp.Body)
					if err != nil {
						errCh <- err
						return
					}
					dataCh <- string(data)
				}()
				time.Sleep(time.Second)
			}
		}
	}(ctx)
	return dataCh, errCh
}
