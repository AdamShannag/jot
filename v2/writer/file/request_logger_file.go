package file

import (
	"fmt"

	"github.com/AdamShannag/jot/v2/writer/template"
)

type RequestLoggerFile struct {
	name string
	tpl  template.Template
}

func NewRequestLoggerFile() *RequestLoggerFile {
	return &RequestLoggerFile{"request_logger", template.LoggerMiddleware}
}

func (f *RequestLoggerFile) Write(path string) {
	template.Create(path, fmt.Sprintf("%s.go", f.name), string(f.tpl), nil)
}
