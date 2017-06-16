package syntax

import "github.com/gtong/gojs/types"

type WhileNode struct {
	Expression     Node
	Statements   *StatementsNode
}

func (n WhileNode) Eval(ctx *types.Context) (types.Value, error) {
	loop := true
	var retVal types.Value
	for loop {
		expr, err := n.Expression.Eval(ctx)
		if err != nil {
			return nil, err
		}
		exprBool, err := expr.ToBooleanValue(ctx)
		if err != nil {
			return nil, err
		}
		loop = exprBool.Value
		if loop {
			retVal, err = n.Statements.Eval(ctx)
			if err != nil {
				return nil, err
			}
		}
	}
	return retVal, nil
}
