package template

import (
	"embed"
	"text/template"

	"github.com/AdamShannag/jot/v2/writer/util"
)

var (
	//go:embed templates/*
	resources embed.FS
	tmpl      = template.Must(template.ParseFS(resources, "templates/*"))
)

func Create(path, name, tpl string, data any) {
	util.CreateFile(tmpl, path, name, tpl, data)
}
