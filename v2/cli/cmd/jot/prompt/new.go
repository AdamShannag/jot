package prompt

import (
	"github.com/AdamShannag/jot/v2/api/project"
	"github.com/AdamShannag/jot/v2/writer"
)

func (p *PrompterImpl) new() {
	p.walkServices()

	switch p.creationPrompt() {
	case SERVICE:
		p.servicePrompt()
	case ENDPOINT:
		p.endpointPrompt()
	case MIDDLEWARE:
		p.middlewarePrompt()
	}

	p.projectPrompt()

	writer.NewProjectWriter(project.NewBuilder().Name("").Services(p.services).Build()).Write(p.projectPath)
}
