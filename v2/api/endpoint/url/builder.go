package url

import (
	"github.com/AdamShannag/jot/v2/types/model"
)

func NewBuilder() *Url {
	return &Url{&actions{}}
}

func (b *Url) Path(path string) *Url {
	b.actions.modifications = append(b.actions.modifications, func(url *model.Url) {
		url.Path = path
	})

	return b
}

func (b *Url) Handler(handler string) *Url {
	b.actions.modifications = append(b.actions.modifications, func(url *model.Url) {
		url.Handler = handler
	})
	return b
}

func (b *Url) Method(method string) *Url {
	b.actions.modifications = append(b.actions.modifications, func(url *model.Url) {
		url.Method = method
	})
	return b
}

func (b *Url) Build() model.Url {
	q := model.Url{}
	for _, action := range b.actions.modifications {
		action(&q)
	}
	return q
}
