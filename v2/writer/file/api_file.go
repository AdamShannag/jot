package file

import (
	"fmt"

	"github.com/AdamShannag/jot/v2/writer/template"
)

type ApiFile struct {
	name string
	data map[string]string
	tpl  template.Template
}

func NewApiFile(name string, data map[string]string) *ApiFile {
	return &ApiFile{name, data, template.API}
}

func (f *ApiFile) Write(path string) {
	template.Create(path, fmt.Sprintf("%s.go", f.name), string(f.tpl), f.data)
}
