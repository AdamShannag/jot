package command

import (
	"github.com/urfave/cli/v2"
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
