package directory

import (
	"log"
	"os"

	"github.com/AdamShannag/jot/v2/writer/fs"
	"github.com/spf13/afero"
)

type DirectoryWriter interface {
	Write(string)
}

type DefualtDirectoryWriter struct{}

func (w DefualtDirectoryWriter) Write(path string) {
	if w.pathExists(path) {
		return
	}
	if err := fs.Get().MkdirAll(path, os.ModePerm); err != nil {
		log.Fatalf("Error creating directories at %s, err: %v", path, err)
	}
}

func (p DefualtDirectoryWriter) pathExists(path string) bool {
	ok, err := afero.Exists(fs.Get(), path)
	if err != nil {
		log.Fatalf("Error occured: %v", err)
	}
	return ok
}
