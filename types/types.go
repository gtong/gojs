package types

import (
	"errors"
)

type Context struct {
	Parent    *Context
	Variables map[string]Value
}

func (c *Context) Set(s string, v Value) Value {
	if c.Variables == nil {
		c.Variables = make(map[string]Value)
	}
	av, err := v.ToActualValue(c)
	if err != nil {
		return nil
	}
	c.Variables[s] = av
	return v
}

func (c Context) Get(s string) (Value, error) {
	v, ok := c.Variables[s]
	if !ok {
		if c.Parent != nil {
			return c.Parent.Get(s)
		}
		return nil, errors.New("ReferenceError: " + s + " is not defined")
	}
	return v, nil
}

func (c *Context) FindContext(s string) *Context {
	if _, ok := c.Variables[s]; ok {
		return c
	}
	if c.Parent != nil {
		return c.Parent.FindContext(s)
	}
	return nil
}

type Value interface {
	ToString(*Context) (string, error)

	ToActualValue(*Context) (Value, error)
	ToStringValue(*Context) (StringValue, error)
	ToNumberValue(*Context) (NumberValue, error)
	ToBooleanValue(*Context) (BooleanValue, error)

	// Rules for addition:
	// If either operand is a string, do string concatenation
	// If both operands are numbers, do addition
	Add(*Context, Value) (Value, error)

	// Rules for subtraction:
	// Convert both operands to numbers, do subtraction
	Subtract(*Context, Value) (Value, error)

	// Rules for assignment:
	Assign(*Context, Value) (Value, error)
	Increment(*Context, int) (Value, error)

	// Rules for Compare:
	// Returns 0 if they are equal
	// Returns >0 if this value is > passed-in value
	// Returns <0 if this value is < passed-in value
	Compare(*Context, Value, bool) (int, bool, error)

	Call(*Context, []Value) (Value, error)
}
