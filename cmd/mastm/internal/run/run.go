package run

import "fmt"

func Run(args []string) int {
	handler := wire()

	if err := handler.Handle(args[1:]); err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		return 1
	}
	return 0
}
