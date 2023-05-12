package endpoint

import (
	"fmt"

	"github.com/AdamShannag/jot/internal/command/log"
	"github.com/AdamShannag/jot/internal/command/module"
	p "github.com/AdamShannag/jot/internal/command/path"
	s "github.com/AdamShannag/jot/internal/command/suffix"
	"github.com/AdamShannag/jot/internal/template"
	"github.com/AdamShannag/jot/internal/types"
	"github.com/AdamShannag/jot/internal/types/tpls"
)

func UpdateAll(endpoints []string, specs *types.Specs, i int, service string, withCrud bool) {
	newEps := []string{}

	// create endpoints if new, ignore if existing
	for _, ep := range endpoints {
		if !s.Contains(specs.Services[i].Endpoints, ep) {
			New(service, ep, withCrud)
			newEps = append(newEps, ep)
		} else {
			log.Info("Endpoint", log.IGNORED)
		}
	}

	// update specs file
	specs.Services[0].Endpoints = append(specs.Services[0].Endpoints, newEps...)
}

func CreateAll(service string, endpoints []string, withCrud bool) {
	for _, e := range endpoints {
		New(s.ServiceSuffix(service), e, withCrud)
	}
}

func New(service string, endpoint string, withCrud bool) {
	newEndpoints(service, endpoint, withCrud)
	if withCrud {
		newCrud(service, endpoint)
	}
}

func newEndpoints(service string, endpoint string, withCrud bool) {
	handlerData := tpls.Handler{
		EndpointName: s.TitleCase(endpoint),
		PackageName:  s.LowerCase(endpoint),
		Crud:         withCrud,
	}
	handlerData.AddModules(module.GoChi)
	// create handlers files
	template.Create(p.HandlerTpl, fmt.Sprintf(p.HandlerPath, s.ServiceSuffix(service), endpoint), s.GoSuffix(endpoint), handlerData)
}

func newCrud(service string, endpoint string) {
	handlerData := tpls.Crud{
		EndpointName: s.TitleCase(endpoint),
		PackageName:  s.LowerCase(endpoint),
	}
	handlerData.AddModules(module.NetHttp)

	// create crud files
	template.Create(p.CrudTpl, fmt.Sprintf(p.CrudPath, s.ServiceSuffix(service), endpoint), p.CrudFileName, handlerData)
}
