package syntax

import (
	"github.com/gtong/gojs/types"
)

type Node interface {
	Eval(*types.Context) (types.Value, error)
}
