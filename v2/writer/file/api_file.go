package file

import (
	"fmt"

	"github.com/AdamShannag/jot/v2/writer/template"
)

type ApiFile struct {
	name string
	tpl  template.Template
}

func NewApiFile() *ApiFile {
	return &ApiFile{"api", template.API}
}

func (f *ApiFile) Write(path string) {
	template.Create(path, fmt.Sprintf("%s.go", f.name), string(f.tpl), nil)
}
