package types

import (
	"errors"
)

var _ Value = FunctionValue{}

type Evaluable interface {
	Eval(*Context) (Value, error)
}

type FunctionValue struct {
	context    *Context
	parameters []string
	statements Evaluable
}

func NewFunctionValue(ctx *Context, parameters []string, statements Evaluable) *FunctionValue {
	return &FunctionValue{context: ctx, parameters: parameters, statements: statements}
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
	values := make([]Value, len(a.parameters))
	for i, arg := range args {
		if i >= len(a.parameters) {
			break
		}
		value, err := arg.ToActualValue(ctx)
		if err != nil {
			return nil, err
		}
		values[i] = value
	}
	newContext := &Context{Parent: a.context}
	for i, param := range a.parameters {
		if values[i] == nil {
			values[i] = Null
		}
		newContext.Set(param, values[i])
	}
	return a.statements.Eval(newContext)
}
