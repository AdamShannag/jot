package tpls

type Api struct {
	Imports []string
}

func (a *Api) AddModules(names ...string) {
	a.Imports = append(a.Imports, names...)
}
