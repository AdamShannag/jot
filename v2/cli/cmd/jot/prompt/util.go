package prompt

import (
	"log"

	f "github.com/AdamShannag/jot/v2/writer/fs"
	"github.com/spf13/afero"
)

func (p *PrompterImpl) servicesSlice() (srv []string) {
	for _, s := range p.services {
		srv = append(srv, s.Name)
	}
	srv = append(srv, "../")
	return
}

func (p *PrompterImpl) pathExists(path string) bool {
	ok, err := afero.Exists(f.Get(), path)
	if err != nil {
		log.Fatalf("Error occured: %v", err)
	}
	return ok
}
