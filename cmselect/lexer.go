package cmselect

import (
	"fmt"
	"strings"

	"github.com/cznic/lex"
)

const (
	IDENTIFIER = iota + 1
	STRING_LITERAL
	NUMBER_LITERAL
	SELECT
	FROM
	WHERE
	AND
	OR
	NOT
	EQUAL
	NOTEQUAL
	LESSTHAN
	GREATERTHAN
)

var tokens = []string{
	"EOF",
	"IDENTIFIER",
	"STRING_LITERAL",
	"NUMBER_LITERAL",
	"SELECT",
	"FROM",
	"WHERE",
	"AND",
	"OR",
	"NOT",
	"EQUAL",
	"NOTEQUAL",
	"LESSTHAN",
	"GREATERTHAN",
}

var lexRules = `
	\s+                   ;
	[a-zA-Z_][a-zA-Z0-9_]* { return yy.Symbol(yylex, IDENTIFIER, "IDENTIFIER") }
	\d+(\.\d+)?            { return yy.Symbol(yylex, NUMBER_LITERAL, "NUMBER_LITERAL") }
	"'"[^']*"'|"[^"]*"     { return yy.Symbol(yylex, STRING_LITERAL, "STRING_LITERAL") }
	select                 { return SELECT }
	from                   { return FROM }
	where                  { return WHERE }
	and                    { return AND }
	or                     { return OR }
	not                    { return NOT }
	==                     { return EQUAL }
	!=                     { return NOTEQUAL }
	<                      { return LESSTHAN }
	>                      { return GREATERTHAN }
	.                      { return int(yylex.Text()[0]) }
`

type lexer struct {
	*lex.Lexer
}

func (l *lexer) Error(s string) {
	fmt.Printf("syntax error: %s\n", s)
}

func newLexer(input string) *lexer {
	l := &lexer{
		Lexer: lex.New(strings.NewReader(input), lexRules),
	}
	return l
}

func (l *lexer) Symbol(lex *lex.Lexer, sym int, typ string) int {
	lval := l.Value().(*yySymType)
	lval.str = typ
	lval.strVal = l.Text()
	return sym
}
