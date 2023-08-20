package configuration

import (
	"context"
	"log/slog"
)

type Configuration struct {
}

func New(ctx context.Context, log *slog.Logger) Configuration {
	return Configuration{}
}
