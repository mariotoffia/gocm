package cmexpr

// PostConditionType specifies the type of the post condition
type PostConditionType string

const (
	// PostConditionPK is a partition key
	PostConditionPK PostConditionType = "pk"
	// PostConditionSK is a secondary key
	PostConditionSK PostConditionType = "sk"
	// PostConditionAttribute is a attribute / property name
	PostConditionAttribute PostConditionType = "att"
	// PostConditionValue is a value, values or a list of values.
	PostConditionValue PostConditionType = "value"
)

// PostConditionImpl is the right hand of the condition.
type PostConditionImpl struct {
	*ExpressionImpl
	name  string
	t     PostConditionType
	value []any
	child *LogicalImpl
}

// Value is one or more values. If multiple it is either a list or
// e.g. a range.
func (pc *PostConditionImpl) Value(values ...any) *LogicalImpl {

	pc.value = values
	pc.t = PostConditionValue

	pc.child = &LogicalImpl{ExpressionImpl: &ExpressionImpl{
		root:      pc.root,
		condition: pc.condition,
	}}

	return pc.child

}

// SK expresses a secondary key
func (pc *PostConditionImpl) SK(name string) *LogicalImpl {

	pc.name = name
	pc.t = PostConditionSK

	pc.child = &LogicalImpl{ExpressionImpl: &ExpressionImpl{
		root:      pc.root,
		condition: pc.condition,
	}}

	return pc.child

}

// PK is expressing a partition key
func (pc *PostConditionImpl) PK(name string) *LogicalImpl {

	pc.name = name
	pc.t = PostConditionPK

	pc.child = &LogicalImpl{ExpressionImpl: &ExpressionImpl{
		root:      pc.root,
		condition: pc.condition,
	}}

	return pc.child

}

// Att expresses an attribute / property name
func (pc *PostConditionImpl) Att(name string) *LogicalImpl {

	pc.name = name
	pc.t = PostConditionAttribute

	pc.child = &LogicalImpl{ExpressionImpl: &ExpressionImpl{
		root:      pc.root,
		condition: pc.condition,
	}}

	return pc.child

}
