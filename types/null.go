package types

import (
	"errors"
)

var _ Value = NullValue{}

// Constants definition
var Null NullValue = NullValue{}

type NullValue struct {
}

func (a NullValue) ToString(ctx *Context) (string, error) {
	return "null", nil
}

func (a NullValue) ToActualValue(ctx *Context) (Value, error) {
	return a, nil
}

func (a NullValue) ToStringValue(ctx *Context) (StringValue, error) {
	return StringValue{Value: "null"}, nil
}

func (a NullValue) ToNumberValue(ctx *Context) (NumberValue, error) {
	return NumberValue{Value: 0}, nil
}

func (a NullValue) ToBooleanValue(ctx *Context) (BooleanValue, error) {
	return BooleanValue{Value: false}, nil
}

func (a NullValue) Add(ctx *Context, b Value) (Value, error) {
	ab, err := b.ToActualValue(ctx)
	if err != nil {
		return nil, err
	}
	if _, ok := ab.(StringValue); ok {
		sa, err := a.ToStringValue(ctx)
		if err != nil {
			return nil, err
		}
		return sa.Add(ctx, b)
	}
	na, err := a.ToNumberValue(ctx)
	if err != nil {
		return nil, err
	}
	return na.Add(ctx, b)
}

func (a NullValue) Subtract(ctx *Context, b Value) (Value, error) {
	na, err := a.ToNumberValue(ctx)
	if err != nil {
		return nil, err
	}
	return na.Subtract(ctx, b)
}

func (a NullValue) Assign(ctx *Context, b Value) (Value, error) {
	return nil, errors.New("ReferenceError: Invalid left-hand side in assignment")
}

func (a NullValue) Increment(ctx *Context, value int) (Value, error) {
	return nil, errors.New("ReferenceError: Invalid left-hand side expression in postfix operation")
}

func (a NullValue) Compare(ctx *Context, b Value, strict bool) (int, bool, error) {
	ab, err := b.ToActualValue(ctx)
	if err != nil {
		return 0, false, err
	}
	if _, ok := ab.(NullValue); ok {
		return 0, false, nil
	}
	return 0, true, nil
}

func (a NullValue) Call(ctx *Context, args []Value) (Value, error) {
	return nil, errors.New("TypeError: null is not a function")
}
