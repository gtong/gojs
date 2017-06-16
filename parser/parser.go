package parser

import (
	"io"

	"github.com/gtong/gojs/syntax"
)

func Parse(reader io.Reader) syntax.Node {
	lexer := NewLexer(reader)
	yyParse(lexer)
	return lexer.parseResult.(yySymType).node
}

func setParseResult(lexer interface{}, o yySymType) {
	lexer.(*Lexer).parseResult = o
}
