package syntax

import "github.com/gtong/gojs/types"

type IfNode struct {
	Expression Node
	IfStatements *StatementsNode
	ElseStatements *StatementsNode
}

func (n IfNode) Eval(ctx *types.Context) (types.Value, error) {
	expr, err := n.Expression.Eval(ctx)
	if err != nil {
		return nil, err
	}
	exprBool, err := expr.ToBooleanValue(ctx)
	if err != nil {
		return nil, err
	}
	if exprBool.Value == true {
		return n.IfStatements.Eval(ctx)
	} else if n.ElseStatements != nil {
		return n.ElseStatements.Eval(ctx)
	}
	return nil, nil
}
