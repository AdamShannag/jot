package project

import "github.com/AdamShannag/jot/v2/types/model"

func NewBuilder() *Project {
	return &Project{&actions{}}
}

func (b *Project) Name(name string) *Project {
	b.actions.modifications = append(b.actions.modifications, func(project *model.Project) {
		project.Name = name
	})
	return b
}

func (b *Project) Services(services []model.Service) *Project {
	b.actions.modifications = append(b.actions.modifications, func(project *model.Project) {
		project.Services = append(project.Services, services...)
	})
	return b
}

func (b *Project) Build() model.Project {
	q := model.Project{}
	for _, action := range b.actions.modifications {
		action(&q)
	}
	return q
}
