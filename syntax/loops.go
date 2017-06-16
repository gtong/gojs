package syntax

import (
	"errors"

	"github.com/gtong/gojs/types"
)

var ErrBreak = errors.New("break")
var ErrContinue = errors.New("continue")

type WhileNode struct {
	Expression Node
	Statements *StatementsNode
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
			if err == ErrBreak {
				return nil, nil
			} else if err == ErrContinue {
				// Continue
			} else if err != nil {
				return nil, err
			}
		}
	}
	return retVal, nil
}

type BreakNode struct{}

func (n BreakNode) Eval(ctx *types.Context) (types.Value, error) {
	return nil, ErrBreak
}

type ContinueNode struct{}

func (n ContinueNode) Eval(ctx *types.Context) (types.Value, error) {
	return nil, ErrContinue
}
