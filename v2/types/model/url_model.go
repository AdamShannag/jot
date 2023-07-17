package model

type Url struct {
	Path    string `json:"path"`
	Handler string `json:"handler"`
	Method  string `json:"method"`
}
