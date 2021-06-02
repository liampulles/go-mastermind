package main

import (
	"os"

	"github.com/liampulles/go-mastermind/cmd/mastm/internal/run"
)

func main() {
	os.Exit(run.Run(os.Args))
}
