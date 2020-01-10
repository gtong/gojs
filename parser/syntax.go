package parser

import (
	"fmt"

	"github.com/gtong/gojs/syntax"
)

func createNullNode(o yySymType) yySymType {
	return yySymType{l: o.l, c: o.c, node: syntax.NullNode{}}
}

func createNumberNode(o yySymType) yySymType {
	node := syntax.NumberNode{
		Line:   o.l,
		Column: o.c,
		Value:  o.s,
	}
	return yySymType{l: o.l, c: o.c, node: node}
}

func createBooleanNode(o yySymType) yySymType {
	node := syntax.BooleanNode{
		Line:   o.l,
		Column: o.c,
		Value:  o.s,
	}
	return yySymType{l: o.l, c: o.c, node: node}
}

func createStringNode(o yySymType) yySymType {
	node := syntax.StringNode{
		Line:   o.l,
		Column: o.c,
		Value:  o.s,
	}
	return yySymType{l: o.l, c: o.c, node: node}
}

func createIdentifierNode(o yySymType) yySymType {
	node := syntax.IdentifierNode{
		Line:   o.l,
		Column: o.c,
		Value:  o.s,
	}
	return yySymType{l: o.l, c: o.c, node: node}
}

func createUnaryOpNode(operator yySymType, left yySymType) yySymType {
	node := syntax.UnaryOpNode{
		Line:     left.l,
		Column:   left.c,
		Left:     left.node,
		Operator: operator.s,
	}
	return yySymType{l: left.l, c: left.c, node: node}
}

func createBinaryOpNode(operator yySymType, left yySymType, right yySymType) yySymType {
	node := syntax.BinaryOpNode{
		Line:     left.l,
		Column:   left.c,
		Left:     left.node,
		Right:    right.node,
		Operator: operator.s,
	}
	return yySymType{l: left.l, c: left.c, node: node}
}

func createIfNode(expr yySymType, ifStatements yySymType) yySymType {
	ifStatementsNode := ifStatements.node.(*syntax.StatementsNode)
	node := syntax.IfNode{
		Expression:   expr.node,
		IfStatements: ifStatementsNode,
	}
	return yySymType{l: expr.l, c: expr.c, node: node}
}

func createIfElseNode(expr yySymType, ifStatements yySymType, elseStatements yySymType) yySymType {
	ifStatementsNode := ifStatements.node.(*syntax.StatementsNode)
	elseStatementsNode := elseStatements.node.(*syntax.StatementsNode)
	node := syntax.IfNode{
		Expression:     expr.node,
		IfStatements:   ifStatementsNode,
		ElseStatements: elseStatementsNode,
	}
	return yySymType{l: expr.l, c: expr.c, node: node}
}

func createWhileNode(expr yySymType, statements yySymType) yySymType {
	statementsNode := statements.node.(*syntax.StatementsNode)
	node := syntax.WhileNode{
		Expression: expr.node,
		Statements: statementsNode,
	}
	return yySymType{l: expr.l, c: expr.c, node: node}
}

func createForNode(initExpr yySymType, testExpr yySymType, incrExpr yySymType, statements yySymType) yySymType {
	statementsNode := statements.node.(*syntax.StatementsNode)
	node := syntax.ForNode{
		InitExpr:   initExpr.node,
		TestExpr:   testExpr.node,
		IncrExpr:   incrExpr.node,
		Statements: statementsNode,
	}
	return yySymType{l: initExpr.l, c: initExpr.c, node: node}
}

func createBreakNode(o yySymType) yySymType {
	node := syntax.BreakNode{
		Line:   o.l,
		Column: o.c,
	}
	return yySymType{l: o.l, c: o.c, node: node}
}

func createContinueNode(o yySymType) yySymType {
	node := syntax.ContinueNode{
		Line:   o.l,
		Column: o.c,
	}
	return yySymType{l: o.l, c: o.c, node: node}
}

func appendStatement(statements *yySymType, statement yySymType) yySymType {
	var node *syntax.StatementsNode
	if statements != nil {
		if statementsNode, ok := statements.node.(*syntax.StatementsNode); ok {
			node = statementsNode
		} else {
			panic(fmt.Sprintf("not a statements node: %+v", statements.node))
		}
	} else {
		node = &syntax.StatementsNode{
			Line:   statement.l,
			Column: statement.c,
		}
	}
	node.Append(statement.node)
	return yySymType{l: node.Line, c: node.Column, node: node}
}

func appendString(strings *yySymType, value yySymType) yySymType {
	var ss []string
	l, c := value.l, value.c
	if strings != nil {
		ss = strings.ss
		l = strings.l
		c = strings.c
	}
	ss = append(ss, value.s)
	return yySymType{l: l, c: c, ss: ss}
}

func appendNode(nodes *yySymType, value yySymType) yySymType {
	var ns []syntax.Node
	l, c := value.l, value.c
	if nodes != nil {
		ns = nodes.nodes
		l = nodes.l
		c = nodes.c
	}
	ns = append(ns, value.node)
	return yySymType{l: l, c: c, nodes: ns}
}

func createFunctionNode(parameters *yySymType, statements yySymType) yySymType {
	statementsNode := statements.node.(*syntax.StatementsNode)
	node := syntax.FunctionNode{
		Statements: statementsNode,
	}
	if parameters != nil {
		node.Parameters = parameters.ss
	}
	return yySymType{node: node}
}

func createReturnNode(expr yySymType) yySymType {
	node := syntax.ReturnNode{
		Expression: expr.node,
	}
	return yySymType{node: node}
}

func createCallNode(expr yySymType, args *yySymType) yySymType {
	node := syntax.CallNode{
		Expression: expr.node,
	}
	if args != nil {
		node.Arguments = args.nodes
	}
	return yySymType{node: node}
}
