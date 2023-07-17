package tpls

type Middleware struct {
	MiddlewareName string
	Imports        []string
}

func (m *Middleware) AddModules(names ...string) {
	m.Imports = append(m.Imports, names...)
}
