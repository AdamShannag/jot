package service

import (
	"github.com/AdamShannag/jot/internal/command/endpoint"
	"github.com/AdamShannag/jot/internal/command/log"
	"github.com/AdamShannag/jot/internal/command/module"
	p "github.com/AdamShannag/jot/internal/command/path"
	s "github.com/AdamShannag/jot/internal/command/suffix"
	"github.com/AdamShannag/jot/internal/io"
	"github.com/AdamShannag/jot/internal/makefile"
	"github.com/AdamShannag/jot/internal/template"
	"github.com/AdamShannag/jot/internal/types"
	"github.com/AdamShannag/jot/internal/types/tpls"
)

func NewRestService(specs *types.Specs, mk *makefile.Makefile, service *types.Service) {
	mk.InitMod(s.ServiceSuffix(service.Name))
	mk.GetGoModules(module.GoChi, module.GoChiCors, module.GoChiMiddleware)

	createRESTService(service.Name)
	createDockerFile(service.Name)
	if len(service.Endpoints) > 0 {
		endpoint.CreateAll(service.Name, service.Endpoints)
	}

	// update specs file
	specs.Services = append(specs.Services, *service)
}

func createRESTService(name string) {
	createDirectories(name)
	createRest(s.ServiceSuffix(name))

	log.Info("Service", log.CREATED)
}

func createRest(service string) {
	// fill in api data
	apiData := tpls.Api{}
	apiData.AddModules(module.GoChi, module.GoChiCors, module.GoChiMiddleware)
	template.Create(p.ApiTpl, p.Path(p.ApiDirPath, service), p.ApiFileName, apiData)

	// create handlers and middelwares
	io.ToDirs(p.Path(p.HandlerDirPath, service))
	io.ToDirs(p.Path(p.MiddelwareDirPath, service))
}
