package middleware

import (
	mid "github.com/AdamShannag/jot/internal/command/middleware"
	"github.com/AdamShannag/jot/internal/config"
	"github.com/AdamShannag/jot/internal/makefile"
	"github.com/AdamShannag/jot/internal/types"
)

func GetSupportedMiddlewares() map[string]func(specs *types.Specs, service string, mk *makefile.Makefile) error {
	return map[string]func(specs *types.Specs, service string, mk *makefile.Makefile) error{
		config.LOGGER_MID: mid.Logger,
	}
}
