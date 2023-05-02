package command

import (
	"errors"

	"github.com/AdamShannag/jot/internal/command/endpoint"
	"github.com/AdamShannag/jot/internal/command/new"
	p "github.com/AdamShannag/jot/internal/command/path"
	srv "github.com/AdamShannag/jot/internal/command/service"
	"github.com/AdamShannag/jot/internal/io"
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
			new.IntFlag("port", 0, "specifies the port of the service", true, "p"),
			new.BoolFlag("rest", false, "make a rest api", false, "rs"),
		},
		func(cCtx *cli.Context) error {
			// check if we have jot.yaml file
			ok, err := io.FileExists(p.JotRelPath + p.JotFile)
			if err != nil {
				return err
			}
			if !ok {
				return errors.New(p.JotFile + " file not found in current working directory")
			}

			// convert jot.yaml to Specs
			specs, err := io.ToSpecs(p.JotRelPath + p.JotFile)
			if err != nil {
				return err
			}

			// read all flags
			isRest := cCtx.Bool("rest")
			port := cCtx.Int("port")
			endpoints := cCtx.StringSlice("endpoints")
			service := cCtx.String("service")

			// endpoints specified but rest flag is off
			if !isRest && len(endpoints) > 0 {
				return errors.New("--endpoints flag is specified but --rest flag is not")
			}

			// check if new service or not
			if ok, i := types.IsExistingService(specs.Services, service); ok {
				endpoint.UpdateAll(endpoints, specs, i, service)
			} else {
				srv.New(service, isRest, endpoints, specs, port)
			}

			// write new specs to jot.yaml file
			if b, err := types.ToYamlString(specs); err == nil {
				io.ToFile(p.JotRelPath, p.JotFile, b)
			} else {
				return err
			}

			return nil
		},
		onUsageError,
	)
	return cmd
}
