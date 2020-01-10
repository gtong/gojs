package syntax

import (
	"errors"

	"github.com/gtong/gojs/types"
)

var ErrBreak = errors.New("break")
var ErrContinue = errors.New("continue")

type BreakNode struct {
	Line   int
	Column int
}

func (n BreakNode) Eval(ctx *types.Context) (types.Value, error) {
	return nil, ErrBreak
}

type ContinueNode struct {
	Line   int
	Column int
}

func (n ContinueNode) Eval(ctx *types.Context) (types.Value, error) {
	return nil, ErrContinue
}

type WhileNode struct {
	Expression Node
	Statements *StatementsNode
}

func (n WhileNode) Eval(ctx *types.Context) (types.Value, error) {
	var retVal types.Value
	var err error

	loop := true
	for loop {
		loop, err = testExpr(ctx, n.Expression)
		if err != nil {
			return nil, err
		}
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

type ForNode struct {
	InitExpr   Node
	TestExpr   Node
	IncrExpr   Node
	Statements *StatementsNode
}

func (n ForNode) Eval(ctx *types.Context) (types.Value, error) {
	// Initialize
	if _, err := n.InitExpr.Eval(ctx); err != nil {
		return nil, err
	}

	var retVal types.Value
	var err error

	loop := true
	for loop {
		loop, err = testExpr(ctx, n.TestExpr)
		if err != nil {
			return nil, err
		}
		if loop {
			retVal, err = n.Statements.Eval(ctx)
			if err == ErrBreak {
				return nil, nil
			} else if err == ErrContinue {
				// Continue
			} else if err != nil {
				return nil, err
			}
			if _, err := n.IncrExpr.Eval(ctx); err != nil {
				return nil, err
			}
		}
	}
	return retVal, nil
}
