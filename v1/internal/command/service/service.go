package service

import (
	"fmt"
	"strings"

	"github.com/AdamShannag/jot/v1/internal/command/module"
	p "github.com/AdamShannag/jot/v1/internal/command/path"
	s "github.com/AdamShannag/jot/v1/internal/command/suffix"
	"github.com/AdamShannag/jot/v1/internal/io"
	"github.com/AdamShannag/jot/v1/internal/makefile"
	"github.com/AdamShannag/jot/v1/internal/template"
	"github.com/AdamShannag/jot/v1/internal/types"
	"github.com/AdamShannag/jot/v1/internal/types/tpls"
)

func NewService(specs *types.Specs, mk *makefile.Makefile, service *types.Service) {
	mk.InitMod(s.ServiceSuffix(service.Name))
	mk.GetGoModules(module.Zerolog, module.ZerologXID, module.ZerologPkgerrors, module.Lumberjack)

	createService(service.Name)
	createDockerFile(service.Name)

	specs.Services = append(specs.Services, *service)
}

func createService(name string) {
	createDirectories(name)
	createZerologPkg(name)
}

func createDirectories(service string) {
	service = strings.Replace(service, "-service", "", 1)
	io.ToDirs(fmt.Sprintf(p.MainDirPath, s.ServiceSuffix(service), service))
	io.ToDirs(p.Path(p.BinDirPath, s.ServiceSuffix(service)))
	io.ToDirs(p.Path(p.ApiDirPath, s.ServiceSuffix(service)))
	io.ToDirs(p.Path(p.PkgDirPath, s.ServiceSuffix(service)))
}

func createDockerFile(service string) {
	template.Create(
		p.DockerImageTpl,
		p.Path(p.DockerImagePath,
			s.ServiceSuffix(service)),
		s.DockerfileSuffix(s.ServiceSuffix(service)),
		tpls.Docker{
			AppName: s.AppSuffix(service),
			AppPath: p.Path(p.AppPath, s.AppSuffix(service)),
		},
	)
}
