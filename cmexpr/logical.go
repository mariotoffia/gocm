package cmexpr

// LogicalImpl expresses logical operations such as AND, OR.
type LogicalImpl struct {
	*ExpressionImpl

	and []Expression
	or  []Expression
	not Expression

	child *LogicalImpl
}

// And is takes a left hand side and a right hand side. It will combine
// to evaluate true in all expressions to be true.
func (l *LogicalImpl) And(expr ...Expression) *LogicalImpl {

	l.and = expr

	l.child = &LogicalImpl{ExpressionImpl: &ExpressionImpl{
		root:      l.root,
		condition: l.condition,
	}}

	return l.child

}

// Or is takes a left hand side and a right hand side. All sides needs
// to be true to result in true.
func (l *LogicalImpl) Or(expr ...Expression) *LogicalImpl {

	l.or = expr

	l.child = &LogicalImpl{ExpressionImpl: &ExpressionImpl{
		root:      l.root,
		condition: l.condition,
	}}

	return l.child

}

// Not is a unary expression to negate the right hand side expression.
func (l *LogicalImpl) Not(expr Expression) *LogicalImpl {

	l.not = expr

	l.child = &LogicalImpl{ExpressionImpl: &ExpressionImpl{
		root:      l.root,
		condition: l.condition,
	}}

	return l.child

}
