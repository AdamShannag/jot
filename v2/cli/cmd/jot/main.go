package main

import (
	"fmt"
	"os"

	"github.com/AdamShannag/jot/v2/cli/cmd/jot/prompt"
)

func main() {
	if len(os.Args) > 1 {
		run(prompt.NewPrompterImpl(), os.Args[1])
	} else {
		fmt.Println("Jot v2.6.5")
	}
}

func run(prompter prompt.Prompter, command string) {
	prompter.Start(command)
}
