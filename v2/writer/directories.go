package writer

import (
	"github.com/AdamShannag/jot/v2/types/model"
	"github.com/AdamShannag/jot/v2/writer/util"
)

// Creates a directory for each service and its related directories
func (w *writer) servicesDirectories() []*Dir {
	servicesDirs := []*Dir{}

	for _, service := range w.project.Services {
		servicesDirs = append(servicesDirs, w.serviceDirectory(&service))
	}

	return servicesDirs
}

// Creates a service directory with its releated directories
func (w *writer) serviceDirectory(srv *model.Service) *Dir {
	sd := &Dir{srv.Name, nil, serviceDFs()}

	for _, d := range sd.Dirs {
		if d.Name == CMD_DIR {
			d.Dirs = append(d.Dirs, &Dir{srv.Name, []File{{MAIN_FILE, "", "", nil}}, nil})
		}
		if d.Name == API_DIR {
			d.Dirs = append(d.Dirs, w.endpointsDirectory(srv.Endpoints))
			d.Dirs = append(d.Dirs, w.middlewaresDirectory(srv.Middlewares))
		}
	}

	return sd
}

// Creates endpoints directory with its related directories
func (w *writer) endpointsDirectory(endpoints []model.Endpoint) *Dir {
	var ed = &Dir{API_ENDPOINTS_DIR, nil, nil}

	for _, endpoint := range endpoints {
		ed.Dirs = append(ed.Dirs, &Dir{
			endpoint.Name,
			[]File{apiFile(endpoint.Name, "endpoint.go.gotpl")},
			nil},
		)
	}

	return ed
}

// Creates middlewares directory with its related directories
func (w *writer) middlewaresDirectory(middlewares []model.Middleware) *Dir {
	var md = &Dir{API_MIDDLEWARES_DIR, nil, nil}

	for _, middleware := range middlewares {
		md.Dirs = append(md.Dirs, &Dir{
			middleware.Name,
			[]File{apiFile(middleware.Name, "middleware.go.gotpl")},
			nil})
	}

	return md
}

func apiFile(name string, tpl string) File {
	return File{
		Name: name,
		Ext:  ".go",
		Tpl:  tpl,
		Data: map[string]string{
			"PackageName": name,
			"Name":        util.TitleCase(name),
		},
	}
}
