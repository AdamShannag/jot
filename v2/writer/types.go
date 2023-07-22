package writer

import "github.com/AdamShannag/jot/v2/types/model"

type writer struct {
	project *model.Project
}

type dirFilesMap map[string][]File
