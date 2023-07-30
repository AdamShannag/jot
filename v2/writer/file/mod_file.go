package file

import (
	"fmt"

	"github.com/AdamShannag/jot/v2/writer/template"
)

type ModFile struct {
	name string
	tpl  template.Template
	data map[string]string
}

func NewModFile(serviceName string) *ModFile {
	return &ModFile{"go", template.MOD, map[string]string{"ServiceName": serviceName}}
}

func (f *ModFile) Write(path string) {
	template.Create(path, fmt.Sprintf("%s.mod", f.name), string(f.tpl), f.data)
}
