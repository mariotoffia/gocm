// Code generated from cmselect.g4 by ANTLR 4.12.0. DO NOT EDIT.

package cmselect // cmselect
import (
	"fmt"
	"strconv"
  "sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}


type cmselectParser struct {
	*antlr.BaseParser
}

var cmselectParserStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  literalNames           []string
  symbolicNames          []string
  ruleNames              []string
  predictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func cmselectParserInit() {
  staticData := &cmselectParserStaticData
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
    "query", "selectClause", "attributeList", "attribute", "fromClause", 
    "tableName", "whereClause", "condition", "parameter", "logicalOperator",
  }
  staticData.predictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 18, 72, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 0, 1, 
	0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 3, 2, 30, 8, 2, 1, 2, 1, 2, 
	1, 2, 3, 2, 35, 8, 2, 5, 2, 37, 8, 2, 10, 2, 12, 2, 40, 9, 2, 1, 3, 1, 
	3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 54, 
	8, 6, 10, 6, 12, 6, 57, 9, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 64, 8, 
	7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 0, 0, 10, 0, 2, 4, 6, 8, 10, 
	12, 14, 16, 18, 0, 2, 1, 0, 4, 8, 1, 0, 9, 11, 67, 0, 20, 1, 0, 0, 0, 2, 
	24, 1, 0, 0, 0, 4, 29, 1, 0, 0, 0, 6, 41, 1, 0, 0, 0, 8, 43, 1, 0, 0, 0, 
	10, 46, 1, 0, 0, 0, 12, 48, 1, 0, 0, 0, 14, 58, 1, 0, 0, 0, 16, 65, 1, 
	0, 0, 0, 18, 69, 1, 0, 0, 0, 20, 21, 3, 2, 1, 0, 21, 22, 3, 8, 4, 0, 22, 
	23, 3, 12, 6, 0, 23, 1, 1, 0, 0, 0, 24, 25, 5, 12, 0, 0, 25, 26, 3, 4, 
	2, 0, 26, 3, 1, 0, 0, 0, 27, 30, 3, 6, 3, 0, 28, 30, 5, 1, 0, 0, 29, 27, 
	1, 0, 0, 0, 29, 28, 1, 0, 0, 0, 30, 38, 1, 0, 0, 0, 31, 34, 5, 2, 0, 0, 
	32, 35, 3, 6, 3, 0, 33, 35, 5, 1, 0, 0, 34, 32, 1, 0, 0, 0, 34, 33, 1, 
	0, 0, 0, 35, 37, 1, 0, 0, 0, 36, 31, 1, 0, 0, 0, 37, 40, 1, 0, 0, 0, 38, 
	36, 1, 0, 0, 0, 38, 39, 1, 0, 0, 0, 39, 5, 1, 0, 0, 0, 40, 38, 1, 0, 0, 
	0, 41, 42, 5, 17, 0, 0, 42, 7, 1, 0, 0, 0, 43, 44, 5, 13, 0, 0, 44, 45, 
	3, 10, 5, 0, 45, 9, 1, 0, 0, 0, 46, 47, 5, 17, 0, 0, 47, 11, 1, 0, 0, 0, 
	48, 49, 5, 14, 0, 0, 49, 55, 3, 14, 7, 0, 50, 51, 3, 18, 9, 0, 51, 52, 
	3, 14, 7, 0, 52, 54, 1, 0, 0, 0, 53, 50, 1, 0, 0, 0, 54, 57, 1, 0, 0, 0, 
	55, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 13, 1, 0, 0, 0, 57, 55, 1, 
	0, 0, 0, 58, 59, 3, 6, 3, 0, 59, 63, 7, 0, 0, 0, 60, 64, 5, 15, 0, 0, 61, 
	64, 5, 16, 0, 0, 62, 64, 3, 16, 8, 0, 63, 60, 1, 0, 0, 0, 63, 61, 1, 0, 
	0, 0, 63, 62, 1, 0, 0, 0, 64, 15, 1, 0, 0, 0, 65, 66, 5, 3, 0, 0, 66, 67, 
	5, 17, 0, 0, 67, 68, 5, 3, 0, 0, 68, 17, 1, 0, 0, 0, 69, 70, 7, 1, 0, 0, 
	70, 19, 1, 0, 0, 0, 5, 29, 34, 38, 55, 63,
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

// cmselectParserInit initializes any static state used to implement cmselectParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewcmselectParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CmselectParserInit() {
  staticData := &cmselectParserStaticData
  staticData.once.Do(cmselectParserInit)
}

