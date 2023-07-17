package tpls

type Handler struct {
	PackageName  string
	EndpointName string
	Crud         bool
	Imports      []string
}

func (h *Handler) AddModules(names ...string) {
	h.Imports = append(h.Imports, names...)
}
