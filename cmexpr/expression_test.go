package cmexpr

import (
	"testing"
)

func TestConditionLogical(t *testing.T) {

	Expr().Cond().PK("PK").Condition(CondOperEqual).Value("P#SAAB").
		And(
			Expr().Cond().SK("SK").Condition(CondOperGreaterThan).Value(14),
			Expr().Cond().SK("brand").Condition(CondOperLessThan).Value(20),
		).
		Or(
			Expr().Cond().SK("SK").Condition(CondOperEqual).Att("sune"),
		).
		Build()
}

func TestProjectionWithOneCondition(t *testing.T) {
	Expr().Project().Attribute("building", "PK").Attribute("room", "SK").
		Cond().SK("room").Condition(CondOperEqual).Value("R#Lgh141").
		Cond().PK("building").Condition(CondOperEqual).Value("B#Måssmyren 18").
		Cond().Att("tenant").Condition(CondOperEqual).Bind("tenant").
		Build()

}
