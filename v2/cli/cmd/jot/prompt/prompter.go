package prompt

import "github.com/AdamShannag/jot/v2/types/model"

type Validator func(string) error

type Prompter interface {
	Start(string)
	Prompt(string, Validator) string
	Select(string, ...string) (int, string)
}

const (
	NEW string = "new"

	SERVICE    string = "Service"
	ENDPOINT   string = "Endpoint"
	MIDDLEWARE string = "Middleware"
)

type PrompterImpl struct {
	services    []model.Service
	projectPath string
}

func NewPrompterImpl() *PrompterImpl {
	return &PrompterImpl{
		[]model.Service{},
		"",
	}
}
