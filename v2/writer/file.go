package writer

type File struct {
	Name, Tpl, Ext string
	Data           any
}

func (f *File) GetTpl() string {
	if f.Tpl != "" {
		return f.Tpl
	}
	return f.Name + ".gotpl"
}

func (f *File) GetFileName() string {
	if f.Ext != "" {
		return f.Name + f.Ext
	}
	return f.Name
}
