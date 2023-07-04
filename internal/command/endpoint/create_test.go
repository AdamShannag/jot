package endpoint

import (
	"fmt"
	"log"
	"testing"

	"github.com/AdamShannag/jot/internal/command/path"
	"github.com/AdamShannag/jot/internal/io"
	"github.com/AdamShannag/jot/internal/types"
	"github.com/spf13/afero"
)

func Test_CreateAll_WithoutCrudFile(t *testing.T) {

	appFs := io.SwitchToMemMap()

	endpoints := []string{"users", "posts"}

	newServiceName := "user-service"

	CreateAll(newServiceName, endpoints, false)

	paths := []string{
		fmt.Sprintf(path.HandlerPath, newServiceName, endpoints[0]),
		fmt.Sprintf(path.HandlerPath, newServiceName, endpoints[1]),
	}

	for _, path := range paths {
		fileExists(t, &appFs, path)
	}
}

func Test_CreateAll_WithCrudFile(t *testing.T) {

	appFs := io.SwitchToMemMap()

	endpoints := []string{"users", "posts"}

	newServiceName := "user-service"

	CreateAll(newServiceName, endpoints, true)

	paths := []string{
		fmt.Sprintf(path.HandlerPath, newServiceName, endpoints[0]),
		fmt.Sprintf(path.HandlerPath, newServiceName, endpoints[1]),
		fullPath(path.CrudPath, newServiceName, endpoints[0], path.CrudFileName),
		fullPath(path.CrudPath, newServiceName, endpoints[1], path.CrudFileName),
	}

	for _, path := range paths {
		fileExists(t, &appFs, path)
	}
}

func Test_UpdateAll(t *testing.T) {

	appFs := io.SwitchToMemMap()

	endpoints := []string{"users"}
	services := []types.Service{}

	newServiceName := "user-service"

	services = append(services, *types.NewService(newServiceName, 8081, endpoints, []string{}))

	s := types.NewSpecs("test", services, nil)

	endpoints = append(endpoints, "posts")
	UpdateAll(endpoints, s, 0, newServiceName, false)

	paths := []string{
		fmt.Sprintf(path.HandlerPath, newServiceName, endpoints[1]),
	}

	for _, path := range paths {
		fileExists(t, &appFs, path)
	}
}

func Test_AddCrudFile_WhenUpdateAll(t *testing.T) {

	appFs := io.SwitchToMemMap()

	endpoints := []string{"users"}
	services := []types.Service{}

	newServiceName := "user-service"

	services = append(services, *types.NewService(newServiceName, 8081, endpoints, []string{}))

	s := types.NewSpecs("test", services, nil)

	endpoints = append(endpoints, "posts")
	UpdateAll(endpoints, s, 0, newServiceName, true)

	paths := []string{
		fullPath(path.CrudPath, newServiceName, endpoints[0], path.CrudFileName),
		fullPath(path.CrudPath, newServiceName, endpoints[1], path.CrudFileName),
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

func fullPath(format, folder1, folder2, filename string) string {
	return fmt.Sprintf("%s%s", fmt.Sprintf(format, folder1, folder2), filename)
}