// NewcmselectParser produces a new parser instance for the optional input antlr.TokenStream.
func NewcmselectParser(input antlr.TokenStream) *cmselectParser {
	CmselectParserInit()
	this := new(cmselectParser)
	this.BaseParser = antlr.NewBaseParser(input)
  staticData := &cmselectParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "cmselect.g4"

	return this
}


// cmselectParser tokens.
const (
	cmselectParserEOF = antlr.TokenEOF
	cmselectParserT__0 = 1
	cmselectParserT__1 = 2
	cmselectParserT__2 = 3
	cmselectParserEQUAL = 4
	cmselectParserNOT_EQUAL = 5
	cmselectParserLESS_THAN = 6
	cmselectParserGREATER_THAN = 7
	cmselectParserNOT_EQUAL_ALT = 8
	cmselectParserAND = 9
	cmselectParserOR = 10
	cmselectParserNOT = 11
	cmselectParserSELECT = 12
	cmselectParserFROM = 13
	cmselectParserWHERE = 14
	cmselectParserNUMBER = 15
	cmselectParserSTRING = 16
	cmselectParserIDENTIFIER = 17
	cmselectParserWS = 18
)

// cmselectParser rules.
const (
	cmselectParserRULE_query = 0
	cmselectParserRULE_selectClause = 1
	cmselectParserRULE_attributeList = 2
	cmselectParserRULE_attribute = 3
	cmselectParserRULE_fromClause = 4
	cmselectParserRULE_tableName = 5
	cmselectParserRULE_whereClause = 6
	cmselectParserRULE_condition = 7
	cmselectParserRULE_parameter = 8
	cmselectParserRULE_logicalOperator = 9
)

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SelectClause() ISelectClauseContext
	FromClause() IFromClauseContext
	WhereClause() IWhereClauseContext

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cmselectParserRULE_query
	return p
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cmselectParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) SelectClause() ISelectClauseContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectClauseContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectClauseContext)
}

func (s *QueryContext) FromClause() IFromClauseContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFromClauseContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFromClauseContext)
}

func (s *QueryContext) WhereClause() IWhereClauseContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereClauseContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereClauseContext)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *QueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cmselectListener); ok {
		listenerT.EnterQuery(s)
	}
}

func (s *QueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cmselectListener); ok {
		listenerT.ExitQuery(s)
	}
}




func (p *cmselectParser) Query() (localctx IQueryContext) {
	this := p
	_ = this

	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, cmselectParserRULE_query)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(20)
		p.SelectClause()
	}
	{
		p.SetState(21)
		p.FromClause()
	}
	{
		p.SetState(22)
		p.WhereClause()
	}



	return localctx
}


// ISelectClauseContext is an interface to support dynamic dispatch.
type ISelectClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SELECT() antlr.TerminalNode
	AttributeList() IAttributeListContext

	// IsSelectClauseContext differentiates from other interfaces.
	IsSelectClauseContext()
}

type SelectClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectClauseContext() *SelectClauseContext {
	var p = new(SelectClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cmselectParserRULE_selectClause
	return p
}

func (*SelectClauseContext) IsSelectClauseContext() {}

func NewSelectClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectClauseContext {
	var p = new(SelectClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cmselectParserRULE_selectClause

	return p
}

func (s *SelectClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectClauseContext) SELECT() antlr.TerminalNode {
	return s.GetToken(cmselectParserSELECT, 0)
}

func (s *SelectClauseContext) AttributeList() IAttributeListContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeListContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeListContext)
}

func (s *SelectClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SelectClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cmselectListener); ok {
		listenerT.EnterSelectClause(s)
	}
}

