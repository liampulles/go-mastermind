package cli

import (
	"fmt"
	"strings"

	"github.com/liampulles/go-mastermind/pkg/driver/print"
	"github.com/liampulles/go-mastermind/pkg/usecase"
)

type Handler interface {
	Handle(args []string) error
}

type HandlerImpl struct {
	engine usecase.HumanEngine
}

var _ Handler = &HandlerImpl{}

func NewHandlerImpl(engine usecase.HumanEngine) *HandlerImpl {
	return &HandlerImpl{
		engine: engine,
	}
}

func (h *HandlerImpl) Handle(args []string) error {
	command, err := h.resolveCommand(args)
	if err != nil {
		return err
	}
	return command(args[1:])
}

func (h *HandlerImpl) new(args []string) error {
	if len(args) != 0 {
		return fmt.Errorf("the new command takes no arguments")
	}
	if err := h.engine.NewGame(); err != nil {
		return fmt.Errorf("could not create game: %w", err)
	}
	return nil
}

func (h *HandlerImpl) guess(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("the new command takes execatly one argument, which is the guess")
	}
	request := usecase.Request(args[0])

	response, err := h.engine.Evaluate(&request)
	if err != nil {
		return fmt.Errorf("could not evaluate guess: %w", err)
	}

	print.Print(response)
	return nil
}

type Command func(args []string) error

func (h *HandlerImpl) resolveCommand(args []string) (Command, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("must provide a command. valid commands: %s", h.commandsDescription())
	}
	commandName := args[0]
	commandMap := h.commandMap()
	command, ok := commandMap[commandName]
	if !ok {
		return nil, fmt.Errorf("no such command %s. valid commands: %s", commandName, h.commandsDescription())
	}
	return command, nil
}

func (h *HandlerImpl) commandsDescription() string {
	var names []string
	for k := range h.commandMap() {
		names = append(names, k)
	}
	return strings.Join(names, ", ")
}

func (h *HandlerImpl) commandMap() map[string]Command {
	return map[string]Command{
		"new":   h.new,
		"guess": h.guess,
	}
}
