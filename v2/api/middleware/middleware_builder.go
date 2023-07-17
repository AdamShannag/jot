package middleware

import "github.com/AdamShannag/jot/v2/types/model"

type modification func(*model.Middleware)
type actions struct{ modifications []modification }
type Middleware struct{ actions *actions }
