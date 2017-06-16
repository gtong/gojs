package types

import (
	"errors"
)

var _ Value = FunctionValue{}

type Evaluable interface {
	Eval(*Context) (Value, error)
}

type FunctionValue struct {
	Statements Evaluable
}

func (a FunctionValue) ToString(ctx *Context) (string, error) {
	return "[Function]", nil
}

func (a FunctionValue) ToActualValue(ctx *Context) (Value, error) {
	return a, nil
}

func (a FunctionValue) ToStringValue(ctx *Context) (StringValue, error) {
	return StringValue{Value: "function() {}"}, nil
}

func (a FunctionValue) ToNumberValue(ctx *Context) (NumberValue, error) {
	return NaN, nil
}

func (a FunctionValue) ToBooleanValue(ctx *Context) (BooleanValue, error) {
	return BooleanValue{Value: true}, nil
}

func (a FunctionValue) Add(ctx *Context, b Value) (Value, error) {
	sa, err := a.ToStringValue(ctx)
	if err != nil {
		return nil, err
	}
	return sa.Add(ctx, b)
}

func (a FunctionValue) Subtract(ctx *Context, b Value) (Value, error) {
	na, err := a.ToNumberValue(ctx)
	if err != nil {
		return nil, err
	}
	return na.Subtract(ctx, b)
}

func (a FunctionValue) Assign(ctx *Context, b Value) (Value, error) {
	return nil, errors.New("ReferenceError: Invalid left-hand side in assignment")
}

func (a FunctionValue) Increment(ctx *Context, value int) (Value, error) {
	na, err := a.ToNumberValue(ctx)
	if err != nil {
		return nil, err
	}
	return na.Increment(ctx, value)
}

func (a FunctionValue) Compare(ctx *Context, b Value, strict bool) (int, bool, error) {
	aa, err := a.ToActualValue(ctx)
	if err != nil {
		return 0, false, err
	}
	ab, err := a.ToActualValue(ctx)
	if err != nil {
		return 0, false, err
	}
	if aa == ab {
		return 0, false, nil
	} else {
		return 0, true, nil
	}
}

func (a FunctionValue) Call(ctx *Context, args []Value) (Value, error) {
	newContext := &Context{}
	return a.Statements.Eval(newContext)
}
