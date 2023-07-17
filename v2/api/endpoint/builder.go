package endpoint

import "github.com/AdamShannag/jot/v2/types/model"

func NewBuilder() *Endpoint {
	return &Endpoint{&actions{}}
}

func (b *Endpoint) Name(name string) *Endpoint {
	b.actions.modifications = append(b.actions.modifications, func(endpoint *model.Endpoint) {
		endpoint.Name = name
	})
	return b
}

func (b *Endpoint) Urls(urls []model.Url) *Endpoint {
	b.actions.modifications = append(b.actions.modifications, func(endpoint *model.Endpoint) {
		endpoint.Urls = append(endpoint.Urls, urls...)
	})
	return b
}

func (b *Endpoint) Build() model.Endpoint {
	q := model.Endpoint{}
	for _, action := range b.actions.modifications {
		action(&q)
	}
	return q
}
