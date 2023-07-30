package template

type Template string

const (
	MOD               Template = "go.mod.gotpl"
	API               Template = "api.go.gotpl"
	MAIN              Template = "main.go.gotpl"
	Endpoint          Template = "endpoint.go.gotpl"
	DefaultMiddleware Template = "middleware.go.gotpl"
	Logger            Template = "logger.go.gotpl"
	LoggerMiddleware  Template = "logger_middleware.go.gotpl"
)
