package cmexpr

// ExpressionType specifies the type of Expression.
type ExpressionType string

const (
	// Projection specifies what attributes shall be returned, i.e. projected.
	Projection ExpressionType = "projection"
	// KeyCondition specifies the PK and possibly a SK condition
	KeyCondition = "key-condition"
	// Condition specifies a condition to be satisfied before the operation may proceed.
	Condition = "condition"
	// Filter specifies a filter condition that may not be executed
	// in the remote technology.
	//
	// Depending on the implementation capability this may be executed remotely or not.
	Filter = "filter"
	// Update is a update expression
	Update = "update"
)
