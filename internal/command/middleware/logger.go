package middleware

import (
	"github.com/AdamShannag/jot/internal/config"
	"github.com/AdamShannag/jot/internal/makefile"
	"github.com/AdamShannag/jot/internal/types"
)

func Logger(specs *types.Specs, service string, mk *makefile.Makefile) error {
	if ok, i := types.IsExistingService(specs.Services, service); ok {
		if !existingMiddleware(specs, i, config.LOGGER_MID) {

		}
	}
	return nil
}
