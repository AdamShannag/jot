package constructor

import (
	"github.com/AdamShannag/jot/v2/writer/directory"
)

type DirectoryConstructor interface {
	Construct() *directory.Directory
}
