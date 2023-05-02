package path

import "fmt"

const (
	// jot.yaml
	JotRelPath = "./"
	JotFile    = "jot.yaml"

	// directories
	ProjectDirPath    = "%s/%s/"
	MainDirPath       = "./%s/cmd/%s/"
	BinDirPath        = "./%s/bin/"
	AppPath           = "../../bin/%s"
	HandlerDirPath    = "./%s/api/handler/"
	MiddelwareDirPath = "./%s/api/middleware/"

	// go.mod
	GoModFileName = "go.mod"
	GoModTpl      = "go.mod.gotpl"
	GoModPath     = "./%s/"

	// .dockerfile
	DockerImageFileExt = ".dockerfile"
	DockerImageTpl     = "image.gotpl"
	DockerImagePath    = "./%s/deploy/image/"

	// api.go
	ApiDirPath  = "./%s/api/"
	ApiFileName = "api.go"
	ApiTpl      = "api.go.gotpl"

	// handlers.go
	HandlerTpl  = "handler.go.gotpl"
	HandlerPath = "./%s/api/handler/%s/"
)

func Path(format string, arg string) string {
	return fmt.Sprintf(format, arg)
}
