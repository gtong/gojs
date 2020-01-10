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
	Line       int
	Column     int
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
	Line       int
	Column     int
	Parameters []string
	Statements *StatementsNode
}

func (n FunctionNode) Eval(ctx *types.Context) (types.Value, error) {
	return types.NewFunctionValue(ctx, n.Parameters, n.Statements), nil
}

type CallNode struct {
	Line       int
	Column     int
	Expression Node
	Arguments  []Node
}

func (n CallNode) Eval(ctx *types.Context) (types.Value, error) {
	val, err := n.Expression.Eval(ctx)
	if err != nil {
		return nil, err
	}
	var args []types.Value
	for _, arg := range n.Arguments {
		arg, err := arg.Eval(ctx)
		if err != nil {
			return nil, err
		}
		args = append(args, arg)
	}
	if _, err := val.Call(ctx, args); err != nil {
		if errReturn, ok := err.(*ErrorReturn); ok {
			return errReturn.Value, nil
		}
		return nil, err
	}
	return nil, nil
}
