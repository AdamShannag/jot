package tpls

type Crud struct {
	PackageName  string
	EndpointName string
	Imports      []string
}

func (h *Crud) AddModules(names ...string) {
	h.Imports = append(h.Imports, names...)
}
