// Code generated from cmselect.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser

import (
	"fmt"
  "sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter


type cmselectLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var cmselectlexerLexerStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  channelNames           []string
  modeNames              []string
  literalNames           []string
  symbolicNames          []string
  ruleNames              []string
  predictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func cmselectlexerLexerInit() {
  staticData := &cmselectlexerLexerStaticData
  staticData.channelNames = []string{
    "DEFAULT_TOKEN_CHANNEL", "HIDDEN",
  }
  staticData.modeNames = []string{
    "DEFAULT_MODE",
  }
  staticData.literalNames = []string{
    "", "'*'", "','", "'=='", "'<>'", "'<'", "'>'", "'!='", "'AND'", "'OR'", 
    "'NOT'", "'SELECT'", "'FROM'", "'WHERE'",
  }
  staticData.symbolicNames = []string{
    "", "", "", "EQUAL", "NOT_EQUAL", "LESS_THAN", "GREATER_THAN", "NOT_EQUAL_ALT", 
    "AND", "OR", "NOT", "SELECT", "FROM", "WHERE", "NUMBER", "STRING", "IDENTIFIER", 
    "WS",
  }
  staticData.ruleNames = []string{
    "T__0", "T__1", "EQUAL", "NOT_EQUAL", "LESS_THAN", "GREATER_THAN", "NOT_EQUAL_ALT", 
    "AND", "OR", "NOT", "SELECT", "FROM", "WHERE", "NUMBER", "STRING", "IDENTIFIER", 
    "LETTER", "DIGIT", "WS",
  }
  staticData.predictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 0, 17, 132, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 
	4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 
	10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 
	7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 1, 0, 1, 0, 1, 1, 1, 1, 
	1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 
	1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 
	1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 
	11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 3, 13, 87, 
	8, 13, 1, 13, 4, 13, 90, 8, 13, 11, 13, 12, 13, 91, 1, 13, 1, 13, 4, 13, 
	96, 8, 13, 11, 13, 12, 13, 97, 3, 13, 100, 8, 13, 1, 14, 1, 14, 1, 14, 
	1, 14, 5, 14, 106, 8, 14, 10, 14, 12, 14, 109, 9, 14, 1, 14, 1, 14, 1, 
	15, 1, 15, 1, 15, 1, 15, 5, 15, 117, 8, 15, 10, 15, 12, 15, 120, 9, 15, 
	1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 4, 18, 127, 8, 18, 11, 18, 12, 18, 128, 
	1, 18, 1, 18, 0, 0, 19, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 
	8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 0, 
	35, 0, 37, 17, 1, 0, 5, 1, 0, 39, 39, 2, 0, 46, 46, 95, 95, 2, 0, 65, 90, 
	97, 122, 1, 0, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32, 139, 0, 1, 1, 0, 0, 
	0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 
	0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 
	0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 
	0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 37, 
	1, 0, 0, 0, 1, 39, 1, 0, 0, 0, 3, 41, 1, 0, 0, 0, 5, 43, 1, 0, 0, 0, 7, 
	46, 1, 0, 0, 0, 9, 49, 1, 0, 0, 0, 11, 51, 1, 0, 0, 0, 13, 53, 1, 0, 0, 
	0, 15, 56, 1, 0, 0, 0, 17, 60, 1, 0, 0, 0, 19, 63, 1, 0, 0, 0, 21, 67, 
	1, 0, 0, 0, 23, 74, 1, 0, 0, 0, 25, 79, 1, 0, 0, 0, 27, 86, 1, 0, 0, 0, 
	29, 101, 1, 0, 0, 0, 31, 112, 1, 0, 0, 0, 33, 121, 1, 0, 0, 0, 35, 123, 
	1, 0, 0, 0, 37, 126, 1, 0, 0, 0, 39, 40, 5, 42, 0, 0, 40, 2, 1, 0, 0, 0, 
	41, 42, 5, 44, 0, 0, 42, 4, 1, 0, 0, 0, 43, 44, 5, 61, 0, 0, 44, 45, 5, 
	61, 0, 0, 45, 6, 1, 0, 0, 0, 46, 47, 5, 60, 0, 0, 47, 48, 5, 62, 0, 0, 
	48, 8, 1, 0, 0, 0, 49, 50, 5, 60, 0, 0, 50, 10, 1, 0, 0, 0, 51, 52, 5, 
	62, 0, 0, 52, 12, 1, 0, 0, 0, 53, 54, 5, 33, 0, 0, 54, 55, 5, 61, 0, 0, 
	55, 14, 1, 0, 0, 0, 56, 57, 5, 65, 0, 0, 57, 58, 5, 78, 0, 0, 58, 59, 5, 
	68, 0, 0, 59, 16, 1, 0, 0, 0, 60, 61, 5, 79, 0, 0, 61, 62, 5, 82, 0, 0, 
	62, 18, 1, 0, 0, 0, 63, 64, 5, 78, 0, 0, 64, 65, 5, 79, 0, 0, 65, 66, 5, 
	84, 0, 0, 66, 20, 1, 0, 0, 0, 67, 68, 5, 83, 0, 0, 68, 69, 5, 69, 0, 0, 
	69, 70, 5, 76, 0, 0, 70, 71, 5, 69, 0, 0, 71, 72, 5, 67, 0, 0, 72, 73, 
	5, 84, 0, 0, 73, 22, 1, 0, 0, 0, 74, 75, 5, 70, 0, 0, 75, 76, 5, 82, 0, 
	0, 76, 77, 5, 79, 0, 0, 77, 78, 5, 77, 0, 0, 78, 24, 1, 0, 0, 0, 79, 80, 
	5, 87, 0, 0, 80, 81, 5, 72, 0, 0, 81, 82, 5, 69, 0, 0, 82, 83, 5, 82, 0, 
	0, 83, 84, 5, 69, 0, 0, 84, 26, 1, 0, 0, 0, 85, 87, 5, 45, 0, 0, 86, 85, 
	1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 89, 1, 0, 0, 0, 88, 90, 3, 35, 17, 
	0, 89, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 91, 92, 
	1, 0, 0, 0, 92, 99, 1, 0, 0, 0, 93, 95, 5, 46, 0, 0, 94, 96, 3, 35, 17, 
	0, 95, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 97, 98, 
	1, 0, 0, 0, 98, 100, 1, 0, 0, 0, 99, 93, 1, 0, 0, 0, 99, 100, 1, 0, 0, 
	0, 100, 28, 1, 0, 0, 0, 101, 107, 5, 39, 0, 0, 102, 106, 8, 0, 0, 0, 103, 
	104, 5, 39, 0, 0, 104, 106, 5, 39, 0, 0, 105, 102, 1, 0, 0, 0, 105, 103, 
	1, 0, 0, 0, 106, 109, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 107, 108, 1, 0, 
	0, 0, 108, 110, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 110, 111, 5, 39, 0, 0, 
	111, 30, 1, 0, 0, 0, 112, 118, 3, 33, 16, 0, 113, 117, 3, 33, 16, 0, 114, 
	117, 3, 35, 17, 0, 115, 117, 7, 1, 0, 0, 116, 113, 1, 0, 0, 0, 116, 114, 
	1, 0, 0, 0, 116, 115, 1, 0, 0, 0, 117, 120, 1, 0, 0, 0, 118, 116, 1, 0, 
	0, 0, 118, 119, 1, 0, 0, 0, 119, 32, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 
	121, 122, 7, 2, 0, 0, 122, 34, 1, 0, 0, 0, 123, 124, 7, 3, 0, 0, 124, 36, 
	1, 0, 0, 0, 125, 127, 7, 4, 0, 0, 126, 125, 1, 0, 0, 0, 127, 128, 1, 0, 
	0, 0, 128, 126, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 
	130, 131, 6, 18, 0, 0, 131, 38, 1, 0, 0, 0, 10, 0, 86, 91, 97, 99, 105, 
	107, 116, 118, 128, 1, 6, 0, 0,
}
  deserializer := antlr.NewATNDeserializer(nil)
  staticData.atn = deserializer.Deserialize(staticData.serializedATN)
  atn := staticData.atn
  staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
  decisionToDFA := staticData.decisionToDFA
  for index, state := range atn.DecisionToState {
    decisionToDFA[index] = antlr.NewDFA(state, index)
  }
}

// cmselectLexerInit initializes any static state used to implement cmselectLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewcmselectLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CmselectLexerInit() {
  staticData := &cmselectlexerLexerStaticData
  staticData.once.Do(cmselectlexerLexerInit)
}

// NewcmselectLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewcmselectLexer(input antlr.CharStream) *cmselectLexer {
  CmselectLexerInit()
	l := new(cmselectLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
  staticData := &cmselectlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "cmselect.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// cmselectLexer tokens.
const (
	cmselectLexerT__0 = 1
	cmselectLexerT__1 = 2
	cmselectLexerEQUAL = 3
	cmselectLexerNOT_EQUAL = 4
	cmselectLexerLESS_THAN = 5
	cmselectLexerGREATER_THAN = 6
	cmselectLexerNOT_EQUAL_ALT = 7
	cmselectLexerAND = 8
	cmselectLexerOR = 9
	cmselectLexerNOT = 10
	cmselectLexerSELECT = 11
	cmselectLexerFROM = 12
	cmselectLexerWHERE = 13
	cmselectLexerNUMBER = 14
	cmselectLexerSTRING = 15
	cmselectLexerIDENTIFIER = 16
	cmselectLexerWS = 17
)

