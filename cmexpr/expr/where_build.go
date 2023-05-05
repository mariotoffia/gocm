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
	var log func(l *LogicalExprPropertyBuilders, parent *cmexpr.LogicalOperatorScope, op string)

	log = func(l *LogicalExprPropertyBuilders, parent *cmexpr.LogicalOperatorScope, op string) {
		if len(l.and) == 0 && len(l.or) == 0 && l.not == nil {
			return
		}

		scp := &cmexpr.LogicalOperatorScope{}

		if parent != nil {
			parent.Scopes = append(parent.Scopes, scp)
		} else {
			if s.Where == nil {
				s.Where = &cmexpr.WhereExpression{}
			}

			s.Where.LogicScopes = append(s.Where.LogicScopes, scp)
		}

		switch op {
		case "AND":
			expr := &cmexpr.AndExpression{}
			for _, v := range l.conditions {
				expr.Conditions = append(expr.Conditions, v.toExpression())
			}

			scp.AndExpressions = append(scp.AndExpressions, expr)
		case "OR":
			expr := &cmexpr.OrExpression{}
			for _, v := range l.conditions {
				expr.Conditions = append(expr.Conditions, v.toExpression())
			}

			scp.OrExpressions = append(scp.OrExpressions, expr)
		case "NOT":
			if len(l.conditions) > 1 {
				panic(fmt.Sprintf("not expression can only have one condition, got %d", len(l.conditions)))
			}

			scp.NotExpression = &cmexpr.NotExpression{
				Condition: l.not.conditions[0].toExpression(),
			}
		}

		for _, and := range l.and {
			expr := &cmexpr.AndExpression{}
			for _, c := range and.conditions {
				expr.Conditions = append(expr.Conditions, c.toExpression())
			}

			scp.AndExpressions = append(scp.AndExpressions, expr)

			log(and, scp, "AND")
		}

		for _, or := range l.or {
			expr := &cmexpr.OrExpression{}
			for _, c := range or.conditions {
				expr.Conditions = append(expr.Conditions, c.toExpression())
			}

			scp.OrExpressions = append(scp.OrExpressions, expr)

			log(or, scp, "OR")
		}

		if l.not != nil {
			if len(l.not.conditions) > 1 {
				panic(fmt.Sprintf("not expression can only have one condition, got %d", len(l.not.conditions)))
			}

			scp.NotExpression = &cmexpr.NotExpression{
				Condition: l.not.conditions[0].toExpression(),
			}

			log(l.not, scp, "NOT")
		}
	}

	for _, v := range b.and {
		log(v, nil, "AND")
	}

	for _, v := range b.or {
		log(v, nil, "OR")
	}

	if b.not != nil {
		log(b.not, nil, "NOT")
	}

	return s
}
