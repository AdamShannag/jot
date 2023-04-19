package command

import (
	"github.com/urfave/cli/v2"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var commands = []*cli.Command{}

func Commands() []*cli.Command {
	addCommands(
		ini(),
		add(),
	)
	return commands
}

func addCommands(cmd ...*cli.Command) {
	commands = append(commands, cmd...)
}

func contains(slice []string, item string) bool {
	for _, i := range slice {
		if i == item {
			return true
		}
	}
	return false
}

func titleCase(s string) string {
	return cases.Title(language.English, cases.NoLower).String(s)
}
