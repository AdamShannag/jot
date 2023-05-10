package feature

import (
	"errors"

	e "github.com/AdamShannag/jot/internal/command/endpoint"
	srv "github.com/AdamShannag/jot/internal/command/service"
	"github.com/AdamShannag/jot/internal/makefile"
	"github.com/AdamShannag/jot/internal/types"
)

func newRestApi(isRest, isCrud bool) *restAPI {
	return &restAPI{
		rest: isRest,
		crud: isCrud,
	}
}

func (r *restAPI) Build(specs *types.Specs, mk *makefile.Makefile, service string, port int, endpoints []string) error {
	if !r.rest && len(endpoints) > 0 {
		return errors.New("--endpoints flag is specified but --rest flag is not")
	}

	if ok, i := types.IsExistingService(specs.Services, service); ok {
		e.UpdateAll(endpoints, specs, i, service)
	} else {
		if r.rest {
			srv.NewRestService(specs, mk, &types.Service{Name: service, Port: port, Endpoints: endpoints})
		} else {
			srv.NewService(specs, mk, &types.Service{Name: service, Port: port, Endpoints: endpoints})
		}
	}

	mk.GoTidy()
	mk.GoFmt()

	return nil
}
