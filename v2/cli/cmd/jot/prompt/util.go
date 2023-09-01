package prompt

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"slices"

	"github.com/AdamShannag/jot/v2/api/endpoint"
	"github.com/AdamShannag/jot/v2/api/middleware"
	"github.com/AdamShannag/jot/v2/api/service"
	"github.com/AdamShannag/jot/v2/types/model"
	f "github.com/AdamShannag/jot/v2/writer/fs"
	"github.com/spf13/afero"
)

func (p *PrompterImpl) servicesSlice() (srv []string) {
	for _, s := range p.services {
		srv = append(srv, s.Name)
	}
	return
}

func (p *PrompterImpl) pathExists(path string) bool {
	ok, err := afero.Exists(f.Get(), path)
	if err != nil {
		log.Fatalf("Error occured: %v", err)
	}
	return ok
}

func (p *PrompterImpl) search(path string, d fs.DirEntry, _ error) error {
	current, err := d.Info()

	if err != nil {
		return err
	}

	if current.IsDir() {
		if p.pathExists(filepath.Join(path, "cmd")) {
			filepath.WalkDir(filepath.Join(path, "cmd"), func(path string, d fs.DirEntry, _ error) error {
				current, err := d.Info()
				if err != nil {
					return err
				}
				if current.Name() == "main.go" {
					servicePath := strings.Split(strings.ReplaceAll(filepath.Dir(path), "\\", "/"), "/")
					serviceRoot := slices.Delete[[]string](servicePath, len(servicePath)-2, len(servicePath))
					if p.projectPath == "" {
						p.projectPath = filepath.Join(slices.Delete[[]string](serviceRoot, len(serviceRoot)-1, len(serviceRoot))...)
						if p.projectPath == "" {
							p.projectPath = "./"
						}
					}
					endpointsDir := filepath.Join(filepath.Join(serviceRoot...), "api/endpoints")
					middlewaresDir := filepath.Join(filepath.Join(serviceRoot...), "api/middlewares")

					endpointEntries, _ := os.ReadDir(endpointsDir)
					middlewareEntries, _ := os.ReadDir(middlewaresDir)

					var endpoints []model.Endpoint
					var middlewares []model.Middleware

					for _, e := range endpointEntries {
						endpoints = append(endpoints, endpoint.NewBuilder().Name(e.Name()).Build())
					}

					for _, m := range middlewareEntries {
						middlewares = append(middlewares, middleware.NewBuilder().Name(m.Name()).Build())
					}

					p.services = append(p.services, service.NewBuilder().
						Name(servicePath[len(servicePath)-1]).
						Endpoints(endpoints).
						Middlewares(middlewares).
						Build())
				}
				return nil
			})
		}
	}
	return nil
}

func (p *PrompterImpl) walkServices() {
	err := filepath.WalkDir(".", p.search)
	if err != nil {
		log.Fatal(err)
	}
}
