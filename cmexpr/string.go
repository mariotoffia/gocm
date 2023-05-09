package cmexpr

import (
	"fmt"
	"strings"
)

func (s *SelectExpression) String() string {
	str := "SELECT "

	if s.Projection != nil {
		for _, v := range s.Projection.Expressions {
			str += v.String() + ", "
		}

		str = str[:len(str)-2]
	}

	if s.Where != nil {
		str += " WHERE " + s.Where.String()
	}

	return str
}

func (w *WhereExpression) String() string {
	str := ""

	if w.Condition != nil {
		str += w.Condition.String()
	}

	for _, v := range w.LogicScopes {
		str += " " + v.String()
	}

	return str
}

func (l *LogicalOperatorScope) String() string {
	str := ""

	for op, v := range l.Scopes {
		str += string(op) + " (" + v.String() + ") "
	}

	if len(l.AndExpressions) > 0 {
		if len(l.Scopes) > 0 {
			str += "AND "
		}

		for _, v := range l.AndExpressions {
			str += v.String() + " AND "
		}

		str = str[:len(str)-5]
	}

	if len(l.OrExpressions) > 0 {
		if len(l.Scopes) > 0 && len(l.AndExpressions) == 0 {
			str += "OR "
		}

		for _, v := range l.OrExpressions {
			str += v.String() + " "
		}

		str = str[:len(str)-4]
	}

	if l.NotExpression != nil {
		str += l.NotExpression.String() + " "
	}

	return strings.TrimRight(str, " ")
}

func (a *AndExpression) String() string {
	str := ""

	for _, v := range a.Conditions {
		str += v.String() + " AND "
	}

	return str[:len(str)-5]
}

func (o *OrExpression) String() string {
	str := ""

	for _, v := range o.Conditions {
		str += v.String() + " OR "
	}

	return str[:len(str)-4]
}

func (n *NotExpression) String() string {
	return "NOT " + n.Condition.String()
}

func (c *ConditionExpression) String() string {
	return c.Left.String() + " " + string(c.Condition) + " " + c.Right.String()
}

func (p *PropertyExpression) String() string {
	return p.Property
}

func (r *RHExpression) String() string {
	if r.VariableBinding != "" {
		return ":" + r.VariableBinding + ":"
	}

	if r.Value != nil {
		return fmt.Sprint(r.Value)
	}

	return r.Property.String()
}

func (p *ProjectionExpressions) String() string {
	str := ""

	for _, v := range p.Expressions {
		str += v.String() + ", "
	}

	return str[:len(str)-2]
}

func (p *ProjectionExpression) String() string {
	if p.Alias == "" {
		return p.Property.String()
	}

	return p.Property.String() + " AS " + p.Alias
}
