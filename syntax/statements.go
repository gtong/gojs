package syntax

import "github.com/gtong/gojs/types"

type StatementsNode struct {
	Statements []Node
}

func (n *StatementsNode) Append(statement Node) {
	n.Statements = append(n.Statements, statement)
}

func (n StatementsNode) Eval(ctx *types.Context) (types.Value, error) {
	var ret types.Value
	for _, statement := range n.Statements {
		var err error
		ret, err = statement.Eval(ctx)
		if err != nil {
			return nil, err
		}
	}
	return ret.ToActualValue(ctx)
}
