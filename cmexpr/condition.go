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

// ConditionImpl is a binary / unary condition.
type ConditionImpl struct {
	ExpressionImpl

	oper  ConditionOperation
	child *PostConditionImpl
}

// Expr is the generic condition function.
func (c *ConditionImpl) Expr(oper ConditionOperation) *PostConditionImpl {
	c.oper = oper
	c.child = &PostConditionImpl{ExpressionImpl: ExpressionImpl{parent: c}}
	c.childs = append(c.childs, c.child)

	return c.child
}
