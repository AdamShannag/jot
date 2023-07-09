package path

import "fmt"

const (
	// jot.yaml
	JotRelPath = "./"
	JotFile    = "jot.yaml"

	// directories
	ProjectDirPath = "%s/%s/"
	MainDirPath    = "./%s/cmd/%s/"
	BinDirPath     = "./%s/bin/"
	AppPath        = "../../bin/%s"
	HandlerDirPath = "./%s/api/handler/"

	// go.mod
	GoModFileName = "go.mod"
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

	// middleware.go
	DefaultMiddlewareDirPath = "./%s/api/middleware/"
	DefaultMiddlewareTpl     = "default_middleware.go.gotpl"

	// crud.go
	CrudTpl      = "crud.go.gotpl"
	CrudPath     = "./%s/api/handler/%s/"
	CrudFileName = "crud.go"

	// pkg
	PkgDirPath = "./%s/pkg/"

	// pkg.logger.go
	ZerologPkgname          = "logger"
	ZerologMiddleware       = "request_logger"
	ZerologPkgTpl           = "pkg_zerolog.go.gotpl"
	ZerologMiddlewarePkgTpl = "zerolog_middleware.go.gotpl"
	ZerologPkgPath          = "./%s/pkg/logger/"
)

func Path(format string, arg string) string {
	return fmt.Sprintf(format, arg)
}
