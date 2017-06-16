package syntax

import "github.com/gtong/gojs/types"

type ConditionalNode struct {
	Expression Node
	Statements *StatementsNode
}

func (n ConditionalNode) Eval(ctx *types.Context) (types.Value, error) {
	expr, err := n.Expression.Eval(ctx)
	if err != nil {
		return nil, err
	}
	exprBool, err := expr.ToBooleanValue(ctx)
	if err != nil {
		return nil, err
	}
	if exprBool.Value == true {
		return n.Statements.Eval(ctx)
	}
	return nil, nil
}
