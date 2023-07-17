package service

import (
	"github.com/AdamShannag/jot/v2/types/model"
)

func NewBuilder() *Service {
	return &Service{&actions{}}
}

func (b *Service) Name(name string) *Service {
	b.actions.modifications = append(b.actions.modifications, func(service *model.Service) {
		service.Name = name
	})
	return b
}

func (b *Service) Port(port int) *Service {
	b.actions.modifications = append(b.actions.modifications, func(service *model.Service) {
		service.Port = port
	})
	return b
}

func (b *Service) Endpoints(endpoints []model.Endpoint) *Service {
	b.actions.modifications = append(b.actions.modifications, func(service *model.Service) {
		service.Endpoints = append(service.Endpoints, endpoints...)
	})
	return b
}

func (b *Service) Middlewares(middlewares []model.Middleware) *Service {
	b.actions.modifications = append(b.actions.modifications, func(service *model.Service) {
		service.Middlewares = append(service.Middlewares, middlewares...)
	})
	return b
}

func (b *Service) Build() model.Service {
	q := model.Service{}
	for _, action := range b.actions.modifications {
		action(&q)
	}
	return q
}
