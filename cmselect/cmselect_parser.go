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
    "", "','", "'*'", "':'", "'=='", "'!='", "'<'", "'>'", "'AND'", "'OR'", 
    "'NOT'", "'AS'", "'SELECT'", "'FROM'", "'WHERE'",
  }
  staticData.symbolicNames = []string{
    "", "", "", "", "EQUAL", "NOT_EQUAL", "LESS_THAN", "GREATER_THAN", "AND", 
    "OR", "NOT", "AS", "SELECT", "FROM", "WHERE", "NUMBER", "STRING", "IDENTIFIER", 
    "WS",
  }
  staticData.ruleNames = []string{
    "query", "selectClause", "attributeList", "attribute", "rhAttribute", 
    "fromClause", "tableName", "whereClause", "condition", "rightHandExpression", 
    "parameter", "logicalOperator",
  }
  staticData.predictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 18, 78, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7, 
	10, 2, 11, 7, 11, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 
	1, 2, 5, 2, 35, 8, 2, 10, 2, 12, 2, 38, 9, 2, 1, 3, 1, 3, 1, 3, 3, 3, 43, 
	8, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 
	1, 7, 5, 7, 57, 8, 7, 10, 7, 12, 7, 60, 9, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 
	9, 1, 9, 1, 9, 1, 9, 3, 9, 70, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 
	1, 11, 1, 11, 0, 0, 12, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 0, 3, 
	2, 0, 2, 2, 17, 17, 1, 0, 4, 7, 1, 0, 8, 10, 71, 0, 24, 1, 0, 0, 0, 2, 
	28, 1, 0, 0, 0, 4, 31, 1, 0, 0, 0, 6, 39, 1, 0, 0, 0, 8, 44, 1, 0, 0, 0, 
	10, 46, 1, 0, 0, 0, 12, 49, 1, 0, 0, 0, 14, 51, 1, 0, 0, 0, 16, 61, 1, 
	0, 0, 0, 18, 69, 1, 0, 0, 0, 20, 71, 1, 0, 0, 0, 22, 75, 1, 0, 0, 0, 24, 
	25, 3, 2, 1, 0, 25, 26, 3, 10, 5, 0, 26, 27, 3, 14, 7, 0, 27, 1, 1, 0, 
	0, 0, 28, 29, 5, 12, 0, 0, 29, 30, 3, 4, 2, 0, 30, 3, 1, 0, 0, 0, 31, 36, 
	3, 6, 3, 0, 32, 33, 5, 1, 0, 0, 33, 35, 3, 6, 3, 0, 34, 32, 1, 0, 0, 0, 
	35, 38, 1, 0, 0, 0, 36, 34, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 5, 1, 0, 
	0, 0, 38, 36, 1, 0, 0, 0, 39, 42, 7, 0, 0, 0, 40, 41, 5, 11, 0, 0, 41, 
	43, 5, 17, 0, 0, 42, 40, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 7, 1, 0, 0, 
	0, 44, 45, 5, 17, 0, 0, 45, 9, 1, 0, 0, 0, 46, 47, 5, 13, 0, 0, 47, 48, 
	3, 12, 6, 0, 48, 11, 1, 0, 0, 0, 49, 50, 5, 17, 0, 0, 50, 13, 1, 0, 0, 
	0, 51, 52, 5, 14, 0, 0, 52, 58, 3, 16, 8, 0, 53, 54, 3, 22, 11, 0, 54, 
	55, 3, 16, 8, 0, 55, 57, 1, 0, 0, 0, 56, 53, 1, 0, 0, 0, 57, 60, 1, 0, 
	0, 0, 58, 56, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 15, 1, 0, 0, 0, 60, 58, 
	1, 0, 0, 0, 61, 62, 3, 6, 3, 0, 62, 63, 7, 1, 0, 0, 63, 64, 3, 18, 9, 0, 
	64, 17, 1, 0, 0, 0, 65, 70, 5, 16, 0, 0, 66, 70, 5, 15, 0, 0, 67, 70, 3, 
	20, 10, 0, 68, 70, 3, 8, 4, 0, 69, 65, 1, 0, 0, 0, 69, 66, 1, 0, 0, 0, 
	69, 67, 1, 0, 0, 0, 69, 68, 1, 0, 0, 0, 70, 19, 1, 0, 0, 0, 71, 72, 5, 
	3, 0, 0, 72, 73, 5, 17, 0, 0, 73, 74, 5, 3, 0, 0, 74, 21, 1, 0, 0, 0, 75, 
	76, 7, 2, 0, 0, 76, 23, 1, 0, 0, 0, 4, 36, 42, 58, 69,
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
	cmselectParserAND = 8
	cmselectParserOR = 9
	cmselectParserNOT = 10
	cmselectParserAS = 11
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
	cmselectParserRULE_rhAttribute = 4
	cmselectParserRULE_fromClause = 5
	cmselectParserRULE_tableName = 6
	cmselectParserRULE_whereClause = 7
	cmselectParserRULE_condition = 8
	cmselectParserRULE_rightHandExpression = 9
	cmselectParserRULE_parameter = 10
	cmselectParserRULE_logicalOperator = 11
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
		p.SetState(24)
		p.SelectClause()
	}
	{
		p.SetState(25)
		p.FromClause()
	}
	{
		p.SetState(26)
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
		p.SetState(28)
		p.Match(cmselectParserSELECT)
	}
	{
		p.SetState(29)
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
	{
		p.SetState(31)
		p.Attribute()
	}
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == cmselectParserT__0 {
		{
			p.SetState(32)
			p.Match(cmselectParserT__0)
		}
		{
			p.SetState(33)
			p.Attribute()
		}


		p.SetState(38)
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
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AS() antlr.TerminalNode

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

func (s *AttributeContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(cmselectParserIDENTIFIER)
}

func (s *AttributeContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(cmselectParserIDENTIFIER, i)
}

func (s *AttributeContext) AS() antlr.TerminalNode {
	return s.GetToken(cmselectParserAS, 0)
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
		p.SetState(39)
		_la = p.GetTokenStream().LA(1)

		if !(_la == cmselectParserT__1 || _la == cmselectParserIDENTIFIER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == cmselectParserAS {
		{
			p.SetState(40)
			p.Match(cmselectParserAS)
		}
		{
			p.SetState(41)
			p.Match(cmselectParserIDENTIFIER)
		}

	}



	return localctx
}


// IRhAttributeContext is an interface to support dynamic dispatch.
type IRhAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsRhAttributeContext differentiates from other interfaces.
	IsRhAttributeContext()
}

type RhAttributeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRhAttributeContext() *RhAttributeContext {
	var p = new(RhAttributeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cmselectParserRULE_rhAttribute
	return p
}

func (*RhAttributeContext) IsRhAttributeContext() {}

func NewRhAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RhAttributeContext {
	var p = new(RhAttributeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cmselectParserRULE_rhAttribute

	return p
}

func (s *RhAttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *RhAttributeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(cmselectParserIDENTIFIER, 0)
}

func (s *RhAttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RhAttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RhAttributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cmselectListener); ok {
		listenerT.EnterRhAttribute(s)
	}
}

func (s *RhAttributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cmselectListener); ok {
		listenerT.ExitRhAttribute(s)
	}
}




func (p *cmselectParser) RhAttribute() (localctx IRhAttributeContext) {
	this := p
	_ = this

	localctx = NewRhAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, cmselectParserRULE_rhAttribute)

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
		p.SetState(44)
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
	p.EnterRule(localctx, 10, cmselectParserRULE_fromClause)

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
		p.Match(cmselectParserFROM)
	}
	{
		p.SetState(47)
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
	p.EnterRule(localctx, 12, cmselectParserRULE_tableName)

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
		p.SetState(49)
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
	p.EnterRule(localctx, 14, cmselectParserRULE_whereClause)
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
		p.SetState(51)
		p.Match(cmselectParserWHERE)
	}
	{
		p.SetState(52)
		p.Condition()
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 1792) != 0) {
		{
			p.SetState(53)
			p.LogicalOperator()
		}
		{
			p.SetState(54)
			p.Condition()
		}


		p.SetState(60)
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
	RightHandExpression() IRightHandExpressionContext
	EQUAL() antlr.TerminalNode
	NOT_EQUAL() antlr.TerminalNode
	LESS_THAN() antlr.TerminalNode
	GREATER_THAN() antlr.TerminalNode

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

func (s *ConditionContext) RightHandExpression() IRightHandExpressionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRightHandExpressionContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRightHandExpressionContext)
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
	p.EnterRule(localctx, 16, cmselectParserRULE_condition)
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
		p.SetState(61)
		p.Attribute()
	}
	{
		p.SetState(62)
		_la = p.GetTokenStream().LA(1)

		if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 240) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(63)
		p.RightHandExpression()
	}



	return localctx
}


