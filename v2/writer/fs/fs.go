package fs

import (
	"sync"

	"github.com/spf13/afero"
)

var once sync.Once

var fs afero.Fs

func Get() afero.Fs {
	once.Do(func() {
		fs = afero.NewOsFs()
	})

	return fs
}
