package cmexpr

import "context"

// ContextTag is used to be embedded into value context.
type ContextTag struct {
	key, val interface{}
}

// ConditionExpressionVisitor shall be implemented by ...
//
// Context handling is done through a base context of which gets extended
// with sub-context (value context) in parts of the system. It also means
// if e.g. nested _And_, _Or_, _Not_ it may have more than one sub-context since it
// is pushed as a stack. For each end the context is popped until the last
// `EndExpression()` where the initial context (in `BeginExpression`) is passed.
type ConditionExpressionVisitor interface {
	// BeginExpression is called when the parsing begins.
	//
	// It will receives a initialized context.
	BeginExpression(ctx context.Context)
	// BeginExpression is called when the parsing begins.
	//
	// It will receives the same context as it got when `BeginExpression` was called.
	EndExpression(ctx context.Context)
	// BeginCondition is invoked when starting a (local) condition expression
	//
	// The returned _ContextTag_ is used to create a new context using `context.WithValue(ctx, key, value)`.
	// The sub-context is then passed on all methods including `EndCondition()`. After that function the
	// sub-context is discarded and the parent is used in subsequent functions.
	BeginCondition(ctx context.Context) *ContextTag
	EndCondition(ctx context.Context)
	PartitionKey(ctx context.Context, name string)
	SecondaryKey(ctx context.Context, name string)
	Attribute(ctx context.Context, name string)
	Condition(ctx context.Context, oper ConditionOperation)
	Value(ctx context.Context, value ...interface{})
	// Begin And is invoked when a `And` operation is encountered.
	//
	// The returned _ContextTag_ is used to create a new context using `context.WithValue(ctx, key, value)`.
	// The sub-context is then passed to all methods including `EndAnd()`. After `EndAnd()` the sub-context
	// is destroyed and the parent is used.
	BeginAnd(ctx context.Context) *ContextTag
	EndAnd(ctx context.Context)
	// Begin Or is invoked when a `Or` operation is encountered.
	//
	// The returned _ContextTag_ is used to create a new context using `context.WithValue(ctx, key, value)`.
	// The sub-context is then passed to all methods including `EndOr()`. After `EndOr()` the sub-context
	// is destroyed and the parent is used.
	BeginOr(ctx context.Context) *ContextTag
	EndOr(ctx context.Context)
	// Begin Not is invoked when a `Not` operation is encountered.
	//
	// The returned _ContextTag_ is used to create a new context using `context.WithValue(ctx, key, value)`.
	// The sub-context is then passed to all methods including `EndNot()`. After `EndNot()` the sub-context
	// is destroyed and the parent is used.
	BeginNot(ctx context.Context) *ContextTag
	EndNot(ctx context.Context)
}
