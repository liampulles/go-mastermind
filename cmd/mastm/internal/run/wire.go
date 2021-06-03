package run

import (
	"github.com/liampulles/go-mastermind/pkg/domain"
	"github.com/liampulles/go-mastermind/pkg/driver/cli"
	"github.com/liampulles/go-mastermind/pkg/driver/yaml"
	"github.com/liampulles/go-mastermind/pkg/usecase"
)

func wire() cli.Handler {
	engine := domain.NewEngineImpl()

	factory := usecase.NewFactoryImpl(engine)
	yamlStore := yaml.NewGameStore()

	humanEngine := usecase.NewHumanEngineImpl(factory, yamlStore, engine)

	return cli.NewHandlerImpl(humanEngine)
}
