package util

import (
	"log"
	"os"

	"github.com/spf13/afero"
)

var fs afero.Fs = afero.NewOsFs()

func CreateDirs(path string) {
	if err := fs.MkdirAll(path, os.ModePerm); err != nil {
		log.Fatalf("Error creating directories at %s, err: %v", path, err)
	}
}

func IsExistingDirOrFile(path string) bool {
	ok, err := afero.Exists(fs, path)
	if err != nil {
		log.Fatalf("Error occured: %v", err)
	}
	return ok
}
