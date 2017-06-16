package syntax

import (
	"strconv"

	"github.com/gtong/gojs/types"
)

type IdentifierNode struct {
	Value string
}

func (i IdentifierNode) Eval(ctx *types.Context) (types.Value, error) {
	return types.IdentifierValue{Value: i.Value}, nil
}

type BooleanNode struct {
	Value string
}

func (t BooleanNode) Eval(ctx *types.Context) (types.Value, error) {
	if t.Value == "true" {
		return types.BooleanValue{Value: true}, nil
	}
	return types.BooleanValue{Value: false}, nil
}

type NumberNode struct {
	Value string
}

func (t NumberNode) Eval(ctx *types.Context) (types.Value, error) {
	if t.Value == "NaN" {
		return types.NaN, nil
	}
	i, err := strconv.ParseInt(t.Value, 10, 64)
	if err != nil {
		return nil, err
	}
	return types.NumberValue{Value: i}, nil
}

type StringNode struct {
	Value string
}

func (t StringNode) Eval(ctx *types.Context) (types.Value, error) {
	return types.StringValue{Value: t.Value[1 : len(t.Value)-1]}, nil
}
