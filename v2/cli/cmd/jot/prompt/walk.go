package prompt

import (
	"io/fs"
	"log"
	"path/filepath"

	"github.com/AdamShannag/jot/v2/api/service"
)

func (p *PrompterImpl) walkServices() {
	err := filepath.WalkDir(".", p.visit)
	if err != nil {
		log.Fatal(err)
	}
}

func (p *PrompterImpl) walkService(jotService string) {
	if p.pathExists(filepath.Join(jotService, "cmd")) {
		filepath.WalkDir(filepath.Join(jotService, "cmd"), func(path string, d fs.DirEntry, _ error) error {
			dir, err := d.Info()
			if err != nil {
				return err
			}

			if dir.Name() == "main.go" {
				p.setProjectPath(jotService)
				p.services = append(p.services, service.NewBuilder().Name(jotService).Build())
			}

			return nil
		})
	}
}

func (p *PrompterImpl) visit(path string, d fs.DirEntry, _ error) error {
	dir, err := d.Info()
	if err != nil {
		return err
	}

	if skip := p.skipDirs(dir); skip != nil {
		return skip
	}

	if dir.IsDir() {
		p.walkService(path)
	}

	return nil
}

func (p *PrompterImpl) setProjectPath(path string) {
	if p.projectPath == "" {
		pathList := filepath.SplitList(path)
		if len(pathList) > 1 {
			p.projectPath = pathList[0]
		} else {
			p.projectPath = "./"
		}
	}
}

func (p *PrompterImpl) skipDirs(path fs.FileInfo) error {
	if path.IsDir() {
		if !p.pathExists(filepath.Join(path.Name(), "cmd")) && path.Name() != "." {
			return filepath.SkipDir
		}
	}

	return nil
}
