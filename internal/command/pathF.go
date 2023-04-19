package command

import (
	"fmt"
)

const (
	// jot.yaml
	jotRelPath = "./"
	jotFile    = "jot.yaml"

	// directories
	projectDirPath    = "%s/%s/"
	mainDirPath       = "./%s/cmd/%s/"
	binDirPath        = "./%s/bin/"
	appPath           = "../../bin/%s"
	handlerDirPath    = "./%s/api/handler/"
	middelwareDirPath = "./%s/api/middleware/"

	// go.mod
	goModFileName = "go.mod"
	goModTpl      = "go.mod.gotpl"
	goModPath     = "./%s/"

	// .dockerfile
	dockerImageFileExt = ".dockerfile"
	dockerImageTpl     = "image.gotpl"
	dockerImagePath    = "./%s/deploy/image/"

	// api.go
	apiDirPath  = "./%s/api/"
	apiFileName = "api.go"
	apiTpl      = "api.go.gotpl"

	// handlers.go
	handlerTpl  = "handler.go.gotpl"
	handlerPath = "./%s/api/handler/%s/"
)

func path(format string, arg string) string {
	return fmt.Sprintf(format, arg)
}

func dockerfile(filename string) string {
	return fmt.Sprintf("%s%s", filename, dockerImageFileExt)

}

func appSuffix(s string) string {
	return fmt.Sprintf("%sApp", s)

}

func goSuffix(s string) string {
	return fmt.Sprintf("%s.go", s)
}
