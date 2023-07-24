package template

import (
	"embed"
	"log"
	"text/template"

	"github.com/AdamShannag/jot/v2/writer/fs"
)

var (
	//go:embed templates/*
	resources embed.FS
	tmpl      = template.Must(template.ParseFS(resources, "templates/*"))
)

func Create(path, name, tpl string, data any) {
	file, err := fs.Get().Create(path + name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	err = tmpl.ExecuteTemplate(file, tpl, data)
	if err != nil {
		log.Fatal(err)
	}
}
