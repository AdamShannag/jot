package template

import (
	"embed"
	"fmt"
	"text/template"

	"github.com/AdamShannag/jot/internal/io"
	"github.com/fatih/color"
)

var (
	//go:embed templates/*
	resources embed.FS
	tmpl      = template.Must(template.ParseFS(resources, "templates/*"))
)

func Create(tplfile, path, filename string, data any) {
	io.TplToFile(
		tmpl,
		tplfile,
		path,
		filename,
		data,
	)
	fmt.Printf("%s %s\n", color.CyanString("[%s]", filename), color.GreenString("Created!"))
}
