package service

import (
	"fmt"
	"strings"

	"github.com/AdamShannag/jot/internal/command/log"
	p "github.com/AdamShannag/jot/internal/command/path"
	s "github.com/AdamShannag/jot/internal/command/suffix"
	"github.com/AdamShannag/jot/internal/io"
	"github.com/AdamShannag/jot/internal/makefile"
	"github.com/AdamShannag/jot/internal/template"
	"github.com/AdamShannag/jot/internal/types"
	"github.com/AdamShannag/jot/internal/types/tpls"
)

func NewService(specs *types.Specs, mk *makefile.Makefile, service *types.Service) {
	mk.InitMod(s.ServiceSuffix(service.Name))

	createService(service.Name)
	createDockerFile(service.Name)

	specs.Services = append(specs.Services, *service)
}

func createService(name string) {
	createDirectories(name)

	log.Info("Service", log.CREATED)
}

func createDirectories(service string) {
	service = strings.Replace(service, "-service", "", 1)
	io.ToDirs(fmt.Sprintf(p.MainDirPath, s.ServiceSuffix(service), service))
	io.ToDirs(p.Path(p.BinDirPath, s.ServiceSuffix(service)))
	io.ToDirs(p.Path(p.ApiDirPath, s.ServiceSuffix(service)))
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
