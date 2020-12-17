package cmexpr

type logical struct {
	*ExpressionImpl

	and []Expression
	or  []Expression
	not Expression

	child *logical
}

func (l *logical) And(expr ...Expression) *logical {

	l.and = expr

	l.child = &logical{ExpressionImpl: &ExpressionImpl{
		root:      l.root,
		condition: l.condition,
	}}

	return l.child

}

func (l *logical) Or(expr ...Expression) *logical {

	l.or = expr

	l.child = &logical{ExpressionImpl: &ExpressionImpl{
		root:      l.root,
		condition: l.condition,
	}}

	return l.child

}

func (l *logical) Not(expr Expression) *logical {

	l.not = expr

	l.child = &logical{ExpressionImpl: &ExpressionImpl{
		root:      l.root,
		condition: l.condition,
	}}

	return l.child

}
