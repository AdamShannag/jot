package file

import (
	"fmt"

	"github.com/AdamShannag/jot/v2/writer/template"
)

type LoggerFile struct {
	name string
	tpl  template.Template
}

func NewLoggerFile() *LoggerFile {
	return &LoggerFile{"logger", template.Logger}
}

func (f *LoggerFile) Write(path string) {
	template.Create(path, fmt.Sprintf("%s.go", f.name), string(f.tpl), nil)
}
