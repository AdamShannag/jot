package command

import (
	"github.com/AdamShannag/jot/internal/command/new"
	p "github.com/AdamShannag/jot/internal/command/path"
	"github.com/AdamShannag/jot/internal/format"
	"github.com/AdamShannag/jot/internal/io"
	"github.com/AdamShannag/jot/internal/types"
	"github.com/AdamShannag/jot/internal/validate"
	"github.com/urfave/cli/v2"
)

func ini() *cli.Command {
	return new.Command(
		"init",
		[]string{"i"},
		"creation",
		"Create a new project at the specified path.",
		"init [project-path] [project-name]",
		"This command will create a new project at the specified path, along with any necessary files and directories.\nOnce created, you can then begin working on the project as needed.",
		nil,
		func(cCtx *cli.Context) error {

			projectPath := format.Path(cCtx.Args().First())
			projectName := format.ProjectName(cCtx.Args().Get(1))

			if err := validate.Name(projectName); err != nil {
				return err
			}

			if err := validate.Path(projectPath); err != nil {
				return err
			}

			s := types.NewSpecs(projectName, nil, nil)
			if b, err := types.ToYamlString(s); err == nil {
				io.ToEmptyFile(projectPath, p.JotFile, b)
			} else {
				return err
			}

			return nil
		},
		onUsageError,
	)
}
