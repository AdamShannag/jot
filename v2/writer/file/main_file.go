package file

import (
	"fmt"

	"github.com/AdamShannag/jot/v2/writer/template"
)

type MainFile struct {
	name string
	data map[string]string
	tpl  template.Template
}

func NewMainFile(name string, data map[string]string) *MainFile {
	return &MainFile{name, data, template.MAIN}
}

func (f *MainFile) Write(path string) {
	template.Create(path, fmt.Sprintf("%s.go", f.name), string(f.tpl), f.data)
}
