package writer

import (
	"github.com/AdamShannag/jot/v2/writer/template"
	"github.com/AdamShannag/jot/v2/writer/util"
)

type Dir struct {
	Name  string
	Files []File
	Dirs  []*Dir
}

func (d *Dir) Make(path string) {
	d.make(path)
}

func (d *Dir) make(p string) {
	newPath := p + d.Name + "/"

	util.CreateDirs(newPath)
	for _, file := range d.Files {
		template.Create(newPath, file.GetFileName(), file.GetTpl(), file.Data)
	}

	for _, dr := range d.Dirs {
		dr.make(newPath)
	}
}
