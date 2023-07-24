package writer

import (
	"github.com/AdamShannag/jot/v2/types/model"
	d "github.com/AdamShannag/jot/v2/writer/directory"
	f "github.com/AdamShannag/jot/v2/writer/file"
)

// Creates a directory for each service and its related directories
func (w *projectWriter) constructServices() []*d.Directory {
	servicesDirs := []*d.Directory{}

	for _, service := range w.project.Services {
		servicesDirs = append(servicesDirs, w.constructOneService(&service))
	}

	return servicesDirs
}

// Creates a service directory with its releated directories
func (w *projectWriter) constructOneService(srv *model.Service) *d.Directory {
	sd := d.NewDefaultDirectory(srv.Name, defaultServiceStructure())

	for _, dir := range sd.Directories {
		if dir.Name == CMD_DIR {
			dir.Directories = append(dir.Directories, d.NewDefaultDirectory(srv.Name, nil, f.NewMainFile("main", nil)))
		}
		if dir.Name == API_DIR {
			dir.Directories = append(dir.Directories, w.constructEndpoints(srv.Endpoints))
			dir.Directories = append(dir.Directories, w.constructMiddlewares(srv.Middlewares))
		}
	}

	return sd
}

// Creates endpoints directory with its related directories
func (w *projectWriter) constructEndpoints(endpoints []model.Endpoint) *d.Directory {
	var ed = d.NewDefaultDirectory(API_ENDPOINTS_DIR, nil)

	for _, endpoint := range endpoints {
		ed.Directories = append(ed.Directories, d.NewDefaultDirectory(endpoint.Name, nil, f.NewEndpointFile(endpoint.Name, nil)))
	}

	return ed
}

// Creates middlewares directory with its related directories
func (w *projectWriter) constructMiddlewares(middlewares []model.Middleware) *d.Directory {
	var md = d.NewDefaultDirectory(API_MIDDLEWARES_DIR, nil)

	for _, middleware := range middlewares {
		md.Directories = append(md.Directories, d.NewDefaultDirectory(middleware.Name, nil, f.NewDefaultMiddlewareFile(middleware.Name, nil)))
	}

	return md
}
