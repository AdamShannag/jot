package file

import (
	"fmt"

	"github.com/AdamShannag/jot/v2/writer/template"
)

type MainFile struct {
	name string
	tpl  template.Template
}

func NewMainFile() *MainFile {
	return &MainFile{"main", template.MAIN}
}

func (f *MainFile) Write(path string) {
	template.Create(path, fmt.Sprintf("%s.go", f.name), string(f.tpl), nil)
}
