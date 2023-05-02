package cmexpr

// PostConditionType specifies the type of the post condition
type PostConditionType string

const (
	// PostConditionAttribute is a attribute / property name
	PostConditionAttribute PostConditionType = "att"
	// PostConditionValue is a value, values or a list of values.
	PostConditionValue PostConditionType = "value"
	// PostConditionVariableBinding specifies one or more variable binding
	PostConditionVariableBinding PostConditionType = "variable-binding"
)

// PostConditionImpl is the right hand of the condition.
type PostConditionImpl struct {
	*ExpressionImpl
	name      string
	t         PostConditionType
	value     []any
	variables []string
	child     *LogicalImpl
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

// Bind expresses one or more bindings. This is when the
// expression expect a variable to be bound to the expression. This
// enacted when the expression is evaluated.
//
// Hence, it is possible to `Build` an expression that can be reused
// with different variable content.
func (pc *PostConditionImpl) Bind(variable ...string) *LogicalImpl {

	pc.variables = variable
	pc.t = PostConditionValue

	pc.child = &LogicalImpl{ExpressionImpl: &ExpressionImpl{
		root:      pc.root,
		condition: pc.condition,
	}}

	return pc.child

}

// Att expresses an attribute / property name in the data
func (pc *PostConditionImpl) Att(name string) *LogicalImpl {

	pc.name = name
	pc.t = PostConditionAttribute

	pc.child = &LogicalImpl{ExpressionImpl: &ExpressionImpl{
		root:      pc.root,
		condition: pc.condition,
	}}

	return pc.child

}
