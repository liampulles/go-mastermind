package cli

import (
	"fmt"

	"github.com/liampulles/go-mastermind/pkg/usecase"
)

type Handler interface {
	Handle(args []string) error
}

type HandlerImpl struct {
	engine usecase.HumanEngine
}

var _ Handler = &usecase.HumanEngineImpl{}

func NewHandlerImpl(engine usecase.HumanEngine) *HandlerImpl {
	return &HandlerImpl{
		engine: engine,
	}
}

func (h *HandlerImpl) Handle(args []string) error {

}

type Command func(args []string) error

func (h *HandlerImpl) resolveCommand(args []string) (Command, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("must provide a command")
	}
}

func commandDesc()

func (h *HandlerImpl) commandMap() map[string]Command {

}
