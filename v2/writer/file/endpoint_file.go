package file

import (
	"fmt"
	"strings"

	"github.com/AdamShannag/jot/v2/writer/template"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type EndpointFile struct {
	name string
	data map[string]string
	tpl  template.Template
}

func NewEndpointFile(name string, data map[string]string) *EndpointFile {
	if data == nil {
		data = make(map[string]string)
	}
	data["PackageName"] = name
	data["Name"] = cases.Title(language.English, cases.NoLower).String(strings.ToLower(name))

	return &EndpointFile{name, data, template.Endpoint}
}

func (f *EndpointFile) Write(path string) {
	template.Create(path, fmt.Sprintf("%s.go", f.name), string(f.tpl), f.data)
}