// IRightHandExpressionContext is an interface to support dynamic dispatch.
type IRightHandExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	Parameter() IParameterContext
	RhAttribute() IRhAttributeContext

	// IsRightHandExpressionContext differentiates from other interfaces.
	IsRightHandExpressionContext()
}

type RightHandExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRightHandExpressionContext() *RightHandExpressionContext {
	var p = new(RightHandExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cmselectParserRULE_rightHandExpression
	return p
}

func (*RightHandExpressionContext) IsRightHandExpressionContext() {}

func NewRightHandExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RightHandExpressionContext {
	var p = new(RightHandExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cmselectParserRULE_rightHandExpression

	return p
}

func (s *RightHandExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *RightHandExpressionContext) STRING() antlr.TerminalNode {
	return s.GetToken(cmselectParserSTRING, 0)
}

func (s *RightHandExpressionContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(cmselectParserNUMBER, 0)
}

func (s *RightHandExpressionContext) Parameter() IParameterContext {
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

func (s *RightHandExpressionContext) RhAttribute() IRhAttributeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRhAttributeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRhAttributeContext)
}

func (s *RightHandExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RightHandExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RightHandExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cmselectListener); ok {
		listenerT.EnterRightHandExpression(s)
	}
}

func (s *RightHandExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cmselectListener); ok {
		listenerT.ExitRightHandExpression(s)
	}
}




func (p *cmselectParser) RightHandExpression() (localctx IRightHandExpressionContext) {
	this := p
	_ = this

	localctx = NewRightHandExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, cmselectParserRULE_rightHandExpression)

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

	p.SetState(69)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case cmselectParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(65)
			p.Match(cmselectParserSTRING)
		}


	case cmselectParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)
			p.Match(cmselectParserNUMBER)
		}


	case cmselectParserT__2:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(67)
			p.Parameter()
		}


	case cmselectParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(68)
			p.RhAttribute()
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
	p.EnterRule(localctx, 20, cmselectParserRULE_parameter)

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
		p.SetState(71)
		p.Match(cmselectParserT__2)
	}
	{
		p.SetState(72)
		p.Match(cmselectParserIDENTIFIER)
	}
	{
		p.SetState(73)
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
	p.EnterRule(localctx, 22, cmselectParserRULE_logicalOperator)
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
		p.SetState(75)
		_la = p.GetTokenStream().LA(1)

		if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 1792) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