func (s *SelectClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cmselectListener); ok {
		listenerT.ExitSelectClause(s)
	}
}




func (p *cmselectParser) SelectClause() (localctx ISelectClauseContext) {
	this := p
	_ = this

	localctx = NewSelectClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, cmselectParserRULE_selectClause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(24)
		p.Match(cmselectParserSELECT)
	}
	{
		p.SetState(25)
		p.AttributeList()
	}



	return localctx
}


// IAttributeListContext is an interface to support dynamic dispatch.
type IAttributeListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAttribute() []IAttributeContext
	Attribute(i int) IAttributeContext

	// IsAttributeListContext differentiates from other interfaces.
	IsAttributeListContext()
}

type AttributeListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeListContext() *AttributeListContext {
	var p = new(AttributeListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cmselectParserRULE_attributeList
	return p
}

func (*AttributeListContext) IsAttributeListContext() {}

func NewAttributeListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeListContext {
	var p = new(AttributeListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cmselectParserRULE_attributeList

	return p
}

func (s *AttributeListContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeListContext) AllAttribute() []IAttributeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttributeContext); ok {
			len++
		}
	}

	tst := make([]IAttributeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttributeContext); ok {
			tst[i] = t.(IAttributeContext)
			i++
		}
	}

	return tst
}

func (s *AttributeListContext) Attribute(i int) IAttributeContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *AttributeListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AttributeListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cmselectListener); ok {
		listenerT.EnterAttributeList(s)
	}
}

func (s *AttributeListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cmselectListener); ok {
		listenerT.ExitAttributeList(s)
	}
}




func (p *cmselectParser) AttributeList() (localctx IAttributeListContext) {
	this := p
	_ = this

	localctx = NewAttributeListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, cmselectParserRULE_attributeList)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(29)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case cmselectParserIDENTIFIER:
		{
			p.SetState(27)
			p.Attribute()
		}


	case cmselectParserT__0:
		{
			p.SetState(28)
			p.Match(cmselectParserT__0)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == cmselectParserT__1 {
		{
			p.SetState(31)
			p.Match(cmselectParserT__1)
		}
		p.SetState(34)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case cmselectParserIDENTIFIER:
			{
				p.SetState(32)
				p.Attribute()
			}


		case cmselectParserT__0:
			{
				p.SetState(33)
				p.Match(cmselectParserT__0)
			}



		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}


		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// IAttributeContext is an interface to support dynamic dispatch.
type IAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsAttributeContext differentiates from other interfaces.
	IsAttributeContext()
}

type AttributeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeContext() *AttributeContext {
	var p = new(AttributeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cmselectParserRULE_attribute
	return p
}

func (*AttributeContext) IsAttributeContext() {}

func NewAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeContext {
	var p = new(AttributeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cmselectParserRULE_attribute

	return p
}

func (s *AttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(cmselectParserIDENTIFIER, 0)
}

func (s *AttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AttributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cmselectListener); ok {
		listenerT.EnterAttribute(s)
	}
}

func (s *AttributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cmselectListener); ok {
		listenerT.ExitAttribute(s)
	}
}




func (p *cmselectParser) Attribute() (localctx IAttributeContext) {
	this := p
	_ = this

	localctx = NewAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, cmselectParserRULE_attribute)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(41)
		p.Match(cmselectParserIDENTIFIER)
	}



	return localctx
}


// IFromClauseContext is an interface to support dynamic dispatch.
type IFromClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FROM() antlr.TerminalNode
	TableName() ITableNameContext

	// IsFromClauseContext differentiates from other interfaces.
	IsFromClauseContext()
}

type FromClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFromClauseContext() *FromClauseContext {
	var p = new(FromClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cmselectParserRULE_fromClause
	return p
}

func (*FromClauseContext) IsFromClauseContext() {}

func NewFromClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FromClauseContext {
	var p = new(FromClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cmselectParserRULE_fromClause

	return p
}

func (s *FromClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *FromClauseContext) FROM() antlr.TerminalNode {
	return s.GetToken(cmselectParserFROM, 0)
}

func (s *FromClauseContext) TableName() ITableNameContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableNameContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *FromClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FromClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FromClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cmselectListener); ok {
		listenerT.EnterFromClause(s)
	}
}

