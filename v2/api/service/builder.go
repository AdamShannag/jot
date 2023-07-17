package service

import (
	"github.com/AdamShannag/jot/v2/types/model"
)

func (b *Service) New(name string) *Service {
	b.actions.modifications = append(b.actions.modifications, func(service *model.Service) {
		service.Name = name
	})
	return b
}

func (b *Service) Build() *model.Service {
	q := model.Service{}
	for _, action := range b.actions.modifications {
		action(&q)
	}
	return &q
}
