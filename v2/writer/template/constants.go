package template

type Template string

const (
	API               Template = "api.go.gotpl"
	MAIN              Template = "main.go.gotpl"
	Endpoint          Template = "endpoint.go.gotpl"
	DefaultMiddleware Template = "middleware.go.gotpl"
)
