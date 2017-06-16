package parser

import (
	"fmt"
	"io"

	"github.com/gtong/gojs/syntax"
)

func createNumberNode(o yySymType) yySymType {
	node := syntax.NumberNode{
		Value: o.s,
	}
	return yySymType{node: node}
}

func createBooleanNode(o yySymType) yySymType {
	node := syntax.BooleanNode{
		Value: o.s,
	}
	return yySymType{node: node}
}

func createStringNode(o yySymType) yySymType {
	node := syntax.StringNode{
		Value: o.s,
	}
	return yySymType{node: node}
}

func createIdentifierNode(o yySymType) yySymType {
	node := syntax.IdentifierNode{
		Value: o.s,
	}
	return yySymType{node: node}
}

func createUnaryOpNode(operator yySymType, left yySymType) yySymType {
	node := syntax.UnaryOpNode{
		Left:     left.node,
		Operator: operator.s,
	}
	return yySymType{node: node}
}

func createBinaryOpNode(operator yySymType, left yySymType, right yySymType) yySymType {
	node := syntax.BinaryOpNode{
		Left:     left.node,
		Right:    right.node,
		Operator: operator.s,
	}
	return yySymType{node: node}
}

func createIfNode(expr yySymType, ifStatements yySymType) yySymType {
	ifStatementsNode := ifStatements.node.(*syntax.StatementsNode)
	node := syntax.IfNode{
		Expression:   expr.node,
		IfStatements: ifStatementsNode,
	}
	return yySymType{node: node}
}

func createIfElseNode(expr yySymType, ifStatements yySymType, elseStatements yySymType) yySymType {
	ifStatementsNode := ifStatements.node.(*syntax.StatementsNode)
	elseStatementsNode := elseStatements.node.(*syntax.StatementsNode)
	node := syntax.IfNode{
		Expression:     expr.node,
		IfStatements:   ifStatementsNode,
		ElseStatements: elseStatementsNode,
	}
	return yySymType{node: node}
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
		node = &syntax.StatementsNode{}
	}
	node.Append(statement.node)
	return yySymType{node: node}
}

func Parse(reader io.Reader) syntax.Node {
	lexer := NewLexer(reader)
	yyParse(lexer)
	return lexer.parseResult.(yySymType).node
}

func setParseResult(lexer interface{}, o yySymType) {
	lexer.(*Lexer).parseResult = o
}