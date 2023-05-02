package io

import (
	"fmt"
	"log"
	"strings"
	"testing"
	"text/template"

	"github.com/spf13/afero"
)

const test_jot_file_path = "./test_data/jot_test.yaml"

func Test_ToSpecs(t *testing.T) {
	specs, err := ToSpecs(test_jot_file_path)
	if err != nil {
		log.Panic(err)
	}

	if specs.Version != "v1.0.0" {
		t.Errorf("Expected version to be v1.0.0, but got %s", specs.Version)
	}
	if specs.Project.Name != "test" {
		t.Errorf("Expected project.name to be test, but got %s", specs.Project.Name)
	}
	if specs.Project.Created_at != "2 May 2023" {
		t.Errorf("Expected project.created_at to be 2 May 2023, but got %s", specs.Project.Created_at)
	}
	if len(specs.Services) != 0 {
		t.Errorf("Expected number of service to be 0, but got %d", len(specs.Services))
	}
}

func Test_ToEmptyFile(t *testing.T) {
	appFs = afero.NewMemMapFs()

	data := "some data"
	path := "./test/"
	file := "test.file"
	fullpath := fmt.Sprintf("%s%s", path, file)

	ToEmptyFile(path, file, []byte(data))

	// check file content
	bs, err := afero.ReadFile(appFs, fullpath)
	if err != nil {
		log.Panic(err)
	}
	if !strings.Contains(data, string(bs)) {
		t.Errorf("Expected file to contain [%s], but got [%s]", data, string(bs))
	}

	// test if file is created
	if ok, err := afero.Exists(appFs, fullpath); err != nil {
		log.Panic(err)
	} else if !ok {
		t.Errorf("Expected file to exist at %s, but it is not", fullpath)
	}
}

func Test_ToFile(t *testing.T) {
	appFs = afero.NewMemMapFs()

	data := "some data"
	path := "./test/"
	file := "test.file"
	fullpath := fmt.Sprintf("%s%s", path, file)

	ToFile(path, file, []byte(data))

	// check file content
	bs, err := afero.ReadFile(appFs, fullpath)
	if err != nil {
		log.Panic(err)
	}
	if !strings.Contains(data, string(bs)) {
		t.Errorf("Expected file to contain [%s], but got [%s]", data, string(bs))
	}

	// test if file is created
	if ok, err := afero.Exists(appFs, fullpath); err != nil {
		log.Panic(err)
	} else if !ok {
		t.Errorf("Expected file to exist at %s, but it is not", fullpath)
	}
}

func Test_ToDirs(t *testing.T) {
	appFs = afero.NewMemMapFs()

	path := "./test/dir1/dir2/"

	ToDirs(path)

	// test if directories are created
	if ok, err := afero.DirExists(appFs, path); err != nil {
		log.Panic(err)
	} else if !ok {
		t.Errorf("Expected directory at %s, but it is not", path)
	}
}

func Test_FileExists(t *testing.T) {
	appFs = afero.NewMemMapFs()

	data := "some data"
	path := "./test/"
	file := "test.file"
	fullpath := fmt.Sprintf("%s%s", path, file)

	ToFile(path, file, []byte(data))

	if b, err := FileExists(fullpath); err != nil {
		log.Panic(err)
	} else if !b {
		t.Errorf("Expected file to exist at %s, but it is not", fullpath)
	}
}

func Test_TplToFile(t *testing.T) {
	tpl := "test.test.gotpl"
	tplPath := "./test_data/test.test.gotpl"
	filename := "test.test"
	pathTofile := "./test.test"

	tmpl := template.Must(template.ParseFiles(tplPath))
	data := map[string]string{
		"Data": "some data",
	}

	TplToFile(tmpl, tpl, "./", filename, data)

	// test if file is created
	if ok, err := afero.Exists(appFs, pathTofile); err != nil {
		log.Panic(err)
	} else if !ok {
		t.Errorf("Expected file to exist at %s, but it is not", pathTofile)
	}

	// check file content
	bs, err := afero.ReadFile(appFs, pathTofile)
	if err != nil {
		log.Panic(err)
	}
	if !strings.Contains(data["Data"], string(bs)) {
		t.Errorf("Expected file to contain [%s], but got [%s]", data["Data"], string(bs))
	}
}
