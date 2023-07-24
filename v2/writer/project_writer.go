package writer

import (
	"github.com/AdamShannag/jot/v2/types/model"
	d "github.com/AdamShannag/jot/v2/writer/directory"
)

type projectWriter struct {
	project *model.Project
}

func NewProjectWriter(p *model.Project) *projectWriter {
	return &projectWriter{p}
}

func (w *projectWriter) Write(path string) {
	d.NewDefaultDirectory(w.project.Name, w.constructServices()).Create(path)
}
