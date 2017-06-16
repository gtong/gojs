package syntax

import (
	"fmt"

	"github.com/gtong/gojs/types"
)

const (
	ADD_OP                   = "+"
	SUBTRACT_OP              = "-"
	ASSIGNMENT_OP            = "="
	GREATER_THAN_OP          = ">"
	LESS_THAN_OP             = "<"
	GREATER_THAN_OR_EQUAL_OP = ">="
	LESS_THAN_OR_EQUAL_OP    = "<="
	EQUALITY_OP              = "=="
	INEQUALITY_OP            = "!="
	EQUALITY_OP_STRICT       = "==="
	INEQUALITY_OP_STRICT     = "!=="
)

type BinaryOpNode struct {
	Left     Node
	Right    Node
	Operator string
}

func (n BinaryOpNode) Eval(ctx *types.Context) (types.Value, error) {
	lv, err := n.Left.Eval(ctx)
	if err != nil {
		return nil, err
	}
	rv, err := n.Right.Eval(ctx)
	if err != nil {
		return nil, err
	}
	switch n.Operator {
	case ADD_OP:
		return lv.Add(ctx, rv)
	case SUBTRACT_OP:
		return lv.Subtract(ctx, rv)
	case ASSIGNMENT_OP:
		return lv.Assign(ctx, rv)
	case EQUALITY_OP:
		return comp(ctx, lv, rv, false, false, 0)
	case INEQUALITY_OP:
		return comp(ctx, lv, rv, false, true, 0)
	case EQUALITY_OP_STRICT:
		return comp(ctx, lv, rv, true, false, 0)
	case INEQUALITY_OP_STRICT:
		return comp(ctx, lv, rv, true, true, 0)
	case GREATER_THAN_OP:
		return comp(ctx, lv, rv, false, false, 1)
	case GREATER_THAN_OR_EQUAL_OP:
		return comp(ctx, lv, rv, false, false, 0, 1)
	case LESS_THAN_OP:
		return comp(ctx, lv, rv, false, false, -1)
	case LESS_THAN_OR_EQUAL_OP:
		return comp(ctx, lv, rv, false, false, -1, 0)
	default:
		return nil, fmt.Errorf("operator %s not recognized", n.Operator)
	}
}

func comp(ctx *types.Context, lv types.Value, rv types.Value, strict bool, invert bool, expect ...int) (types.Value, error) {
	cmp, forceFalse, err := lv.Compare(ctx, rv, strict)
	if err != nil {
		return nil, err
	}
	if forceFalse {
		if invert {
			return types.BooleanValue{Value: true}, nil
		} else {
			return types.BooleanValue{Value: false}, nil
		}
	}
	successValue := !invert
	for _, val := range expect {
		if val == cmp {
			return types.BooleanValue{Value: successValue}, nil
		}
	}
	return types.BooleanValue{Value: !successValue}, nil
}
