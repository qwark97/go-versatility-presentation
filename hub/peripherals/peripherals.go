package peripherals

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/qwark97/go-versatility-presentation/hub/peripherals/model"
)

var stubbedConfs = []model.Configuration{
	{
		ID: uuid.MustParse("fe310f14-dfb6-4817-a12a-55cbf3417e3e"),
	},
	{
		ID: uuid.MustParse("51a7a0be-4101-4f41-a1b5-2f4184a017d9"),
	},
}

type Peripherals struct {
}

func New(ctx context.Context, log *slog.Logger) Peripherals {
	return Peripherals{}
}

func (c Peripherals) All(ctx context.Context) ([]model.Configuration, error) {
	return stubbedConfs, nil
}

func (c Peripherals) Add(ctx context.Context, configuration model.Configuration) error {
	configuration.ID = uuid.New()
	stubbedConfs = append(stubbedConfs, configuration)
	return nil
}
