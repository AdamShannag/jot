package service

import "github.com/AdamShannag/jot/v2/types/model"

type modification func(*model.Service)
type actions struct{ modifications []modification }
type Service struct{ actions *actions }
