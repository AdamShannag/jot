package directory

import (
	"fmt"

	"github.com/AdamShannag/jot/v2/writer/file"
)

type Directory struct {
	Name        string
	Directories []*Directory
	Files       []file.File
	writer      DirectoryWriter
}

// Instantiates a new Directory with the specified name, directories and files, using the DefualtDirectoryWriter
func NewDefaultDirectory(name string, directories []*Directory, files ...file.File) *Directory {
	return &Directory{name, directories, files, DefualtDirectoryWriter{}}
}

// Instantiates a new Directory with the specified name,directory writer, directories and files
func NewDirectory(name string, directories []*Directory, files []file.File, writer DirectoryWriter) *Directory {
	return &Directory{name, directories, files, writer}
}

// Creates the directories with thier related files and directories at the specified path
func (d *Directory) Create(p string) {
	path := fmt.Sprintf("%s%s/", p, d.Name)

	d.writer.Write(path)

	for _, file := range d.Files {
		file.Write(path)
	}

	for _, dr := range d.Directories {
		dr.Create(path)
	}
}

// Gets a directory by name, returns (nil, false) if not found
func (d *Directory) Get(name string) (*Directory, bool) {
	if d.Name == name {
		return d, true
	}

	for _, dir := range d.Directories {
		dr, ok := dir.Get(name)
		if ok {
			return dr, true
		}
	}

	return nil, false
}

// Inserts a new directory at the given directory name, panics if directory name does not exist
func (d *Directory) InsertAt(name string, dir *Directory) {
	if d, ok := d.Get(name); ok {
		d.Directories = append(d.Directories, dir)
	} else {
		panic("directory not found")
	}
}
