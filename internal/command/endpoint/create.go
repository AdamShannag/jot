package endpoint

import (
	"fmt"

	"github.com/AdamShannag/jot/internal/command/log"
	p "github.com/AdamShannag/jot/internal/command/path"
	s "github.com/AdamShannag/jot/internal/command/suffix"
	"github.com/AdamShannag/jot/internal/template"
	"github.com/AdamShannag/jot/internal/types"
	"github.com/AdamShannag/jot/internal/types/tpls"
)

func UpdateAll(endpoints []string, specs *types.Specs, i int, service string) {
	newEps := []string{}

	// create endpoints if new, ignore if existing
	for _, ep := range endpoints {
		if !s.Contains(specs.Services[i].Endpoints, ep) {
			New(service, ep)
			newEps = append(newEps, ep)
		} else {
			log.Info("Endpoint", log.IGNORED)
		}
	}

	// update specs file
	specs.Services[0].Endpoints = append(specs.Services[0].Endpoints, newEps...)
}

func New(service string, endpoint string) {
	s.ServiceSuffix(&service)

	// create handlers files
	template.Create(p.HandlerTpl, fmt.Sprintf(p.HandlerPath, service, endpoint), s.GoSuffix(endpoint),
		tpls.Handler{EndpointName: s.TitleCase(endpoint)})
}

func CreateAll(service string, endpoints []string) {
	s.ServiceSuffix(&service)
	for _, e := range endpoints {
		New(service, e)
	}
}
