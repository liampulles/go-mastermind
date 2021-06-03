package print

import (
	"fmt"

	"github.com/liampulles/go-mastermind/pkg/domain"
	"github.com/liampulles/go-mastermind/pkg/usecase"
)

func Print(response *usecase.Response) {
	if response.GameWon {
		fmt.Println("You guessed correctly, and won!")
		return
	}
	formatted := format(response.Evaluation)
	fmt.Printf("Evaluation: %s\n", formatted)
}

func format(eval *domain.Evaluation) string {
	result := ""
	for _, rating := range eval {
		switch rating {
		case domain.Nothing:
			result += "[ ]"
		case domain.GoodColourBadPosition:
			result += "[~]"
		case domain.Perfect:
			result += "[✓]"
		default:
			result += "[?]"
		}
	}
	return result
}
