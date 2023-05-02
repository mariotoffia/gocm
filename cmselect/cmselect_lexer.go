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
    "", "'*'", "','", "':'", "'=='", "'<>'", "'<'", "'>'", "'!='", "'AND'", 
    "'OR'", "'NOT'", "'SELECT'", "'FROM'", "'WHERE'",
  }
  staticData.symbolicNames = []string{
    "", "", "", "", "EQUAL", "NOT_EQUAL", "LESS_THAN", "GREATER_THAN", "NOT_EQUAL_ALT", 
    "AND", "OR", "NOT", "SELECT", "FROM", "WHERE", "NUMBER", "STRING", "IDENTIFIER", 
    "WS",
  }
  staticData.ruleNames = []string{
    "T__0", "T__1", "T__2", "EQUAL", "NOT_EQUAL", "LESS_THAN", "GREATER_THAN", 
    "NOT_EQUAL_ALT", "AND", "OR", "NOT", "SELECT", "FROM", "WHERE", "NUMBER", 
    "STRING", "IDENTIFIER", "LETTER", "DIGIT", "WS",
  }
  staticData.predictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 0, 18, 136, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 
	4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 
	10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 
	7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 1, 0, 1, 
	0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 
	5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 
	9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 
	1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 
	13, 1, 13, 1, 14, 3, 14, 91, 8, 14, 1, 14, 4, 14, 94, 8, 14, 11, 14, 12, 
	14, 95, 1, 14, 1, 14, 4, 14, 100, 8, 14, 11, 14, 12, 14, 101, 3, 14, 104, 
	8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 110, 8, 15, 10, 15, 12, 15, 113, 
	9, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 121, 8, 16, 10, 
	16, 12, 16, 124, 9, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 4, 19, 131, 
	8, 19, 11, 19, 12, 19, 132, 1, 19, 1, 19, 0, 0, 20, 1, 1, 3, 2, 5, 3, 7, 
	4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 
	14, 29, 15, 31, 16, 33, 17, 35, 0, 37, 0, 39, 18, 1, 0, 5, 1, 0, 39, 39, 
	2, 0, 46, 46, 95, 95, 2, 0, 65, 90, 97, 122, 1, 0, 48, 57, 3, 0, 9, 10, 
	13, 13, 32, 32, 143, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 
	0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 
	0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 
	0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 
	1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 1, 
	41, 1, 0, 0, 0, 3, 43, 1, 0, 0, 0, 5, 45, 1, 0, 0, 0, 7, 47, 1, 0, 0, 0, 
	9, 50, 1, 0, 0, 0, 11, 53, 1, 0, 0, 0, 13, 55, 1, 0, 0, 0, 15, 57, 1, 0, 
	0, 0, 17, 60, 1, 0, 0, 0, 19, 64, 1, 0, 0, 0, 21, 67, 1, 0, 0, 0, 23, 71, 
	1, 0, 0, 0, 25, 78, 1, 0, 0, 0, 27, 83, 1, 0, 0, 0, 29, 90, 1, 0, 0, 0, 
	31, 105, 1, 0, 0, 0, 33, 116, 1, 0, 0, 0, 35, 125, 1, 0, 0, 0, 37, 127, 
	1, 0, 0, 0, 39, 130, 1, 0, 0, 0, 41, 42, 5, 42, 0, 0, 42, 2, 1, 0, 0, 0, 
	43, 44, 5, 44, 0, 0, 44, 4, 1, 0, 0, 0, 45, 46, 5, 58, 0, 0, 46, 6, 1, 
	0, 0, 0, 47, 48, 5, 61, 0, 0, 48, 49, 5, 61, 0, 0, 49, 8, 1, 0, 0, 0, 50, 
	51, 5, 60, 0, 0, 51, 52, 5, 62, 0, 0, 52, 10, 1, 0, 0, 0, 53, 54, 5, 60, 
	0, 0, 54, 12, 1, 0, 0, 0, 55, 56, 5, 62, 0, 0, 56, 14, 1, 0, 0, 0, 57, 
	58, 5, 33, 0, 0, 58, 59, 5, 61, 0, 0, 59, 16, 1, 0, 0, 0, 60, 61, 5, 65, 
	0, 0, 61, 62, 5, 78, 0, 0, 62, 63, 5, 68, 0, 0, 63, 18, 1, 0, 0, 0, 64, 
	65, 5, 79, 0, 0, 65, 66, 5, 82, 0, 0, 66, 20, 1, 0, 0, 0, 67, 68, 5, 78, 
	0, 0, 68, 69, 5, 79, 0, 0, 69, 70, 5, 84, 0, 0, 70, 22, 1, 0, 0, 0, 71, 
	72, 5, 83, 0, 0, 72, 73, 5, 69, 0, 0, 73, 74, 5, 76, 0, 0, 74, 75, 5, 69, 
	0, 0, 75, 76, 5, 67, 0, 0, 76, 77, 5, 84, 0, 0, 77, 24, 1, 0, 0, 0, 78, 
	79, 5, 70, 0, 0, 79, 80, 5, 82, 0, 0, 80, 81, 5, 79, 0, 0, 81, 82, 5, 77, 
	0, 0, 82, 26, 1, 0, 0, 0, 83, 84, 5, 87, 0, 0, 84, 85, 5, 72, 0, 0, 85, 
	86, 5, 69, 0, 0, 86, 87, 5, 82, 0, 0, 87, 88, 5, 69, 0, 0, 88, 28, 1, 0, 
	0, 0, 89, 91, 5, 45, 0, 0, 90, 89, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 
	93, 1, 0, 0, 0, 92, 94, 3, 37, 18, 0, 93, 92, 1, 0, 0, 0, 94, 95, 1, 0, 
	0, 0, 95, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 103, 1, 0, 0, 0, 97, 
	99, 5, 46, 0, 0, 98, 100, 3, 37, 18, 0, 99, 98, 1, 0, 0, 0, 100, 101, 1, 
	0, 0, 0, 101, 99, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 104, 1, 0, 0, 
	0, 103, 97, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 30, 1, 0, 0, 0, 105, 
	111, 5, 39, 0, 0, 106, 110, 8, 0, 0, 0, 107, 108, 5, 39, 0, 0, 108, 110, 
	5, 39, 0, 0, 109, 106, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 110, 113, 1, 0, 
	0, 0, 111, 109, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 114, 1, 0, 0, 0, 
	113, 111, 1, 0, 0, 0, 114, 115, 5, 39, 0, 0, 115, 32, 1, 0, 0, 0, 116, 
	122, 3, 35, 17, 0, 117, 121, 3, 35, 17, 0, 118, 121, 3, 37, 18, 0, 119, 
	121, 7, 1, 0, 0, 120, 117, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 120, 119, 
	1, 0, 0, 0, 121, 124, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 122, 123, 1, 0, 
	0, 0, 123, 34, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 125, 126, 7, 2, 0, 0, 
	126, 36, 1, 0, 0, 0, 127, 128, 7, 3, 0, 0, 128, 38, 1, 0, 0, 0, 129, 131, 
	7, 4, 0, 0, 130, 129, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 130, 1, 0, 
	0, 0, 132, 133, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 135, 6, 19, 0, 0, 
	135, 40, 1, 0, 0, 0, 10, 0, 90, 95, 101, 103, 109, 111, 120, 122, 132, 
	1, 6, 0, 0,
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
	cmselectLexerT__2 = 3
	cmselectLexerEQUAL = 4
	cmselectLexerNOT_EQUAL = 5
	cmselectLexerLESS_THAN = 6
	cmselectLexerGREATER_THAN = 7
	cmselectLexerNOT_EQUAL_ALT = 8
	cmselectLexerAND = 9
	cmselectLexerOR = 10
	cmselectLexerNOT = 11
	cmselectLexerSELECT = 12
	cmselectLexerFROM = 13
	cmselectLexerWHERE = 14
	cmselectLexerNUMBER = 15
	cmselectLexerSTRING = 16
	cmselectLexerIDENTIFIER = 17
	cmselectLexerWS = 18
)

