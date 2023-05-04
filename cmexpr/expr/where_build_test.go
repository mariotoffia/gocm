package expr

import (
	"fmt"
	"testing"

	"github.com/mariotoffia/gocm/cmexpr"
)

func TestSimpleExpression(t *testing.T) {

	sb := Expr().Select("table").
		Where().
		PK().Condition(cmexpr.ConditionTypeEquals).Value("123").
		SK().Condition(cmexpr.ConditionTypeEquals).Value("456").
		And(
			Property("age").Condition(cmexpr.ConditionTypeGreaterThan).Value(18),
			Property("building").Condition(cmexpr.ConditionTypeEquals).Value("A"),
			And(
				PK().Condition(cmexpr.ConditionTypeEquals).Value("999"),
				SK().Condition(cmexpr.ConditionTypeEquals).Value("888"),
			),
		).
		Build()

	fmt.Println(sb)
}
