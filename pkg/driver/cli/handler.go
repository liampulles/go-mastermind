package cli

import (
	"fmt"
	"strings"

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
		return nil, fmt.Errorf("must provide a command. valid commands: %s", h.commandsDescription())
	}
	commandName := args[0]
	commandMap := h.commandMap()
}

func (h *HandlerImpl) commandsDescription() string {
	var names []string
	for k := range h.commandMap() {
		names = append(names, k)
	}
	return strings.Join(names, ", ")
}

func (h *HandlerImpl) commandMap() map[string]Command {

}
