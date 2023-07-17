package feature

import (
	"github.com/AdamShannag/jot/v1/internal/makefile"
	"github.com/AdamShannag/jot/v1/internal/types"
	"github.com/urfave/cli/v2"
)

func New(specs *types.Specs, mk *makefile.Makefile, cCtx *cli.Context) *feature {
	return &feature{
		Endpoint: endpoint{
			rest: newRestApi(cCtx.Bool("rest"), cCtx.Bool("crud")),
			grpc: false,
		},
		Middleware: newMiddlewares(cCtx.StringSlice("middlewares"), cCtx.Bool("rest")),
		specs:      specs,
		mk:         mk,
		service:    cCtx.String("service"),
		port:       cCtx.Int("port"),
		endpoints:  cCtx.StringSlice("endpoints"),
	}
}

func (f *feature) BuildREST() error {
	return f.Endpoint.rest.Build(f.specs, f.mk, f.service, f.port, f.endpoints)
}

func (f *feature) BuildGRPC() error {
	return nil
}

func (f *feature) BuildMiddlewares() error {
	return f.Middleware.Build(f.specs, f.service, f.mk)
}
