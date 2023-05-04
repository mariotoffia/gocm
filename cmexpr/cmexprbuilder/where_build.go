package cmexprbuilder

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

	if len(b.and) > 0 {
		if s.Where == nil {
			s.Where = &cmexpr.WhereExpression{}
		}

		if len(s.Where.LogicScopes) == 0 {
			s.Where.LogicScopes = []*cmexpr.LogicalOperatorScope{{}}
		}

		scope := s.Where.LogicScopes[len(s.Where.LogicScopes)-1]

		for _, v := range b.and {
			expr := &cmexpr.AndExpression{}
			for _, c := range v.conditions {
				expr.Conditions = append(expr.Conditions, c.toExpression())
			}

			scope.AndExpressions = append(scope.AndExpressions, expr)
		}
	}

	if len(b.or) > 0 {
		if s.Where == nil {
			s.Where = &cmexpr.WhereExpression{}
		}

		if len(s.Where.LogicScopes) == 0 {
			s.Where.LogicScopes = []*cmexpr.LogicalOperatorScope{{}}
		}

		scope := s.Where.LogicScopes[len(s.Where.LogicScopes)-1]

		for _, v := range b.or {
			expr := &cmexpr.OrExpression{}
			for _, c := range v.conditions {
				expr.Conditions = append(expr.Conditions, c.toExpression())
			}

			scope.OrExpressions = append(scope.OrExpressions, expr)
		}
	}

	if b.not != nil {
		if len(b.not.conditions) > 1 {
			panic(fmt.Sprintf("not expression can only have one condition, got %d", len(b.not.conditions)))
		}
		if s.Where == nil {
			s.Where = &cmexpr.WhereExpression{}
		}

		if len(s.Where.LogicScopes) == 0 {
			s.Where.LogicScopes = []*cmexpr.LogicalOperatorScope{{}}
		}

		scope := s.Where.LogicScopes[len(s.Where.LogicScopes)-1]
		scope.NotExpression = &cmexpr.NotExpression{
			Condition: b.not.conditions[0].toExpression(),
		}
	}

	return s
}
