package endpoint

import "github.com/AdamShannag/jot/v2/types/model"

func (b *Endpoint) New(name string) *Endpoint {
	b.actions.modifications = append(b.actions.modifications, func(service *model.Endpoint) {
		service.Name = name
	})
	return b
}

func (b *Endpoint) Build() *model.Endpoint {
	q := model.Endpoint{}
	for _, action := range b.actions.modifications {
		action(&q)
	}
	return &q
}
