package command

import (
	"errors"
	"fmt"
	"strings"

	"github.com/AdamShannag/jot/internal/command/new"
	"github.com/AdamShannag/jot/internal/io"
	"github.com/AdamShannag/jot/internal/template"
	"github.com/AdamShannag/jot/internal/types"
	"github.com/AdamShannag/jot/internal/types/tpls"
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
			ok, err := io.FileExists(jotRelPath + jotFile)
			if err != nil {
				return err
			}
			if !ok {
				return errors.New(jotFile + " file not found in current working directory")
			}

			// convert jot.yaml to Specs
			specs, err := io.ToSpecs(jotRelPath + jotFile)
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
				updateEndpoints(endpoints, specs, i, service)
			} else {
				createNewService(service, isRest, endpoints, specs, port)
			}

			// write new specs to jot.yaml file
			if b, err := types.ToYamlString(specs); err == nil {
				io.ToFile(jotRelPath, jotFile, b)
			} else {
				return err
			}

			return nil
		},
		onUsageError,
	)
	return cmd
}

func createNewService(service string, isRest bool, endpoints []string, specs *types.Specs, port int) {
	srv := service
	serviceSuffix(&srv)

	// create go.mod
	template.Create(goModTpl, path(goModPath, srv), goModFileName,
		tpls.GoModule{ModuleName: srv})

	// create service directoies
	createService(service, isRest)

	// create .dockerfile
	template.Create(dockerImageTpl, path(dockerImagePath, srv), dockerfile(srv),
		tpls.Docker{AppName: appSuffix(service), AppPath: path(appPath, appSuffix(service))})

	// create endpoints directories if rest enabled
	if isRest && len(endpoints) > 0 {
		createEndpoints(service, endpoints)
	}

	// update specs file
	specs.Services = append(specs.Services, *types.NewService(service, port, endpoints))
}

func updateEndpoints(endpoints []string, specs *types.Specs, i int, service string) {
	newEps := []string{}

	// create endpoints if new, ignore if existing
	for _, ep := range endpoints {
		if !contains(specs.Services[i].Endpoints, ep) {
			createEndpoint(service, ep)
			newEps = append(newEps, ep)
		} else {
			logM("Endpoint", IGNORED)
		}
	}

	// update specs file
	specs.Services[0].Endpoints = append(specs.Services[0].Endpoints, newEps...)
}

func createService(name string, rest bool) {
	service := name
	name = strings.Replace(name, "-service", "", 1)
	serviceSuffix(&service)

	// create service directories cmd, bin, api

	io.ToDirs(fmt.Sprintf(mainDirPath, service, name))
	io.ToDirs(path(binDirPath, service))
	io.ToDirs(path(apiDirPath, service))

	// create api directories if rest handler, middleware
	if rest {
		template.Create(apiTpl, path(apiDirPath, service), apiFileName, nil)
		io.ToDirs(path(handlerDirPath, service))
		io.ToDirs(path(middelwareDirPath, service))
	}

	logM("Service", CREATED)
}

func createEndpoint(service string, endpoint string) {
	serviceSuffix(&service)

	// create handlers files
	template.Create(handlerTpl, fmt.Sprintf(handlerPath, service, endpoint), goSuffix(endpoint),
		tpls.Handler{EndpointName: titleCase(endpoint)})
}

func createEndpoints(service string, endpoints []string) {
	serviceSuffix(&service)
	for _, e := range endpoints {
		createEndpoint(service, e)
	}
}

func serviceSuffix(name *string) {
	if !strings.HasSuffix(*name, "-service") {
		*name += "-service"
	}
}
