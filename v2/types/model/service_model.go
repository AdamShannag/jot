package model

type Service struct {
	Name        string       `json:"name"`
	Port        int          `json:"port"`
	Endpoints   []Endpoint   `json:"endpoints,omitempty"`
	Middlewares []Middleware `json:"middlewares,omitempty"`
}
