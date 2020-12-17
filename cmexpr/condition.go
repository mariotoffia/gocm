package cmexpr

// ConditionOperation is where a left and right hand is checked using a logical operation.
type ConditionOperation string

const (
	// CondOperEqual is equal.
	CondOperEqual ConditionOperation = "EQ"
	// CondOperNotEqual is not equal.
	CondOperNotEqual ConditionOperation = "NEQ"
	// CondOperGreaterThan is greater than.
	CondOperGreaterThan ConditionOperation = "GT"
	// CondOperGreaterThanOrEqual is greater than or equal.
	CondOperGreaterThanOrEqual ConditionOperation = "GTE"
	// CondOperLessThan is less than.
	CondOperLessThan ConditionOperation = "LT"
	// CondOperLessThanOrEqual is less or equal.
	CondOperLessThanOrEqual ConditionOperation = "LTE"
	// CondOperIn matches the left side with a set of right hand values.
	CondOperIn ConditionOperation = "IN"
	// CondOperBetween matches the left side that is is between right hand side value.
	CondOperBetween ConditionOperation = "BT"
	// CondOperBeginsWith checks that left side begins with a certain string value.
	CondOperBeginsWith ConditionOperation = "BW"
	// CondOperEndsWith checks that left side begins with a certain string value.
	CondOperEndsWith ConditionOperation = "EW"
)

type condition struct {
	ExpressionImpl

	oper  ConditionOperation
	child *postcondition
}

// Condition is the generic condition function.
func (c *condition) Condition(oper ConditionOperation) *postcondition {

	c.oper = oper

	c.child = &postcondition{ExpressionImpl: &ExpressionImpl{
		root:      c.root,
		condition: c.condition,
	}}

	return c.child
}
