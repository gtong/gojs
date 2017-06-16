package syntax

import (
	"github.com/gtong/gojs/types"
)

type ErrorReturn struct {
	Value types.Value
}

func (e *ErrorReturn) Error() string {
	return "return"
}

type ReturnNode struct {
	Expression Node
}

func (n ReturnNode) Eval(ctx *types.Context) (types.Value, error) {
	val, err := n.Expression.Eval(ctx)
	if err != nil {
		return nil, err
	}
	aVal, err := val.ToActualValue(ctx)
	if err != nil {
		return nil, err
	}
	return nil, &ErrorReturn{Value: aVal}
}

type FunctionNode struct {
	Statements *StatementsNode
}

func (n FunctionNode) Eval(ctx *types.Context) (types.Value, error) {
	return types.FunctionValue{Statements: n.Statements}, nil
}

type CallNode struct {
	Expression Node
}

func (n CallNode) Eval(ctx *types.Context) (types.Value, error) {
	val, err := n.Expression.Eval(ctx)
	if err != nil {
		return nil, err
	}
	if _, err := val.Call(ctx, nil); err != nil {
		if errReturn, ok := err.(*ErrorReturn); ok {
			return errReturn.Value, nil
		}
		return nil, err
	}
	return nil, nil
}
