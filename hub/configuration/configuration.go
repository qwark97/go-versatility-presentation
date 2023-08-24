package configuration

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/qwark97/go-versatility-presentation/hub/configuration/model"
)

var stubbedConfs = []model.ConfigurationEntity{
	{
		ID: uuid.MustParse("fe310f14-dfb6-4817-a12a-55cbf3417e3e"),
	},
	{
		ID: uuid.MustParse("51a7a0be-4101-4f41-a1b5-2f4184a017d9"),
	},
}

type Configuration struct {
}

func New(ctx context.Context, log *slog.Logger) Configuration {
	return Configuration{}
}

func (c Configuration) AllConfigurations(ctx context.Context) ([]model.ConfigurationEntity, error) {
	return stubbedConfs, nil
}

func (c Configuration) AddNewConfiguration(ctx context.Context, configuration model.ConfigurationEntity) error {
	configuration.ID = uuid.New()
	stubbedConfs = append(stubbedConfs, configuration)
	return nil
}
