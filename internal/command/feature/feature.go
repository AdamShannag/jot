package feature

import (
	"github.com/AdamShannag/jot/internal/makefile"
	"github.com/AdamShannag/jot/internal/types"
	"github.com/urfave/cli/v2"
)

func New(specs *types.Specs, mk *makefile.Makefile, cCtx *cli.Context) *feature {
	return &feature{
		Endpoint: endpoint{
			rest: newRestApi(cCtx.Bool("rest"), false),
			grpc: false,
		},
		Middleware: middleware{
			jwt:  false,
			rbac: false,
		},
		specs:     specs,
		mk:        mk,
		service:   cCtx.String("service"),
		port:      cCtx.Int("port"),
		endpoints: cCtx.StringSlice("endpoints"),
	}
}

func (f *feature) BuildREST() error {
	return f.Endpoint.rest.Build(f.specs, f.mk, f.service, f.port, f.endpoints)
}

func (f *feature) BuildGRPC() error {
	return nil
}
