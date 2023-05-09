package expr

import (
	"fmt"

	"github.com/mariotoffia/gocm/cmexpr"
)

func (b *WhereExpressionBuilder) Build() *cmexpr.SelectExpression {

	s := &cmexpr.SelectExpression{
		From: b.parent.(*SelectBuilder).from,
	}

	if len(b.propertyExpressions) == 1 {
		s.Where = &cmexpr.WhereExpression{Condition: b.propertyExpressions[0].toExpression()}

	} else if len(b.propertyExpressions) > 1 {
		s.Where = &cmexpr.WhereExpression{}

		expr := &cmexpr.AndExpression{}
		for _, v := range b.propertyExpressions {
			expr.Conditions = append(expr.Conditions, v.toExpression())
		}

		scope := &cmexpr.LogicalOperatorScope{}
		scope.AndExpressions = append(scope.AndExpressions, expr)
		s.Where.LogicScopes = append(s.Where.LogicScopes, scope)
	}

	// Iterate nested logical expressions and add them nested Scopes expressions
	var log func(
		l *LogicalExprPropertyBuilders,
		parent *cmexpr.LogicalOperatorScope,
		op cmexpr.LogicalOperator,
	)

	log = func(
		l *LogicalExprPropertyBuilders,
		parent *cmexpr.LogicalOperatorScope,
		op cmexpr.LogicalOperator,
	) {
		if len(l.and) == 0 && len(l.or) == 0 && l.not == nil && len(l.conditions) == 0 {
			return
		}

		scp := &cmexpr.LogicalOperatorScope{}

		for _, and := range l.and {
			log(and, scp, cmexpr.LogicalOperatorAnd)
		}

		for _, or := range l.or {
			log(or, scp, cmexpr.LogicalOperatorOr)
		}

		if l.not != nil {
			log(l.not, scp, cmexpr.LogicalOperatorNot)
		}

		if parent != nil {
			if parent.Scopes == nil {
				parent.Scopes = map[cmexpr.LogicalOperator]*cmexpr.LogicalOperatorScope{}
			}

			parent.Scopes[op] = scp
		} else {

			if s.Where == nil {
				s.Where = &cmexpr.WhereExpression{}
			}

			s.Where.LogicScopes = append(s.Where.LogicScopes, scp)
		}

		switch op {
		case cmexpr.LogicalOperatorAnd:
			expr := &cmexpr.AndExpression{}
			for _, v := range l.conditions {
				expr.Conditions = append(expr.Conditions, v.toExpression())
			}

			scp.AndExpressions = append(scp.AndExpressions, expr)
		case cmexpr.LogicalOperatorOr:
			expr := &cmexpr.OrExpression{}
			for _, v := range l.conditions {
				expr.Conditions = append(expr.Conditions, v.toExpression())
			}

			scp.OrExpressions = append(scp.OrExpressions, expr)
		case cmexpr.LogicalOperatorNot:
			if len(l.conditions) > 1 {
				panic(fmt.Sprintf("not expression can only have one condition, got %d", len(l.conditions)))
			}

			scp.NotExpression = &cmexpr.NotExpression{
				Condition: l.not.conditions[0].toExpression(),
			}
		}
	}

	for _, v := range b.and {
		log(v, nil, cmexpr.LogicalOperatorAnd)
	}

	for _, v := range b.or {
		log(v, nil, cmexpr.LogicalOperatorOr)
	}

	if b.not != nil {
		log(b.not, nil, cmexpr.LogicalOperatorNot)
	}

	return s
}