func (s *FromClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cmselectListener); ok {
		listenerT.ExitFromClause(s)
	}
}




func (p *cmselectParser) FromClause() (localctx IFromClauseContext) {
	this := p
	_ = this

	localctx = NewFromClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, cmselectParserRULE_fromClause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)
		p.Match(cmselectParserFROM)
	}
	{
		p.SetState(44)
		p.TableName()
	}



	return localctx
}


// ITableNameContext is an interface to support dynamic dispatch.
type ITableNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsTableNameContext differentiates from other interfaces.
	IsTableNameContext()
}

type TableNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableNameContext() *TableNameContext {
	var p = new(TableNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cmselectParserRULE_tableName
	return p
}

func (*TableNameContext) IsTableNameContext() {}

func NewTableNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableNameContext {
	var p = new(TableNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cmselectParserRULE_tableName

	return p
}

func (s *TableNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TableNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(cmselectParserIDENTIFIER, 0)
}

func (s *TableNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TableNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cmselectListener); ok {
		listenerT.EnterTableName(s)
	}
}

func (s *TableNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cmselectListener); ok {
		listenerT.ExitTableName(s)
	}
}




func (p *cmselectParser) TableName() (localctx ITableNameContext) {
	this := p
	_ = this

	localctx = NewTableNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, cmselectParserRULE_tableName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.Match(cmselectParserIDENTIFIER)
	}



	return localctx
}


// IWhereClauseContext is an interface to support dynamic dispatch.
type IWhereClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHERE() antlr.TerminalNode
	AllCondition() []IConditionContext
	Condition(i int) IConditionContext
	AllLogicalOperator() []ILogicalOperatorContext
	LogicalOperator(i int) ILogicalOperatorContext

	// IsWhereClauseContext differentiates from other interfaces.
	IsWhereClauseContext()
}

type WhereClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhereClauseContext() *WhereClauseContext {
	var p = new(WhereClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cmselectParserRULE_whereClause
	return p
}

func (*WhereClauseContext) IsWhereClauseContext() {}

func NewWhereClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereClauseContext {
	var p = new(WhereClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cmselectParserRULE_whereClause

	return p
}

func (s *WhereClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereClauseContext) WHERE() antlr.TerminalNode {
	return s.GetToken(cmselectParserWHERE, 0)
}

func (s *WhereClauseContext) AllCondition() []IConditionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionContext); ok {
			len++
		}
	}

	tst := make([]IConditionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionContext); ok {
			tst[i] = t.(IConditionContext)
			i++
		}
	}

	return tst
}

func (s *WhereClauseContext) Condition(i int) IConditionContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *WhereClauseContext) AllLogicalOperator() []ILogicalOperatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILogicalOperatorContext); ok {
			len++
		}
	}

	tst := make([]ILogicalOperatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILogicalOperatorContext); ok {
			tst[i] = t.(ILogicalOperatorContext)
			i++
		}
	}

	return tst
}

func (s *WhereClauseContext) LogicalOperator(i int) ILogicalOperatorContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalOperatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalOperatorContext)
}

func (s *WhereClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhereClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *WhereClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cmselectListener); ok {
		listenerT.EnterWhereClause(s)
	}
}

func (s *WhereClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cmselectListener); ok {
		listenerT.ExitWhereClause(s)
	}
}




func (p *cmselectParser) WhereClause() (localctx IWhereClauseContext) {
	this := p
	_ = this

	localctx = NewWhereClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, cmselectParserRULE_whereClause)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Match(cmselectParserWHERE)
	}
	{
		p.SetState(49)
		p.Condition()
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 3584) != 0) {
		{
			p.SetState(50)
			p.LogicalOperator()
		}
		{
			p.SetState(51)
			p.Condition()
		}


		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Attribute() IAttributeContext
	EQUAL() antlr.TerminalNode
	NOT_EQUAL() antlr.TerminalNode
	LESS_THAN() antlr.TerminalNode
	GREATER_THAN() antlr.TerminalNode
	NOT_EQUAL_ALT() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	STRING() antlr.TerminalNode
	Parameter() IParameterContext

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cmselectParserRULE_condition
	return p
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cmselectParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) Attribute() IAttributeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *ConditionContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(cmselectParserEQUAL, 0)
}

