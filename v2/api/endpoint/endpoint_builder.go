package endpoint

import "github.com/AdamShannag/jot/v2/types/model"

type modification func(*model.Endpoint)
type actions struct{ modifications []modification }
type Endpoint struct{ actions *actions }
