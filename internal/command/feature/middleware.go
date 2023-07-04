package feature

import (
	"errors"

	middle "github.com/AdamShannag/jot/internal/command/middleware"
	config "github.com/AdamShannag/jot/internal/config/middleware"
	"github.com/AdamShannag/jot/internal/makefile"
	"github.com/AdamShannag/jot/internal/types"
)

func newMiddlewares(middelwares []string, rest bool) *middleware {
	return &middleware{
		middelwares: middelwares,
		rest:        rest,
	}
}

func (m *middleware) Build(specs *types.Specs, service string, mk *makefile.Makefile) error {
	if !m.rest && len(m.middelwares) > 0 {
		return errors.New("--middleware flag is specified but --rest flag is not")
	}
	var error error = nil
	for _, mid := range m.middelwares {
		if f, ok := config.GetSupportedMiddlewares()[mid]; ok {
			if err := f(specs, service, mk); err != nil {
				error = err
				break
			}
		} else {
			if err := middle.Defualt(mid, specs, service, mk); err != nil {
				error = err
				break
			}
		}
	}

	mk.GoTidy()
	mk.GoFmt()

	return error
}
