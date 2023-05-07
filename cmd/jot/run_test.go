package main

import (
	"os"
	"testing"

	f "github.com/AdamShannag/jot/internal/io"
	"github.com/AdamShannag/jot/internal/types"
	"github.com/spf13/afero"
)

func Test_Run_jot(t *testing.T) {
	args := os.Args[0:1]
	run(args)
}

func Test_Run_jot_init(t *testing.T) {
	appFs := f.SwitchToMemMap()
	runWith("init", ".", "test")

	// test directories and jot.yaml exists
	checkFileOrDir(t, &appFs, "./test/")
	checkFileOrDir(t, &appFs, "./test/jot.yaml")

	// check file content
	expectedSpecs := types.NewSpecs("test", nil, nil)
	actualSpecs, err := f.ToSpecs("./test/jot.yaml")
	if err != nil {
		t.Error(err)
	}

	assertEquals(t, expectedSpecs.Version, actualSpecs.Version)
	assertEquals(t, expectedSpecs.Project.Name, actualSpecs.Project.Name)
	assertEquals(t, expectedSpecs.Project.Created_at, actualSpecs.Project.Created_at)
}

func Test_Run_jot_add_service(t *testing.T) {
	appFs := f.SwitchToMemMap()
	runWith("init")
	runWith("add", "-srv", "test", "-p", "8080")

	// test directories
	checkFileOrDir(t, &appFs, "./test-service/bin/")
	checkFileOrDir(t, &appFs, "./test-service/go.mod")
	checkFileOrDir(t, &appFs, "./test-service/cmd/test/")
	checkFileOrDir(t, &appFs, "./test-service/deploy/image/test-service.dockerfile")

	// check jot.yaml file
	expectedSpecs := types.NewSpecs("test", []types.Service{*types.NewService("test", 8080, nil)}, nil)
	actualSpecs, err := f.ToSpecs("./jot.yaml")

	if err != nil {
		t.Error(err)
	}

	expectedService := expectedSpecs.Services[0]
	actualService := actualSpecs.Services[0]
	assertEquals(t, expectedService.Name, actualService.Name)
	assertEquals(t, string(rune(expectedService.Port)), string(rune(actualService.Port)))

}

func Test_Run_jot_add_service_endpoint(t *testing.T) {
	appFs := f.SwitchToMemMap()
	runWith("init")
	runWith("add", "-srv", "test", "-p", "8080", "--rest", "-endpoints", "users")

	// test directories
	checkFileOrDir(t, &appFs, "./test-service/bin/")
	checkFileOrDir(t, &appFs, "./test-service/go.mod")
	checkFileOrDir(t, &appFs, "./test-service/cmd/test/")
	checkFileOrDir(t, &appFs, "./test-service/deploy/image/test-service.dockerfile")
	checkFileOrDir(t, &appFs, "./test-service/api/api.go")
	checkFileOrDir(t, &appFs, "./test-service/api/handler/users/users.go")

	// check jot.yaml file
	expectedSpecs := types.NewSpecs("test", []types.Service{*types.NewService("test", 8080, []string{"users"})}, nil)
	actualSpecs, err := f.ToSpecs("./jot.yaml")

	if err != nil {
		t.Error(err)
	}

	expectedService := expectedSpecs.Services[0]
	actualService := actualSpecs.Services[0]

	assertEquals(t, expectedService.Name, actualService.Name)
	assertEquals(t, expectedService.Endpoints[0], actualService.Endpoints[0])
	assertEquals(t, string(rune(expectedService.Port)), string(rune(actualService.Port)))
}

func runWith(arguments ...string) {
	args := os.Args[0:1]
	args = append(args, arguments...)
	run(args)
}

func assertEquals(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("Expected to find [%s], but got [%s]", expected, actual)
	}
}

func checkFileOrDir(t *testing.T, appFs *afero.Fs, s string) {
	ok, err := afero.Exists(*appFs, s)
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Errorf("Expected to find file or directory at [%s], but was not", s)
	}
}
