package io

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/AdamShannag/jot/internal/types"
	"github.com/spf13/afero"
	yml "gopkg.in/yaml.v3"
)

var appFs = afero.NewOsFs()

func ToSpecs(path string) (*types.Specs, error) {
	specs := types.Specs{}
	b, err := afero.ReadFile(appFs, path)
	if err != nil {
		return nil, err
	}
	err = yml.Unmarshal(b, &specs)
	if err != nil {
		return nil, err
	}
	return &specs, nil
}

func ToEmptyFile(path, filename string, data []byte) {
	fullpath := fmt.Sprintf("%s%s", path, filename)
	if err := appFs.MkdirAll(path, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	isEmpty, err := afero.IsEmpty(appFs, path)
	if err != nil {
		log.Fatal(err)
	}
	if !isEmpty {
		log.Fatal("the specified directory is not empty")
	}
	if err := afero.WriteFile(appFs, fullpath, data, 0777); err != nil {
		log.Fatal(err)
	}
}

func ToFile(path, filename string, data []byte) {
	fullpath := fmt.Sprintf("%s%s", path, filename)
	if err := appFs.MkdirAll(path, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	if err := afero.WriteFile(appFs, fullpath, data, 0777); err != nil {
		log.Fatal(err)
	}
}

func ToDirs(path string) {
	if err := appFs.MkdirAll(path, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

func FileExists(path string) (bool, error) {
	if ok, err := afero.Exists(appFs, path); err != nil {
		return false, err
	} else {
		return ok, nil
	}
}

func TplToFile(tmpl *template.Template, tpl string, path string, name string, data any) {
	if err := appFs.MkdirAll(path, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	file, err := appFs.Create(path + name)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.ExecuteTemplate(file, tpl, data)
	if err != nil {
		log.Fatal(err)
	}
}
