package url

import "github.com/AdamShannag/jot/v2/types/model"

type modification func(*model.Url)
type actions struct{ modifications []modification }
type Url struct{ actions *actions }
