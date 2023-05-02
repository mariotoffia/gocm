package cmexpr

import (
	"encoding/json"
	"fmt"
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

func (e *ExpressionImpl) Project() *ProjectionImpl {
	return &ProjectionImpl{ExpressionImpl: ExpressionImpl{root: e}}
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
