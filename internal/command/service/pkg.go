package service

import (
	p "github.com/AdamShannag/jot/internal/command/path"
	s "github.com/AdamShannag/jot/internal/command/suffix"
	"github.com/AdamShannag/jot/internal/template"
)

func createZerologPkg(service string) {
	template.Create(
		p.ZerologPkgTpl,
		p.Path(p.ZerologPkgPath, s.ServiceSuffix(service)),
		s.GoSuffix(p.ZerologPkgname),
		nil,
	)
}
