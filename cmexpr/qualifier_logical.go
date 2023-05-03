package cmexpr

// QualifierType specifies the type of the qualifier
type QualifierType string

const (
	// QualifierNone is set when the `QualifierLogicalImpl` is
	// a logical expression and not a qualifier.
	QualifierNone QualifierType = ""
	// QualifierPK is a partition key
	QualifierPK QualifierType = "pk"
	// QualifierSK is a secondary key
	QualifierSK QualifierType = "sk"
	// QualifierAttribute is a attribute / property name
	QualifierAttribute QualifierType = "att"
)

// QualifierLogicalImpl is the left hand side of a condition
// or a logical operation.
type QualifierLogicalImpl struct {
	ExpressionImpl
	// name of the attribute to apply the condition to
	name string
	// t is the type of the attribute
	t QualifierType
	// child is the condition operation
	child *ConditionImpl

	and []Expression
	or  []Expression
	not Expression
}

func (l *QualifierLogicalImpl) logical(expr []Expression) *QualifierLogicalImpl {
	qualifier := l.getRoot().condition
	log := &QualifierLogicalImpl{ExpressionImpl: ExpressionImpl{parent: qualifier}}

	qualifier.childs = append(qualifier.childs, log)

	l.childs = make([]Expression, len(expr))

	for i := range expr {
		child := expr[i].(*ExpressionImpl)
		child.parent = l
		l.childs[i] = child
	}

	return log
}

// And is takes a left hand side and a right hand side. It will combine
// to evaluate true in all expressions to be true.
func (l *QualifierLogicalImpl) And(expr ...Expression) *QualifierLogicalImpl {
	l.and = expr
	return l.logical(expr)
}

// Or is takes a left hand side and a right hand side. All sides needs
// to be true to result in true.
func (l *QualifierLogicalImpl) Or(expr ...Expression) *QualifierLogicalImpl {
	l.or = expr
	return l.logical(expr)
}

// Not is a unary expression to negate the right hand side expression.
func (l *QualifierLogicalImpl) Not(expr Expression) *QualifierLogicalImpl {
	l.not = expr
	return l.logical([]Expression{expr})
}

func (q *QualifierLogicalImpl) create(name string, t QualifierType) *ConditionImpl {
	q.name = name
	q.t = t

	q.child = &ConditionImpl{ExpressionImpl: ExpressionImpl{parent: q}}
	q.childs = append(q.childs, q.child)

	return q.child
}

// SK is expressing a Secondary (or Sort) Key
func (q *QualifierLogicalImpl) SK(name string) *ConditionImpl {
	return q.create(name, QualifierSK)
}

// PK is expressing a Partition Key
func (q *QualifierLogicalImpl) PK(name string) *ConditionImpl {
	return q.create(name, QualifierPK)
}

// Att is expressing an attribute / property name
func (q *QualifierLogicalImpl) Att(name string) *ConditionImpl {
	return q.create(name, QualifierAttribute)
}
