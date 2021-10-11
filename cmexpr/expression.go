package cmexpr

import (
	"encoding/json"
	"fmt"
)

// ExpressionType specifies the type of Expression.
type ExpressionType string

const (
	// Projection specifies what attributes shall be returned, i.e. projected.
	Projection ExpressionType = "projection"
	// KeyCondition specifies the PK and possibly a SK condition
	KeyCondition = "key-condition"
	// Condition specifies a condition to be satisfied before the operation may proceed.
	Condition = "condition"
	// Filter specifies a filter condition that may not be executed
	// in the remote technology.
	//
	// Depending on the implementation capability this may be executed remotely or not.
	Filter = "filter"
	// Update is a update expression
	Update = "update"
)

// Expression is the base interface.
type Expression interface {
	Cond() *QualifierImpl
}

// ExpressionImpl is the base struct
type ExpressionImpl struct {
	root      *ExpressionImpl
	condition *QualifierImpl
}

// Expr creates a new expression
func Expr() *ExpressionImpl {
	return &ExpressionImpl{}
}

// Cond creates a condition expression.
func (e *ExpressionImpl) Cond() *QualifierImpl {

	q := &QualifierImpl{ExpressionImpl: ExpressionImpl{root: e}}
	q.condition = q

	return q
}

// Build will build the expression
func (e *ExpressionImpl) Build() {
	// We got build. lets enumerate the trees.

	if e.condition != nil {

		ProcessQualifier(e.condition)

		// Build until x.child.expression == e (i.e. skip this e since it is a "dead" child
	}
}

// ProcessQualifier will process a single qualifier and any
// and, or, not expressions.
func ProcessQualifier(qualifier *QualifierImpl) {

	cond := qualifier.child
	postCond := qualifier.child.child

	expr := &ConditionExpr{
		LeftType:   qualifier.t,
		LeftName:   qualifier.name,
		Condition:  cond.oper,
		RightName:  postCond.name,
		RightType:  postCond.t,
		RightValue: postCond.value,
	}

	data, _ := json.Marshal(expr)
	fmt.Println(string(data))

	ProcessLogicalOper(postCond.child)
}

// ProcessLogicalOper will process and, or and not
func ProcessLogicalOper(logic *LogicalImpl) {

}
