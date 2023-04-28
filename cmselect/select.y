%{
package cmselect

import "fmt"
%}

%token IDENTIFIER STRING_LITERAL NUMBER_LITERAL
%token SELECT FROM WHERE AND OR NOT EQUAL NOTEQUAL LESSTHAN GREATERTHAN
%left OR
%left AND
%left NOT
%left EQUAL NOTEQUAL
%left LESSTHAN GREATERTHAN

%type <exprNode> Expression WhereClause
%type <cmpNode> Comparison
%type <str> IDENTIFIER STRING_LITERAL NUMBER_LITERAL
%type <projNode> Projection
%type <selectNode> SelectStatement

%%

Expression:
    SelectStatement { $$ = $1 }
    ;

Projection:
    IDENTIFIER { $$ = &ProjectionNode{Field: $1} }
    ;

WhereClause:
    Comparison { $$ = $1 }
    | WhereClause AND Comparison { $$ = &LogicalNode{Op: "AND", Left: $1, Right: $3} }
    | WhereClause OR Comparison { $$ = &LogicalNode{Op: "OR", Left: $1, Right: $3} }
    | NOT WhereClause { $$ = &LogicalNode{Op: "NOT", Operand: $2} }
    ;

Comparison:
    Projection EQUAL Projection { $$ = &ComparisonNode{Op: "==", Left: $1, Right: $3} }
    | Projection NOTEQUAL Projection { $$ = &ComparisonNode{Op: "!=", Left: $1, Right: $3} }
    | Projection LESSTHAN Projection { $$ = &ComparisonNode{Op: "<", Left: $1, Right: $3} }
    | Projection GREATERTHAN Projection { $$ = &ComparisonNode{Op: ">", Left: $1, Right: $3} }
    ;

SelectStatement:
    SELECT Projection FROM IDENTIFIER WhereClause { $$ = &SelectNode{Projection: $2, TableName: $4, Where: $5} }
    ;

%%

type exprNode interface{}

type ProjectionNode struct {
    Field string
}

type ComparisonNode struct {
    Op    string
    Left  *ProjectionNode
    Right *ProjectionNode
}

type LogicalNode struct {
    Op      string
    Left    exprNode
    Right   exprNode
    Operand exprNode
}

type SelectNode struct {
    Projection *ProjectionNode
    TableName  string
    Where      exprNode
}

func main() {
    yyParse(newLexer("SELECT column1 FROM table1 WHERE column1 == 'some_string' AND column2 != 42"))
}
