package read

import (
	"fmt"
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
			return readLastData(uri)
		},
	}
}

func readLastData(uri string) error {
	for {
		fmt.Println("asd: ", uri)
		time.Sleep(time.Second)
	}
}
