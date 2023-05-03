package cmexpr

import (
	"testing"
)

func TestConditionLogical(t *testing.T) {

	Expr().Where().
		PK("PK").Expr(CondOperEqual).Value("P#SAAB").
		And().
		SK("SK").Expr(CondOperEqual).Value("S#9000").
		And(
			Expr().Where().SK("mickey").Expr(CondOperGreaterThan).Value(14),
			Expr().Where().SK("mouse").Expr(CondOperLessThan).Value(20),
		).
		Or(
			Expr().Where().SK("SK").Expr(CondOperEqual).Att("sune"),
		)
}

/*
func TestProjectionWithOneCondition(t *testing.T) {
	Expr().Project().Attribute("b", "building").Attribute("r", "room").
		Cond().SK("SK").Condition(CondOperEqual).Value("R#Lgh141").
		Cond().PK("PK").Condition(CondOperEqual).Value("B#Måssmyren 18").
		Cond().Att("tenant").Condition(CondOperEqual).Bind("tenant").
		Build()

}*/
