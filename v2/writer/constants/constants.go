package constants

import (
	d "github.com/AdamShannag/jot/v2/writer/directory"
	f "github.com/AdamShannag/jot/v2/writer/file"
)

const (
	API_DIR             string = "api"
	API_ENDPOINTS_DIR   string = "endpoints"
	API_MIDDLEWARES_DIR string = "middlewares"
	BIN_DIR             string = "bin"
	CMD_DIR             string = "cmd"
	DEPLOY_DIR          string = "deploy"
	PKG_DIR             string = "pkg"
)

func DefaultServiceStructure() []*d.Directory {
	dirs := []*d.Directory{}

	dfMap := map[string][]f.File{
		API_DIR:    {f.NewApiFile("api", nil)},
		BIN_DIR:    {},
		CMD_DIR:    {},
		DEPLOY_DIR: {},
		PKG_DIR:    {},
	}

	for k, v := range dfMap {
		dirs = append(dirs, d.NewDefaultDirectory(k, nil, v...))
	}

	return dirs
}
