package feature

import (
	"fmt"
	"log"
	"testing"

	"github.com/AdamShannag/jot/internal/command/path"
	"github.com/AdamShannag/jot/internal/command/suffix"
	"github.com/AdamShannag/jot/internal/io"
	"github.com/AdamShannag/jot/internal/makefile"
	"github.com/AdamShannag/jot/internal/types"
	"github.com/spf13/afero"
)

func Test_BuildRESTWithoutCrud(t *testing.T) {
	appFs := io.SwitchToMemMap()

	endpoints := []string{"users", "posts"}
	services := []types.Service{}

	newServiceName := "user-service"

	s := types.NewSpecs("test", services, nil)
	mk := makefile.New(path.Path(path.GoModPath, newServiceName), 10)

	feat := feature{
		Endpoint: endpoint{
			rest: &restAPI{
				rest: true,
				crud: false,
			},
			grpc: false,
		},
		Middleware: middleware{
			jwt:  false,
			rbac: false,
		},
		specs:     s,
		mk:        mk,
		service:   newServiceName,
		port:      8082,
		endpoints: endpoints,
	}

	feat.BuildREST()

	paths := []string{
		fullPath(path.DockerImagePath, newServiceName, suffix.DockerfileSuffix(newServiceName)),
		fullPath(path.ApiDirPath, newServiceName, path.ApiFileName),
		fullPath(path.BinDirPath, newServiceName, ""),
		fullPath(path.HandlerDirPath, newServiceName, ""),
		fullPath(path.MiddelwareDirPath, newServiceName, ""),
		fmt.Sprintf(path.HandlerPath, newServiceName, endpoints[0]),
		fmt.Sprintf(path.HandlerPath, newServiceName, endpoints[1]),
	}

	for _, path := range paths {
		fileExists(t, &appFs, path)
	}
}

func Test_BuildRESTWithCrud(t *testing.T) {
	appFs := io.SwitchToMemMap()

	endpoints := []string{"users", "posts"}
	services := []types.Service{}

	newServiceName := "user-service"

	s := types.NewSpecs("test", services, nil)
	mk := makefile.New(path.Path(path.GoModPath, newServiceName), 10)

	feat := feature{
		Endpoint: endpoint{
			rest: &restAPI{
				rest: true,
				crud: true,
			},
			grpc: false,
		},
		Middleware: middleware{
			jwt:  false,
			rbac: false,
		},
		specs:     s,
		mk:        mk,
		service:   newServiceName,
		port:      8082,
		endpoints: endpoints,
	}

	feat.BuildREST()

	paths := []string{
		fullPath(path.DockerImagePath, newServiceName, suffix.DockerfileSuffix(newServiceName)),
		fullPath(path.ApiDirPath, newServiceName, path.ApiFileName),
		fullPath(path.BinDirPath, newServiceName, ""),
		fullPath(path.HandlerDirPath, newServiceName, ""),
		fullPath(path.MiddelwareDirPath, newServiceName, ""),
		fmt.Sprintf(path.HandlerPath, newServiceName, endpoints[0]),
		fmt.Sprintf(path.HandlerPath, newServiceName, endpoints[1]),
		fullPath2(path.CrudPath, newServiceName, endpoints[0], path.CrudFileName),
		fullPath2(path.CrudPath, newServiceName, endpoints[1], path.CrudFileName),
	}

	for _, path := range paths {
		fileExists(t, &appFs, path)
	}
}

func Test_AddCrudFile_WhenUpdateREST(t *testing.T) {
	appFs := io.SwitchToMemMap()

	endpoints := []string{"users"}
	services := []types.Service{}

	newServiceName := "user-service"

	s := types.NewSpecs("test", services, nil)
	mk := makefile.New(path.Path(path.GoModPath, newServiceName), 10)

	feat := feature{
		Endpoint: endpoint{
			rest: &restAPI{
				rest: true,
				crud: true,
			},
			grpc: false,
		},
		Middleware: middleware{
			jwt:  false,
			rbac: false,
		},
		specs:     s,
		mk:        mk,
		service:   newServiceName,
		port:      8082,
		endpoints: endpoints,
	}

	feat.BuildREST()

	paths := []string{
		fmt.Sprintf(path.HandlerPath, newServiceName, endpoints[0]),
	}

	for _, path := range paths {
		fileExists(t, &appFs, path)
	}

	feat.Endpoint.rest.crud = false
	feat.endpoints = append(feat.endpoints, "posts")

	feat.BuildREST()

	paths = []string{
		fmt.Sprintf(path.HandlerPath, newServiceName, feat.endpoints[0]),
		fmt.Sprintf(path.HandlerPath, newServiceName, feat.endpoints[1]),
		fullPath2(path.CrudPath, newServiceName, feat.endpoints[0], path.CrudFileName),
	}

	for _, path := range paths {
		fileExists(t, &appFs, path)
	}

	feat.Endpoint.rest.crud = true
	feat.endpoints = append(feat.endpoints, "comments")

	feat.BuildREST()

	paths = []string{
		fmt.Sprintf(path.HandlerPath, newServiceName, feat.endpoints[0]),
		fmt.Sprintf(path.HandlerPath, newServiceName, feat.endpoints[1]),
		fmt.Sprintf(path.HandlerPath, newServiceName, feat.endpoints[2]),
		fullPath2(path.CrudPath, newServiceName, feat.endpoints[0], path.CrudFileName),
		fullPath2(path.CrudPath, newServiceName, feat.endpoints[1], path.CrudFileName),
		fullPath2(path.CrudPath, newServiceName, feat.endpoints[2], path.CrudFileName),
	}

	for _, path := range paths {
		fileExists(t, &appFs, path)
	}

}

func fileExists(t *testing.T, appFs *afero.Fs, path string) {
	if ok, err := afero.Exists(*appFs, path); err != nil {
		log.Panic(err)
	} else if !ok {
		t.Errorf("Expected file to exist at %s, but it is not", path)
	}
}

func fullPath(format, folder, filename string) string {
	return fmt.Sprintf("%s%s", path.Path(format, folder), filename)
}

func fullPath2(format, folder1, folder2, filename string) string {
	return fmt.Sprintf("%s%s", fmt.Sprintf(format, folder1, folder2), filename)
}
