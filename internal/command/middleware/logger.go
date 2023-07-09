package middleware

import (
	"github.com/AdamShannag/jot/internal/command/path"
	"github.com/AdamShannag/jot/internal/command/suffix"
	"github.com/AdamShannag/jot/internal/config"
	"github.com/AdamShannag/jot/internal/makefile"
	"github.com/AdamShannag/jot/internal/template"
	"github.com/AdamShannag/jot/internal/types"
)

func Logger(specs *types.Specs, service string, mk *makefile.Makefile) error {
	if ok, i := types.IsExistingService(specs.Services, service); ok {
		if !existingMiddleware(specs, i, config.LOGGER_MID) {
			template.Create(
				path.ZerologMiddlewarePkgTpl,
				path.Path(path.ZerologPkgPath, suffix.ServiceSuffix(service)),
				suffix.GoSuffix(path.ZerologMiddleware),
				nil,
			)
		}
	}
	return nil
}
