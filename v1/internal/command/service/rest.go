package service

import (
	"fmt"

	"github.com/AdamShannag/jot/v1/internal/command/endpoint"
	"github.com/AdamShannag/jot/v1/internal/command/module"
	p "github.com/AdamShannag/jot/v1/internal/command/path"
	s "github.com/AdamShannag/jot/v1/internal/command/suffix"
	"github.com/AdamShannag/jot/v1/internal/io"
	"github.com/AdamShannag/jot/v1/internal/makefile"
	"github.com/AdamShannag/jot/v1/internal/template"
	"github.com/AdamShannag/jot/v1/internal/types"
	"github.com/AdamShannag/jot/v1/internal/types/tpls"
)

func NewRestService(specs *types.Specs, mk *makefile.Makefile, service *types.Service, withCrud bool) {
	mk.InitMod(s.ServiceSuffix(service.Name))
	mk.GetGoModules(module.GoChi, module.GoChiCors, module.GoChiMiddleware)
	mk.GetGoModules(module.Zerolog, module.ZerologXID, module.ZerologPkgerrors, module.Lumberjack)

	createRESTService(service)
	createDockerFile(service.Name)
	if len(service.Endpoints) > 0 {
		endpoint.CreateAll(service.Name, service.Endpoints, withCrud)
	}

	// update specs file
	specs.Services = append(specs.Services, *service)
}

func createRESTService(service *types.Service) {
	createDirectories(service.Name)
	createZerologPkg(service.Name)
	createRest(service)
}

func createRest(service *types.Service) {
	serviceName := s.ServiceSuffix(service.Name)
	// fill in api data
	handlers, imports := extractHandlers(service)
	apiData := tpls.Api{
		Handlers: handlers,
	}
	apiData.AddModules(module.GoChi, module.GoChiCors, module.GoChiMiddleware)
	apiData.AddModules(imports...)
	template.Create(p.ApiTpl, p.Path(p.ApiDirPath, serviceName), p.ApiFileName, apiData)

	// create handlers and middelwares
	io.ToDirs(p.Path(p.HandlerDirPath, serviceName))
	io.ToDirs(p.Path(p.DefaultMiddlewareDirPath, serviceName))
}

func extractHandlers(service *types.Service) ([]tpls.Handler, []string) {
	var handlers []tpls.Handler
	var imports []string
	for _, ep := range service.Endpoints {
		handler := tpls.Handler{
			EndpointName: s.TitleCase(ep),
			PackageName:  s.LowerCase(ep),
		}
		imports = append(imports, importFormat(s.ServiceSuffix(service.Name), handler))
		handlers = append(handlers, handler)
	}
	return handlers, imports
}

func importFormat(service string, handler tpls.Handler) string {
	handlerPath := handlerPathFormat()
	return fmt.Sprintf(handlerPath, service, handler.PackageName)
}

func handlerPathFormat() string {
	return p.HandlerPath[len("./") : len(p.HandlerPath)-1]
}
