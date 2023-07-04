package middleware

import (
	"github.com/AdamShannag/jot/internal/command/log"
	"github.com/AdamShannag/jot/internal/command/suffix"
	"github.com/AdamShannag/jot/internal/types"
)

func existingMiddleware(specs *types.Specs, index int, name string) bool {
	if !suffix.Contains(specs.Services[index].Middlewares, name) {
		specs.Services[index].Middlewares = append(specs.Services[index].Middlewares, name)
		return false
	}
	log.Info(name, log.IGNORED)

	return true
}
