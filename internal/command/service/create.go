package service

import (
	"fmt"
	"strings"

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

func New(service string, isRest bool, endpoints []string, specs *types.Specs, port int) {
	srv := service
	s.ServiceSuffix(&srv)

	// create service directoies
	createService(service, isRest)

	// create .dockerfile
	template.Create(p.DockerImageTpl, p.Path(p.DockerImagePath, srv), s.DockerfileSuffix(srv),
		tpls.Docker{AppName: s.AppSuffix(service), AppPath: p.Path(p.AppPath, s.AppSuffix(service))})

	// create endpoints directories if rest enabled
	if isRest && len(endpoints) > 0 {
		endpoint.CreateAll(service, endpoints)
	}

	// update specs file
	specs.Services = append(specs.Services, *types.NewService(service, port, endpoints))
}
func createService(name string, rest bool) {
	service := name
	name = strings.Replace(name, "-service", "", 1)
	s.ServiceSuffix(&service)

	mk := makefile.New(p.Path(p.GoModPath, service), 10)
	// go mod init
	mk.InitMod(service)

	defer mk.Build()

	// create service directories cmd, bin, api
	io.ToDirs(fmt.Sprintf(p.MainDirPath, service, name))
	io.ToDirs(p.Path(p.BinDirPath, service))
	io.ToDirs(p.Path(p.ApiDirPath, service))

	// create api directories if rest handler, middleware
	if rest {
		// go get
		mk.GetGoModules(module.GoChi, module.GoChiCors, module.GoChiMiddleware)
		createRest(service)
	}

	// go tidy
	mk.GoTidy()
	// go fmt
	mk.GoFmt()

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
