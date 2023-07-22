package writer

import (
	"github.com/AdamShannag/jot/v2/types/model"
)

func NewWriter(p *model.Project) *writer {
	return &writer{p}
}

func (w *writer) Write(path string) {
	projectDir := Dir{
		Name:  w.project.Name,
		Files: nil,
		Dirs:  w.servicesDirectories()}

	projectDir.Make(path)
}
