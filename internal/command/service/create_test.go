package service

import (
	"fmt"
	"log"
	"testing"

	"github.com/AdamShannag/jot/internal/command/path"
	"github.com/AdamShannag/jot/internal/command/suffix"
	"github.com/AdamShannag/jot/internal/io"
	"github.com/AdamShannag/jot/internal/types"
	"github.com/spf13/afero"
)

func Test_New(t *testing.T) {
	appFs := io.SwitchToMemMap()

	endpoints := []string{"users", "posts"}
	services := []types.Service{}

	newServiceName := "user-service"

	s := types.NewSpecs("test", services, nil)

	New(newServiceName, true, endpoints, s, 8082)

	paths := []string{
		fullPath(path.GoModPath, newServiceName, path.GoModFileName),
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
