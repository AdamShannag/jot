package middleware

import (
	"fmt"

	"github.com/AdamShannag/jot/internal/command/module"
	"github.com/AdamShannag/jot/internal/command/path"
	"github.com/AdamShannag/jot/internal/command/suffix"
	"github.com/AdamShannag/jot/internal/makefile"
	"github.com/AdamShannag/jot/internal/template"
	"github.com/AdamShannag/jot/internal/types"
	"github.com/AdamShannag/jot/internal/types/tpls"
)

func Defualt(name string, specs *types.Specs, service string, mk *makefile.Makefile) error {
	if ok, i := types.IsExistingService(specs.Services, service); ok {
		if !existingMiddleware(specs, i, name) {
			middlewareData := tpls.Middleware{
				MiddlewareName: suffix.TitleCase(name),
				Imports:        []string{module.NetHttp},
			}
			template.Create(path.DefaultMiddlewareTpl,
				fmt.Sprintf(path.DefaultMiddlewareDirPath, suffix.ServiceSuffix(service)),
				suffix.GoSuffix(suffix.LowerCase(name)), middlewareData)
		}
	}
	return nil
}
