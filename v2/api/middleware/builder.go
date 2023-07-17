package middleware

import "github.com/AdamShannag/jot/v2/types/model"

func (b *Middleware) New(name string) *Middleware {
	b.actions.modifications = append(b.actions.modifications, func(service *model.Middleware) {
		service.Name = name
	})
	return b
}

func (b *Middleware) Build() *model.Middleware {
	q := model.Middleware{}
	for _, action := range b.actions.modifications {
		action(&q)
	}
	return &q
}
