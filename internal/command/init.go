package command

import (
	"errors"
	"fmt"

	"github.com/AdamShannag/jot/internal/command/new"
	"github.com/AdamShannag/jot/internal/io"
	"github.com/AdamShannag/jot/internal/types"
	"github.com/urfave/cli/v2"
)

func ini() *cli.Command {
	return new.Command(
		"init",
		[]string{"i"},
		"creation",
		"Create a new project with the specified path and name.",
		"init [project-path] [project-name]",
		"This command will create a new project with the specified path and name, along with any necessary files and directories.\nOnce created, you can then begin working on the project as needed.",
		nil,
		func(cCtx *cli.Context) error {
			if cCtx.Args().Len() < 2 {
				return errors.New("you must specify a path and project name when using this command, see help for more")
			}

			s := types.NewSpecs(cCtx.Args().Get(1), nil, nil)

			path := fmt.Sprintf(projectDirPath, cCtx.Args().Get(0), cCtx.Args().Get(1))
			if b, err := types.ToYamlString(s); err == nil {
				io.ToEmptyFile(path, jotFile, b)
			} else {
				return err
			}

			return nil
		},
		onUsageError,
	)
}
