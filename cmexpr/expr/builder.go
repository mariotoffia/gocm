package expr

import (
	"fmt"

	"github.com/mariotoffia/gocm/cmexpr"
)

type builderBase interface {
	getParent() builderBase
	getChildren() []builderBase
}
type builderBaseImpl struct {
	parent   builderBase
	children []builderBase
}

func (b *builderBaseImpl) getParent() builderBase {
	return b.parent
}

func (b *builderBaseImpl) getChildren() []builderBase {
	return b.children
}

type ExprBuilder struct {
	selectBuilder *SelectBuilder
}

func Expr() *ExprBuilder {
	return &ExprBuilder{}
}

func (b *ExprBuilder) Select(from string) *SelectBuilder {
	b.selectBuilder = &SelectBuilder{
		from: from,
	}

	b.selectBuilder.children = append(b.selectBuilder.children, b.selectBuilder)

	return b.selectBuilder
}

type SelectBuilder struct {
	builderBaseImpl
	from  string
	where *WhereExpressionBuilder
}

func (b *SelectBuilder) Where() *WhereExpressionBuilder {
	b.where = &WhereExpressionBuilder{
		builderBaseImpl: builderBaseImpl{parent: b},
	}

	b.children = append(b.children, b.where)

	return b.where
}

type LogicalExprPropertyBuilders struct {
	builderBaseImpl
	conditions []*ConditionBuilder
	and        []*LogicalExprPropertyBuilders
	or         []*LogicalExprPropertyBuilders
	not        *LogicalExprPropertyBuilders
}

type WhereExpressionBuilder struct {
	builderBaseImpl
	// propertyExpressions are automatically ANDed together.
	propertyExpressions []*ConditionBuilder
	and                 []*LogicalExprPropertyBuilders
	or                  []*LogicalExprPropertyBuilders
	not                 *LogicalExprPropertyBuilders
}

type PropertyExpressionType string

const (
	AttributeTypePK       PropertyExpressionType = "PK"
	AttributeTypeSK       PropertyExpressionType = "SK"
	AttributeTypeProperty PropertyExpressionType = "Property"
)

func (b *WhereExpressionBuilder) property(name string, t cmexpr.PropertyExpressionType) *ConditionBuilder {
	attr := &ConditionBuilder{
		builderBaseImpl: builderBaseImpl{parent: b},
		name:            name,
		attrType:        t,
	}

	b.children = append(b.children, attr)
	b.propertyExpressions = append(b.propertyExpressions, attr)

	return attr
}

func (b *WhereExpressionBuilder) Property(name string) *ConditionBuilder {
	return b.property(name, cmexpr.PropertyExpressionTypeProperty)
}

func (b *WhereExpressionBuilder) PK(name ...string) *ConditionBuilder {
	if len(name) == 0 {
		name = []string{"PK"}
	}

	return b.property(name[0], cmexpr.PropertyExpressionTypePartitionKey)
}

func (b *WhereExpressionBuilder) SK(name ...string) *ConditionBuilder {
	if len(name) == 0 {
		name = []string{"SK"}
	}
	return b.property(name[0], cmexpr.PropertyExpressionTypeTypeSortKey)
}

func And(and ...*WhereExpressionBuilder) *WhereExpressionBuilder {
	w := &WhereExpressionBuilder{}
	return w.logical("AND", and)
}

