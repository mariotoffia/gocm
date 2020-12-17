package cmexpr

type QualifierImpl struct {
	ExpressionImpl

	sk    string
	pk    string
	att   string
	child *condition
}

func (q *QualifierImpl) SK(name string) *condition {

	q.sk = name

	q.child = &condition{ExpressionImpl: ExpressionImpl{
		root:      q.root,
		condition: q.condition,
	}}

	return q.child
}

func (q *QualifierImpl) PK(name string) *condition {

	q.pk = name

	q.child = &condition{ExpressionImpl: ExpressionImpl{
		root:      q.root,
		condition: q.condition,
	}}

	return q.child
}

func (q *QualifierImpl) Att(name string) *condition {

	q.att = name

	q.child = &condition{ExpressionImpl: ExpressionImpl{
		root:      q.root,
		condition: q.condition,
	}}

	return q.child
}
