package directory

import (
	"log"
	"os"

	"github.com/AdamShannag/jot/v2/writer/fs"
)

type DirectoryWriter interface {
	Write(string)
}

type DefualtDirectoryWriter struct{}

func (w DefualtDirectoryWriter) Write(path string) {
	if err := fs.Get().MkdirAll(path, os.ModePerm); err != nil {
		log.Fatalf("Error creating directories at %s, err: %v", path, err)
	}
}
