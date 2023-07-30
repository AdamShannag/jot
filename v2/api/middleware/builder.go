package middleware

import "github.com/AdamShannag/jot/v2/types/model"

func NewBuilder() *Middleware {
	return &Middleware{&actions{}}
}

func (b *Middleware) Name(name string) *Middleware {
	b.actions.modifications = append(b.actions.modifications, func(middleware *model.Middleware) {
		middleware.Name = name
	})
	return b
}

func (b *Middleware) Build() model.Middleware {
	q := model.Middleware{}
	for _, action := range b.actions.modifications {
		action(&q)
	}
	return q
}
