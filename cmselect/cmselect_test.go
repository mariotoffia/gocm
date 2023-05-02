package cmselect

import (
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func TestSimpleSQL(t *testing.T) {
	// Define a sample SQL query that matches your grammar
	input := "SELECT attribute1, attribute2 FROM tableName WHERE attribute1 == :param1: AND attribute2 != 'value'"

	// Create an ANTLR input stream from the sample SQL query
	inputStream := antlr.NewInputStream(input)

	// Create a lexer and parser for the input stream
	lexer := NewcmselectLexer(inputStream)
	parser := NewcmselectParser(antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel))

	// Parse the query using the entry point rule in your grammar
	tree := parser.Query()

	// Check if the parsing was successful
	if tree.GetRuleIndex() == -1 {
		t.Errorf("Failed to parse input: %s", input)
	}
}
