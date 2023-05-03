package cmexpr

// LogicalImpl is the left hand side of a condition
// or a logical operation.
type LogicalImpl struct {
	ExpressionImpl

	and []Expression
	or  []Expression
	not Expression
}

func (l *LogicalImpl) logical(expr []Expression) *QualifierLogicalImpl {
	qualifier := l.getRoot().condition
	log := &QualifierLogicalImpl{ExpressionImpl: ExpressionImpl{parent: qualifier}}

	qualifier.childs = append(qualifier.childs, log)

	l.childs = make([]Expression, len(expr))
	for i := range expr {
		child := expr[i].(*ExpressionImpl)
		child.parent = l
		l.childs[i] = child
	}

	return log
}

// And is takes a left hand side and a right hand side. It will combine
// to evaluate true in all expressions to be true.
func (l *LogicalImpl) And(expr ...Expression) *QualifierLogicalImpl {
	l.and = expr
	return l.logical(expr)
}

// Or is takes a left hand side and a right hand side. All sides needs
// to be true to result in true.
func (l *LogicalImpl) Or(expr ...Expression) *QualifierLogicalImpl {
	l.or = expr
	return l.logical(expr)
}

// Not is a unary expression to negate the right hand side expression.
func (l *LogicalImpl) Not(expr Expression) *QualifierLogicalImpl {
	l.not = expr
	return l.logical([]Expression{expr})
}
