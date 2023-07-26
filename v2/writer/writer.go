package writer

import (
	"github.com/AdamShannag/jot/v2/types/model"
	"github.com/AdamShannag/jot/v2/writer/constructor"
)

type Writer struct {
	directoryConstructor constructor.DirectoryConstructor
}

func NewProjectWriter(p model.Project) *Writer {
	return &Writer{constructor.NewProjectDirectoryConstructor(p)}
}

func (w *Writer) Write(path string) {
	w.directoryConstructor.Construct().Create(path)
}
