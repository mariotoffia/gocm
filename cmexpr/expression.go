package cmexpr

import (
	"encoding/json"
	"fmt"
)

// Expression is the base interface.
type Expression interface {
	getChilds() []Expression
	getParent() Expression
	getRoot() *RootExpressionImpl
}

func (expr *ExpressionImpl) getChilds() []Expression {
	return expr.childs
}

func (expr *ExpressionImpl) getParent() Expression {
	return expr.parent
}

func (expr *ExpressionImpl) getRoot() *RootExpressionImpl {
	var e Expression = expr

	for {
		if e.getParent() == nil {
			break
		}
		e = e.getParent()
	}

	return e.(*RootExpressionImpl)
}

// ExpressionImpl is the base struct
type ExpressionImpl struct {
	parent Expression
	childs []Expression
}

type RootExpressionImpl struct {
	ExpressionImpl
	condition  *QualifierLogicalImpl
	projection *ProjectionImpl
}

// Expr creates a new expression
func Expr() *RootExpressionImpl {
	return &RootExpressionImpl{}
}

// Where creates a condition expression.
func (e *RootExpressionImpl) Where() *QualifierLogicalImpl {

	if e.condition != nil {
		panic("condition expression already created")
	}

	e.condition = &QualifierLogicalImpl{ExpressionImpl: ExpressionImpl{parent: e}}

	e.childs = append(e.childs, e.condition)
	return e.condition
}

func (e *RootExpressionImpl) Project() *ProjectionImpl {
	if e.projection != nil {
		panic("projection expression already created")
	}

	e.projection = &ProjectionImpl{
		ExpressionImpl: ExpressionImpl{parent: e},
	}

	e.childs = append(e.childs, e.projection)
	return e.projection
}

// Build will build the expression
func (e *RootExpressionImpl) Build() {
	// We got build. lets enumerate the trees.

	if e.condition != nil {

		ProcessQualifier(e.condition)

		// Build until x.child.expression == e (i.e. skip this e since it is a "dead" child
	}
}

// ProcessQualifier will process a single qualifier and any
// and, or, not expressions.
func ProcessQualifier(qualifier *QualifierLogicalImpl) {

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
}
