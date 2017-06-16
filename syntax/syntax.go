package syntax

import (
	"github.com/gtong/gojs/types"
)

type Node interface {
	Eval(*types.Context) (types.Value, error)
}

func testExpr(ctx *types.Context, expr Node) (bool, error) {
	val, err := expr.Eval(ctx)
	if err != nil {
		return false, err
	}
	boolVal, err := val.ToBooleanValue(ctx)
	if err != nil {
		return false, err
	}
	return boolVal.Value, nil
}
