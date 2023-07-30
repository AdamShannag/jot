package constructor

import (
	m "github.com/AdamShannag/jot/v2/types/model"
	c "github.com/AdamShannag/jot/v2/writer/constants"
	d "github.com/AdamShannag/jot/v2/writer/directory"
	f "github.com/AdamShannag/jot/v2/writer/file"
)

type projectDirectoryConstructor struct {
	project m.Project
}

func NewProjectDirectoryConstructor(p m.Project) *projectDirectoryConstructor {
	return &projectDirectoryConstructor{p}
}

func (pdc *projectDirectoryConstructor) Construct() *d.Directory {
	return d.NewDefaultDirectory(pdc.project.Name, pdc.services())
}

// Creates a directory for each service and its related directories
func (pdc *projectDirectoryConstructor) services() []*d.Directory {
	servicesDirs := []*d.Directory{}

	for _, service := range pdc.project.Services {
		servicesDirs = append(servicesDirs, pdc.service(&service))
	}

	return servicesDirs
}

// Creates a service directory with its releated directories
func (pdc *projectDirectoryConstructor) service(srv *m.Service) *d.Directory {
	sd := d.NewDefaultDirectory(srv.Name, nil, f.NewModFile(srv.Name))

	sd.InsertAt(sd.Name,
		d.NewDefaultDirectory(c.API_DIR, nil, f.NewApiFile()),
		d.NewDefaultDirectory(c.BIN_DIR, nil),
		d.NewDefaultDirectory(c.CMD_DIR, nil),
		d.NewDefaultDirectory(c.DEPLOY_DIR, nil),
		d.NewDefaultDirectory(c.PKG_DIR, nil),
	)

	sd.InsertAt(c.API_DIR,
		pdc.endpoints(srv.Endpoints),
		pdc.middlewares(srv.Middlewares),
	)

	sd.InsertAt(c.CMD_DIR, d.NewDefaultDirectory(srv.Name, nil, f.NewMainFile()))

	sd.InsertAt(c.PKG_DIR,
		d.NewDefaultDirectory(c.LOGGER_DIR, nil,
			f.NewLoggerFile(),
			f.NewRequestLoggerFile(),
		),
	)

	return sd
}

// Creates endpoints directory with its related directories
func (pdc *projectDirectoryConstructor) endpoints(endpoints []m.Endpoint) *d.Directory {
	var ed = d.NewDefaultDirectory(c.API_ENDPOINTS_DIR, nil)

	for _, endpoint := range endpoints {
		ed.Directories = append(ed.Directories, d.NewDefaultDirectory(endpoint.Name, nil, f.NewEndpointFile(endpoint.Name, nil)))
	}

	return ed
}

// Creates middlewares directory with its related directories
func (pdc *projectDirectoryConstructor) middlewares(middlewares []m.Middleware) *d.Directory {
	var md = d.NewDefaultDirectory(c.API_MIDDLEWARES_DIR, nil)

	for _, middleware := range middlewares {
		md.Directories = append(md.Directories, d.NewDefaultDirectory(middleware.Name, nil, f.NewMiddlewareFile(middleware.Name, nil)))
	}

	return md
}
