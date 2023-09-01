package main

import (
	"os"

	"github.com/AdamShannag/jot/v2/cli/cmd/jot/prompt"
)

func main() {
	run(prompt.NewPrompterImpl(), os.Args[1])
}

func run(prompter prompt.Prompter, command string) {
	prompter.Start(command)
}
