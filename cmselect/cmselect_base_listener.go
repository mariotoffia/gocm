// Code generated from cmselect.g4 by ANTLR 4.12.0. DO NOT EDIT.

package cmselect // cmselect
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BasecmselectListener is a complete listener for a parse tree produced by cmselectParser.
type BasecmselectListener struct{}

var _ cmselectListener = &BasecmselectListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasecmselectListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasecmselectListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasecmselectListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasecmselectListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQuery is called when production query is entered.
func (s *BasecmselectListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BasecmselectListener) ExitQuery(ctx *QueryContext) {}

// EnterSelectClause is called when production selectClause is entered.
func (s *BasecmselectListener) EnterSelectClause(ctx *SelectClauseContext) {}

// ExitSelectClause is called when production selectClause is exited.
func (s *BasecmselectListener) ExitSelectClause(ctx *SelectClauseContext) {}

// EnterAttributeList is called when production attributeList is entered.
func (s *BasecmselectListener) EnterAttributeList(ctx *AttributeListContext) {}

// ExitAttributeList is called when production attributeList is exited.
func (s *BasecmselectListener) ExitAttributeList(ctx *AttributeListContext) {}

// EnterAttribute is called when production attribute is entered.
func (s *BasecmselectListener) EnterAttribute(ctx *AttributeContext) {}

// ExitAttribute is called when production attribute is exited.
func (s *BasecmselectListener) ExitAttribute(ctx *AttributeContext) {}

// EnterRhAttribute is called when production rhAttribute is entered.
func (s *BasecmselectListener) EnterRhAttribute(ctx *RhAttributeContext) {}

// ExitRhAttribute is called when production rhAttribute is exited.
func (s *BasecmselectListener) ExitRhAttribute(ctx *RhAttributeContext) {}

// EnterFromClause is called when production fromClause is entered.
func (s *BasecmselectListener) EnterFromClause(ctx *FromClauseContext) {}

// ExitFromClause is called when production fromClause is exited.
func (s *BasecmselectListener) ExitFromClause(ctx *FromClauseContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BasecmselectListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BasecmselectListener) ExitTableName(ctx *TableNameContext) {}

// EnterWhereClause is called when production whereClause is entered.
func (s *BasecmselectListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BasecmselectListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterCondition is called when production condition is entered.
func (s *BasecmselectListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BasecmselectListener) ExitCondition(ctx *ConditionContext) {}

// EnterRightHandExpression is called when production rightHandExpression is entered.
func (s *BasecmselectListener) EnterRightHandExpression(ctx *RightHandExpressionContext) {}

// ExitRightHandExpression is called when production rightHandExpression is exited.
func (s *BasecmselectListener) ExitRightHandExpression(ctx *RightHandExpressionContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BasecmselectListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BasecmselectListener) ExitParameter(ctx *ParameterContext) {}

// EnterLogicalOperator is called when production logicalOperator is entered.
func (s *BasecmselectListener) EnterLogicalOperator(ctx *LogicalOperatorContext) {}

// ExitLogicalOperator is called when production logicalOperator is exited.
func (s *BasecmselectListener) ExitLogicalOperator(ctx *LogicalOperatorContext) {}
