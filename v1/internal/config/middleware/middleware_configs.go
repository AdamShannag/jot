package middleware

import (
	mid "github.com/AdamShannag/jot/v1/internal/command/middleware"
	"github.com/AdamShannag/jot/v1/internal/config"
	"github.com/AdamShannag/jot/v1/internal/makefile"
	"github.com/AdamShannag/jot/v1/internal/types"
)

func GetSupportedMiddlewares() map[string]func(specs *types.Specs, service string, mk *makefile.Makefile) error {
	return map[string]func(specs *types.Specs, service string, mk *makefile.Makefile) error{
		config.LOGGER_MID: mid.Logger,
	}
}
