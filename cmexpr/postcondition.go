package cmexpr

type postcondition struct {
	*ExpressionImpl

	sk    string
	pk    string
	att   string
	value []interface{}
	child *logical
}

func (pc *postcondition) Value(values ...interface{}) *logical {

	pc.value = values

	pc.child = &logical{ExpressionImpl: &ExpressionImpl{
		root:      pc.root,
		condition: pc.condition,
	}}

	return pc.child

}

func (pc *postcondition) SK(name string) *logical {

	pc.sk = name

	pc.child = &logical{ExpressionImpl: &ExpressionImpl{
		root:      pc.root,
		condition: pc.condition,
	}}

	return pc.child

}

func (pc *postcondition) PK(name string) *logical {

	pc.pk = name

	pc.child = &logical{ExpressionImpl: &ExpressionImpl{
		root:      pc.root,
		condition: pc.condition,
	}}

	return pc.child

}

func (pc *postcondition) Att(name string) *logical {

	pc.att = name

	pc.child = &logical{ExpressionImpl: &ExpressionImpl{
		root:      pc.root,
		condition: pc.condition,
	}}

	return pc.child

}
