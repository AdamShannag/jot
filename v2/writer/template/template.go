package template

import (
	"embed"
	"fmt"
	"log"
	"path/filepath"
	"text/template"

	"github.com/AdamShannag/jot/v2/writer/fs"
	"github.com/fatih/color"
	"github.com/spf13/afero"
)

var (
	//go:embed templates/*
	resources embed.FS
	tmpl      = template.Must(template.ParseFS(resources, "templates/*"))
)

func Create(path, name, tpl string, data any) {
	filePath := filepath.Join(path, name)
	if pathExists(filePath) {
		return
	}
	file, err := fs.Get().Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	err = tmpl.ExecuteTemplate(file, tpl, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(color.GreenString("CREATE"), filePath)
}

func pathExists(path string) bool {
	ok, err := afero.Exists(fs.Get(), path)
	if err != nil {
		log.Fatalf("Error occured: %v", err)
	}
	return ok
}
