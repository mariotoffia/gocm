package cmexpr

import (
	"testing"
)

func TestVariousThings(t *testing.T) {

	Expr().Cond().PK("PK").Condition(CondOperEqual).Value("P#SAAB").
		And(
			Expr().Cond().SK("SK").Condition(CondOperGreaterThan).Value(14),
			Expr().Cond().SK("SK").Condition(CondOperLessThan).Value(20),
		).
		Or(
			Expr().Cond().SK("SK").Condition(CondOperEqual).Att("sune"),
		).
		Build()

}
