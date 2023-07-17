package template

import (
	"embed"
	"fmt"
	"text/template"

	"github.com/AdamShannag/jot/v1/internal/command/log"
	"github.com/AdamShannag/jot/v1/internal/io"
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
	log.Info(fmt.Sprintf("%s%s", path, filename), log.CREATED)
}
