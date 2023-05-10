package feature

import (
	"github.com/AdamShannag/jot/internal/makefile"
	"github.com/AdamShannag/jot/internal/types"
)

type feature struct {
	Endpoint   endpoint
	Middleware middleware
	specs      *types.Specs
	mk         *makefile.Makefile
	service    string
	port       int
	endpoints  []string
}

type endpoint struct {
	rest *restAPI
	grpc bool
}

type restAPI struct {
	rest bool
	crud bool
}

type middleware struct {
	jwt  bool
	rbac bool
}