func (s *ConditionContext) NOT_EQUAL() antlr.TerminalNode {
	return s.GetToken(cmselectParserNOT_EQUAL, 0)
}

func (s *ConditionContext) LESS_THAN() antlr.TerminalNode {
	return s.GetToken(cmselectParserLESS_THAN, 0)
}

func (s *ConditionContext) GREATER_THAN() antlr.TerminalNode {
	return s.GetToken(cmselectParserGREATER_THAN, 0)
}

func (s *ConditionContext) NOT_EQUAL_ALT() antlr.TerminalNode {
	return s.GetToken(cmselectParserNOT_EQUAL_ALT, 0)
}

func (s *ConditionContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(cmselectParserNUMBER, 0)
}

func (s *ConditionContext) STRING() antlr.TerminalNode {
	return s.GetToken(cmselectParserSTRING, 0)
}

func (s *ConditionContext) Parameter() IParameterContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cmselectListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cmselectListener); ok {
		listenerT.ExitCondition(s)
	}
}




func (p *cmselectParser) Condition() (localctx IConditionContext) {
	this := p
	_ = this

	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, cmselectParserRULE_condition)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.Attribute()
	}
	{
		p.SetState(59)
		_la = p.GetTokenStream().LA(1)

		if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 496) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(63)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case cmselectParserNUMBER:
		{
			p.SetState(60)
			p.Match(cmselectParserNUMBER)
		}


	case cmselectParserSTRING:
		{
			p.SetState(61)
			p.Match(cmselectParserSTRING)
		}


	case cmselectParserT__2:
		{
			p.SetState(62)
			p.Parameter()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}



	return localctx
}


// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cmselectParserRULE_parameter
	return p
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cmselectParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(cmselectParserIDENTIFIER, 0)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cmselectListener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cmselectListener); ok {
		listenerT.ExitParameter(s)
	}
}




func (p *cmselectParser) Parameter() (localctx IParameterContext) {
	this := p
	_ = this

	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, cmselectParserRULE_parameter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(65)
		p.Match(cmselectParserT__2)
	}
	{
		p.SetState(66)
		p.Match(cmselectParserIDENTIFIER)
	}
	{
		p.SetState(67)
		p.Match(cmselectParserT__2)
	}



	return localctx
}


// ILogicalOperatorContext is an interface to support dynamic dispatch.
type ILogicalOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode
	NOT() antlr.TerminalNode

	// IsLogicalOperatorContext differentiates from other interfaces.
	IsLogicalOperatorContext()
}

type LogicalOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalOperatorContext() *LogicalOperatorContext {
	var p = new(LogicalOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cmselectParserRULE_logicalOperator
	return p
}

func (*LogicalOperatorContext) IsLogicalOperatorContext() {}

func NewLogicalOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalOperatorContext {
	var p = new(LogicalOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cmselectParserRULE_logicalOperator

	return p
}

func (s *LogicalOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalOperatorContext) AND() antlr.TerminalNode {
	return s.GetToken(cmselectParserAND, 0)
}

func (s *LogicalOperatorContext) OR() antlr.TerminalNode {
	return s.GetToken(cmselectParserOR, 0)
}

func (s *LogicalOperatorContext) NOT() antlr.TerminalNode {
	return s.GetToken(cmselectParserNOT, 0)
}

func (s *LogicalOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LogicalOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cmselectListener); ok {
		listenerT.EnterLogicalOperator(s)
	}
}

func (s *LogicalOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cmselectListener); ok {
		listenerT.ExitLogicalOperator(s)
	}
}




func (p *cmselectParser) LogicalOperator() (localctx ILogicalOperatorContext) {
	this := p
	_ = this

	localctx = NewLogicalOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, cmselectParserRULE_logicalOperator)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		_la = p.GetTokenStream().LA(1)

		if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 3584) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


