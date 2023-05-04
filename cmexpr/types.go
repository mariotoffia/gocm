package cmexpr

type SelectExpression struct {
	From       string
	Where      *WhereExpression
	Projection *ProjectionExpressions
}

type WhereExpression struct {
	// Condition when single condition and no logical expressions.
	Condition      *ConditionExpression
	AndExpressions []*AndExpression
	OrExpressions  []*OrExpression
	NotExpression  *NotExpression
}

type AndExpression struct {
	Conditions []*ConditionExpression
}

type OrExpression struct {
	Conditions []*ConditionExpression
}

type NotExpression struct {
	Condition *ConditionExpression
}

type ConditionType string

const (
	ConditionTypeEquals              ConditionType = "=="
	ConditionTypeNotEquals           ConditionType = "!="
	ConditionTypeGreaterThan         ConditionType = ">"
	ConditionTypeGreaterThanOrEquals ConditionType = ">="
	ConditionTypeLessThan            ConditionType = "<"
	ConditionTypeLessThanOrEquals    ConditionType = "<="
)

type ConditionExpression struct {
	Left      *PropertyExpression
	Right     *RHExpression
	Condition ConditionType
}

type PropertyExpressionType string

const (
	PropertyExpressionTypeProperty     PropertyExpressionType = "property"
	PropertyExpressionTypePartitionKey PropertyExpressionType = "pk"
	PropertyExpressionTypeTypeSortKey  PropertyExpressionType = "sk"
)

type PropertyExpression struct {
	Property string
	Type     PropertyExpressionType
}

type RHExpression struct {
	Value           any
	Property        *PropertyExpression
	VariableBinding string
}

type ProjectionExpressions struct {
	Expressions []*ProjectionExpression
}

type ProjectionExpression struct {
	Property *PropertyExpression
	Alias    string
}
