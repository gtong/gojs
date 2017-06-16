%{
package parser

import "github.com/gtong/gojs/syntax"
%}

%union {
  s string
  node syntax.Node
}

%token BOOLEAN
%token NUMBER
%token STRING
%token IDENTIFIER

%token UNARY_OP
%token BIN_OP_1
%token BIN_OP_2
%token ASSIGNMENT

%token END
%token LP
%token RP
%token LB
%token RB
%token IF
%token ELSE
%token WHILE
%token FOR
%token BREAK
%token CONTINUE
%token FUNCTION
%token RETURN

%right ASSIGNMENT
%right RETURN
%left LP
%left RP
%left UNARY_OP
%left BIN_OP_1
%left BIN_OP_2

%%
program: statements
{
  setParseResult(yylex, $1)
}

statements: statement
{
  $$ = appendStatement(nil, $1)
}
| statements statement
{
  $$ = appendStatement(&$1, $2)
}

statement: expr END
{
  $$ = $1
}
| IF LP expr RP LB statements RB
{
  $$ = createIfNode($3, $6)
}
| IF LP expr RP LB statements RB ELSE LB statements RB
{
  $$ = createIfElseNode($3, $6, $10)
}
| WHILE LP expr RP LB statements RB
{
  $$ = createWhileNode($3, $6)
}
| FOR LP expr END expr END expr RP LB statements RB
{
  $$ = createForNode($3, $5, $7, $10)
}

expr: BOOLEAN
{
  $$ = createBooleanNode($1)
} | NUMBER
{
  $$ = createNumberNode($1)
}
| STRING
{
  $$ = createStringNode($1)
}
| IDENTIFIER
{
  $$ = createIdentifierNode($1)
}
| BREAK
{
  $$ = createBreakNode()
}
| CONTINUE
{
  $$ = createContinueNode()
}
| expr BIN_OP_2 expr
{
  $$ = createBinaryOpNode($2, $1, $3)
}
| expr BIN_OP_1 expr
{
  $$ = createBinaryOpNode($2, $1, $3)
}
| expr ASSIGNMENT expr
{
  $$ = createBinaryOpNode($2, $1, $3)
}
| expr UNARY_OP
{
  $$ = createUnaryOpNode($2, $1)
}
| FUNCTION LP RP LB statements RB
{
  $$ = createFunctionNode($5)
}
| RETURN expr
{
  $$ = createReturnNode($2)
}
| expr LP RP
{
  $$ = createCallNode($1)
}
%%