func (b *WhereExpressionBuilder) logical(oper string, log []*WhereExpressionBuilder) *WhereExpressionBuilder {
	if len(log) == 0 {
		panic(fmt.Sprintf("not supported: empty %s", oper))
	}

	l := &LogicalExprPropertyBuilders{
		builderBaseImpl: builderBaseImpl{
			parent: b,
		},
	}

	for i := range log {
		log := log[i]
		if log.parent != nil {
			panic("expecting a detached builder")
		}

		for _, pe := range log.propertyExpressions {
			pe.parent = l
			l.conditions = append(l.conditions, pe)
			l.children = append(l.children, pe)
		}

		log.propertyExpressions = nil

		if len(log.and) > 0 {
			for _, a := range log.and {
				a.parent = l
				l.and = append(l.and, a)
				l.children = append(l.children, a)
			}
		}

		if len(log.or) > 0 {
			for _, o := range log.or {
				o.parent = l
				l.or = append(l.or, o)
				l.children = append(l.children, o)
			}
		}

		if log.not != nil {
			log.not.parent = l
			l.not = log.not
			l.children = append(l.children, log.not)
		}
	}

	switch oper {
	case "AND":
		b.and = append(b.and, l)
	case "OR":
		b.or = append(b.or, l)
	case "NOT":
		if b.not != nil {
			panic("not supported: multiple NOT")
		}
		b.not = l
	}

	return b
}

func (b *WhereExpressionBuilder) And(and ...*WhereExpressionBuilder) *WhereExpressionBuilder {
	return b.logical("AND", and)
}

func Property(name string) *ConditionBuilder {
	w := &WhereExpressionBuilder{}
	return w.property(name, cmexpr.PropertyExpressionTypeProperty)
}

func PK(name ...string) *ConditionBuilder {
	w := &WhereExpressionBuilder{}
	return w.PK(name...)
}

func SK(name ...string) *ConditionBuilder {
	w := &WhereExpressionBuilder{}
	return w.SK(name...)
}

type ConditionBuilder struct {
	builderBaseImpl
	name      string
	attrType  cmexpr.PropertyExpressionType
	condition *RHCBuilder
}

func (b *ConditionBuilder) toExpression() *cmexpr.ConditionExpression {
	cond := &cmexpr.ConditionExpression{
		Left: &cmexpr.PropertyExpression{
			Type:     b.attrType,
			Property: b.name,
		},
		Condition: b.condition.condition,
		Right: &cmexpr.RHExpression{
			Value:           b.condition.value,
			VariableBinding: b.condition.varbind,
		},
	}

	if b.condition.property != "" {
		cond.Right.Property = &cmexpr.PropertyExpression{
			Type:     b.condition.propertyType,
			Property: b.condition.property,
		}
	}

	return cond
}

func (b *ConditionBuilder) Condition(condition cmexpr.ConditionType) *RHCBuilder {
	b.condition = &RHCBuilder{
		builderBaseImpl: builderBaseImpl{parent: b},
		condition:       condition,
	}

	b.children = append(b.children, b.condition)

	return b.condition
}

type RHCBuilder struct {
	builderBaseImpl
	condition    cmexpr.ConditionType
	value        any
	varbind      string
	property     string
	propertyType cmexpr.PropertyExpressionType
}

func (b *RHCBuilder) Value(value any) *WhereExpressionBuilder {
	b.value = value
	return parent[WhereExpressionBuilder](b.parent)
}

func (b *RHCBuilder) VarBind(variable string) *WhereExpressionBuilder {
	b.varbind = variable
	return parent[WhereExpressionBuilder](b.parent)
}

func (b *RHCBuilder) prop(property string, pt cmexpr.PropertyExpressionType) *WhereExpressionBuilder {
	b.property = property
	b.propertyType = pt
	return parent[WhereExpressionBuilder](b.parent)
}

func (b *RHCBuilder) Property(property string) *WhereExpressionBuilder {
	return b.prop(property, cmexpr.PropertyExpressionTypeProperty)
}

func (b *RHCBuilder) PK(property ...string) *WhereExpressionBuilder {
	if len(property) == 0 {
		property = []string{"PK"}
	}

	return b.prop(property[0], cmexpr.PropertyExpressionTypeProperty)
}

func (b *RHCBuilder) SK(property ...string) *WhereExpressionBuilder {
	if len(property) == 0 {
		property = []string{"SK"}
	}

	return b.prop(property[0], cmexpr.PropertyExpressionTypeProperty)
}

func parent[T any](builder builderBase) *T {
	current := builder

	for {
		if child, ok := any(current).(*T); ok {
			return child
		}

		if current.getParent() == nil {
			break
		}

		current = current.getParent()
	}

	panic("builder not found")
}
