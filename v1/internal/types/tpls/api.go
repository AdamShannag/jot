package tpls

type Api struct {
	Handlers []Handler
	Imports  []string
}

func (a *Api) AddModules(names ...string) {
	a.Imports = append(a.Imports, names...)
}
