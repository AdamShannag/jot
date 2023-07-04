package command

import (
	"errors"

	"github.com/AdamShannag/jot/internal/command/feature"
	"github.com/AdamShannag/jot/internal/command/new"
	p "github.com/AdamShannag/jot/internal/command/path"
	"github.com/AdamShannag/jot/internal/command/suffix"
	"github.com/AdamShannag/jot/internal/config"
	"github.com/AdamShannag/jot/internal/io"
	"github.com/AdamShannag/jot/internal/makefile"
	"github.com/AdamShannag/jot/internal/spinner"
	"github.com/AdamShannag/jot/internal/types"
	"github.com/urfave/cli/v2"
)

func add() *cli.Command {
	cmd := new.Command(
		"add",
		[]string{"a"},
		"creation",
		"Create a new service with the specified type and name.",
		"add -srv [service] -p [port]",
		"This command will create a new service or resource with the specified type and name, along with any necessary files and directories.\nOnce created, you can then begin working on the project as needed.",
		[]cli.Flag{
			new.StringFlag("service", "", "creates a service", true, "srv"),
			new.StringSliceFlag("endpoints", "", "create enpoints", false, "end"),
			new.StringSliceFlag("middlewares", "", "create middlewares", false, "mid"),
			new.IntFlag("port", 0, "specifies the port of the service", true, "p"),
			new.BoolFlag("rest", false, "make a rest api", false, "rs"),
			new.BoolFlag("crud", false, "make a crud file", false, "c"),
		},
		func(cCtx *cli.Context) error {

			specs, err := getSpecs()
			if err != nil {
				return err
			}

			mk := makefile.New(p.Path(p.GoModPath, suffix.ServiceSuffix(cCtx.String("service"))), config.MAKEFILE_TIMEOUT, spinner.New("cyan"))
			defer mk.Build()

			feat := feature.New(specs, mk, cCtx)
			if err := feat.BuildREST(); err != nil {
				return err
			}
			if err := feat.BuildGRPC(); err != nil {
				return err
			}
			if err := feat.BuildMiddlewares(); err != nil {
				return err
			}

			if err := updateSpecs(specs); err != nil {
				return err
			}

			return nil
		},
		onUsageError,
	)
	return cmd
}

func getSpecs() (*types.Specs, error) {
	ok, err := io.FileExists(p.JotRelPath + p.JotFile)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New(p.JotFile + " file not found in current working directory")
	}

	specs, err := io.ToSpecs(p.JotRelPath + p.JotFile)
	if err != nil {
		return nil, err
	}

	return specs, nil
}

func updateSpecs(specs *types.Specs) error {
	if b, err := types.ToYamlString(specs); err == nil {
		io.ToFile(p.JotRelPath, p.JotFile, b)
	} else {
		return err
	}

	return nil
}
