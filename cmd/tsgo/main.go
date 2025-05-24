package main

import (
	"os"

	"github.com/Forge-AI/sand-typescript-go/public/execute"
)

func main() {
	os.Exit(runMain())
}

func runMain() int {
	args := os.Args[1:]
	if len(args) > 0 {
		switch args[0] {
		case "--lsp":
			return runLSP(args[1:])
		case "--api":
			return runAPI(args[1:])
		}
	}
	return int(execute.CommandLine(newSystem(), nil, args))
}
