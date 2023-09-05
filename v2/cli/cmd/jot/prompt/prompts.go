package prompt

import (
	"fmt"
	"log"
	"strconv"

	"github.com/AdamShannag/jot/v2/api/endpoint"
	"github.com/AdamShannag/jot/v2/api/middleware"
	"github.com/AdamShannag/jot/v2/types/model"
)

func (p *PrompterImpl) creationPrompt() string {
	if len(p.services) == 0 {
		_, choice := p.Select("Create New", SERVICE)
		return choice
	}
	_, choice := p.Select("Create New", SERVICE, ENDPOINT, MIDDLEWARE)
	return choice
}

func (p *PrompterImpl) servicePrompt() {
	var srv model.Service
	srv.Name = p.Prompt("Service Name", invalidStringValidator)
	port, err := strconv.ParseInt(p.Prompt("Service Port", invalidStringValidator), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	srv.Port = int(port)
	p.services = append(p.services, srv)
}

func (p *PrompterImpl) endpointPrompt() {
	i, _ := p.Select("Select service", p.servicesSlice()...)
	if i == len(p.services) {
		p.startPrompts()
		return
	}
	newEndpoint := endpoint.NewBuilder().Name(p.Prompt("New Endpoint Name", invalidStringValidator)).Build()
	p.services[i].Endpoints = append(p.services[i].Endpoints, newEndpoint)
}

func (p *PrompterImpl) middlewarePrompt() {
	i, _ := p.Select("Select service", p.servicesSlice()...)
	if i == len(p.services) {
		p.startPrompts()
		return
	}
	newMiddleware := middleware.NewBuilder().Name(p.Prompt("New Middleware Name", invalidStringValidator)).Build()
	p.services[i].Middlewares = append(p.services[i].Middlewares, newMiddleware)
}

func (p *PrompterImpl) projectPrompt() {
	if p.projectPath == "" {
		fmt.Println("project path not found!")
		p.projectPath = p.Prompt("Project Path", invalidStringValidator)
	}
}
