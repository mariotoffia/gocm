package cmexpr

// QualifierType specifies the type of the qualifier
type QualifierType string

const (
	// QualifierPK is a partition key
	QualifierPK QualifierType = "pk"
	// QualifierSK is a secondary key
	QualifierSK QualifierType = "sk"
	// QualifierAttribute is a attribute / property name
	QualifierAttribute QualifierType = "att"
)

// QualifierImpl is the left hand side of a condition.
type QualifierImpl struct {
	ExpressionImpl

	name  string
	t     QualifierType
	child *ConditionImpl
}

// SK is expressing a Secondary (or Sort) Key
func (q *QualifierImpl) SK(name string) *ConditionImpl {

	q.name = name
	q.t = QualifierSK

	q.child = &ConditionImpl{ExpressionImpl: ExpressionImpl{
		root:      q.root,
		condition: q.condition,
	}}

	return q.child
}

// PK is expressing a Partition Key
func (q *QualifierImpl) PK(name string) *ConditionImpl {

	q.name = name
	q.t = QualifierPK

	q.child = &ConditionImpl{ExpressionImpl: ExpressionImpl{
		root:      q.root,
		condition: q.condition,
	}}

	return q.child
}

// Att is expressing an attribute / property name
func (q *QualifierImpl) Att(name string) *ConditionImpl {

	q.name = name
	q.t = QualifierAttribute

	q.child = &ConditionImpl{ExpressionImpl: ExpressionImpl{
		root:      q.root,
		condition: q.condition,
	}}

	return q.child
}
