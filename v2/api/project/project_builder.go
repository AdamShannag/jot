package project

import "github.com/AdamShannag/jot/v2/types/model"

type modification func(*model.Project)
type actions struct{ modifications []modification }
type Project struct{ actions *actions }
