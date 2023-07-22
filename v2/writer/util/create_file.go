package util

import (
	"log"
	"text/template"
)

func CreateFile(tmpl *template.Template, path, name, tpl string, data any) {
	file, err := fs.Create(path + name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	err = tmpl.ExecuteTemplate(file, tpl, data)
	if err != nil {
		log.Fatal(err)
	}
}
