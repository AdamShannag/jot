package project

import "github.com/AdamShannag/jot/v2/types/model"

func (b *Project) New(name string) *Project {
	b.actions.modifications = append(b.actions.modifications, func(service *model.Project) {
		service.Name = name
	})
	return b
}

func (b *Project) Build() *model.Project {
	q := model.Project{}
	for _, action := range b.actions.modifications {
		action(&q)
	}
	return &q
}
