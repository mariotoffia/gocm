package idexpr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyExpression(t *testing.T) {

	err := Parse("", &ParserCallback{
		DivFunc: func(index int) {
			assert.Fail(t, "Should not get a divider")
		},
		CompFunc: func(index int, component string) {
			assert.Fail(t, "Should not get a component")
		},
		IDFunc: func(index int, optional bool, name string) {
			assert.Fail(t, "Should not get any id")
		},
	})

	assert.Equal(t, nil, err)
}

func TestIDNotClosingIDGivesError(t *testing.T) {

	err := Parse("{id", &ParserCallback{
		DivFunc: func(index int) {
			assert.Fail(t, "Should not get a divider")
		},
		CompFunc: func(index int, component string) {
			assert.Fail(t, "Should not get a component")
		},
		IDFunc: func(index int, optional bool, name string) {
			assert.Fail(t, "Should not get any id")
		},
	})

	assert.Error(t, err)
}

func TestIDNotClosing2IDGivesError(t *testing.T) {

	err := Parse("{id#", &ParserCallback{
		DivFunc: func(index int) {
			assert.Fail(t, "Should not get a divider")
		},
		CompFunc: func(index int, component string) {
			assert.Fail(t, "Should not get a component")
		},
		IDFunc: func(index int, optional bool, name string) {
			assert.Fail(t, "Should not get any id")
		},
	})

	assert.Error(t, err)
}

func TestOptionalIDNotClosingIDGivesError(t *testing.T) {

	err := Parse("{?id", &ParserCallback{
		DivFunc: func(index int) {
			assert.Fail(t, "Should not get a divider")
		},
		CompFunc: func(index int, component string) {
			assert.Fail(t, "Should not get a component")
		},
		IDFunc: func(index int, optional bool, name string) {
			assert.Fail(t, "Should not get any id")
		},
	})

	assert.Error(t, err)
}

func TestIDNotOpeningIDGivesError(t *testing.T) {

	err := Parse("id}", &ParserCallback{
		DivFunc: func(index int) {
			assert.Fail(t, "Should not get a divider")
		},
		CompFunc: func(index int, component string) {
			assert.Fail(t, "Should not get a component")
		},
		IDFunc: func(index int, optional bool, name string) {
			assert.Fail(t, "Should not get any id")
		},
	})

	assert.Error(t, err)
}

func TestOptionalIDNotOpeningIDGivesError(t *testing.T) {

	err := Parse("?id}", &ParserCallback{
		DivFunc: func(index int) {
			assert.Fail(t, "Should not get a divider")
		},
		CompFunc: func(index int, component string) {
			assert.Fail(t, "Should not get a component")
		},
		IDFunc: func(index int, optional bool, name string) {
			assert.Fail(t, "Should not get any id")
		},
	})

	assert.Error(t, err)
}

func TestSingleComponent(t *testing.T) {

	components := 0
	err := Parse("BLD", &ParserCallback{
		DivFunc: func(index int) {
			assert.Fail(t, "Should not get a divider")
		},
		CompFunc: func(index int, component string) {
			components++
			assert.Equal(t, 1, components)
			assert.Equal(t, 0, index)
			assert.Equal(t, "BLD", component)
		},
		IDFunc: func(index int, optional bool, name string) {
			assert.Fail(t, "Should not get any id")
		},
	})

	assert.Equal(t, nil, err)
}

func TestSingleID(t *testing.T) {

	ids := 0
	err := Parse("{pelle}", &ParserCallback{
		DivFunc: func(index int) {
			assert.Fail(t, "Should not get a divider")
		},
		CompFunc: func(index int, component string) {
			assert.Fail(t, "Should not get any components")
		},
		IDFunc: func(index int, optional bool, name string) {
			ids++
			assert.Equal(t, 1, ids)
			assert.Equal(t, 0, index)
			assert.Equal(t, false, optional)
			assert.Equal(t, "pelle", name)
		},
	})

	assert.Equal(t, nil, err)
}

func TestSingleOptionalID(t *testing.T) {

	ids := 0
	err := Parse("{?pelle}", &ParserCallback{
		DivFunc: func(index int) {
			assert.Fail(t, "Should not get a divider")
		},
		CompFunc: func(index int, component string) {
			assert.Fail(t, "Should not get any components")
		},
		IDFunc: func(index int, optional bool, name string) {
			ids++
			assert.Equal(t, 1, ids)
			assert.Equal(t, 0, index)
			assert.Equal(t, true, optional)
			assert.Equal(t, "pelle", name)
		},
	})

	assert.Equal(t, nil, err)
}

func TestDividerAtFirstPosWithNothingElseShallFail(t *testing.T) {

	err := Parse("#", &ParserCallback{
		DivFunc: func(index int) {
			assert.Fail(t, "Should not get a divider")
		},
		CompFunc: func(index int, component string) {
			assert.Fail(t, "Should not get any components")
		},
		IDFunc: func(index int, optional bool, name string) {
			assert.Fail(t, "Should not get any id")
		},
	})

	assert.Equal(t, nil, err)
}
func TestTwoComponents(t *testing.T) {

	components := 0
	err := Parse("BLD#TS", &ParserCallback{
		DivFunc: func(index int) {
			assert.Equal(t, 1, index)
		},
		CompFunc: func(index int, component string) {
			components++
			assert.True(t, component == "BLD" || component == "TS")
		},
		IDFunc: func(index int, optional bool, name string) {
			assert.Fail(t, "Should not get any id")
		},
	})

	assert.Equal(t, nil, err)
	assert.Equal(t, 2, components)

}

func TestTwoIDs(t *testing.T) {

	ids := 0
	err := Parse("{id1}#{id2}", &ParserCallback{
		DivFunc: func(index int) {
			assert.Equal(t, 1, index)
		},
		CompFunc: func(index int, component string) {
			assert.Fail(t, "Should not get any component")
		},
		IDFunc: func(index int, optional bool, name string) {
			ids++
			assert.True(t, name == "id1" || name == "id2")
			assert.Equal(t, false, optional)

		},
	})

	assert.Equal(t, nil, err)
	assert.Equal(t, 2, ids)

}

func TestTwoOptionalIDsIsNotAllowed(t *testing.T) {

	ids := 0
	err := Parse("{?id1}#{?id2}", &ParserCallback{
		DivFunc: func(index int) {
			assert.Equal(t, 1, index)
		},
		CompFunc: func(index int, component string) {
			assert.Fail(t, "Should not get any component")
		},
		IDFunc: func(index int, optional bool, name string) {
			ids++
			assert.True(t, name == "id1" || name == "id2")
			assert.Equal(t, true, optional)
		},
	})

	assert.Equal(t, 2, ids)
	assert.Error(t, err)
}

func TestSingleComponentAndID(t *testing.T) {

	invokes := 0
	err := Parse("BLD#{pelle}", &ParserCallback{
		DivFunc: func(index int) {
			assert.Equal(t, 1, index)
			invokes++
		},
		CompFunc: func(index int, component string) {
			assert.Equal(t, 0, index)
			assert.Equal(t, "BLD", component)
			invokes++
		},
		IDFunc: func(index int, optional bool, name string) {
			assert.Equal(t, 2, index)
			assert.Equal(t, "pelle", name)
			assert.Equal(t, false, optional)
			invokes++
		},
	})

	assert.Equal(t, nil, err)
	assert.Equal(t, 3, invokes)

}

func TestSingleComponentAndOptionalID(t *testing.T) {

	invokes := 0
	err := Parse("BLD#{?pelle}", &ParserCallback{
		DivFunc: func(index int) {
			assert.Equal(t, 1, index)
			invokes++
		},
		CompFunc: func(index int, component string) {
			assert.Equal(t, 0, index)
			assert.Equal(t, "BLD", component)
			invokes++
		},
		IDFunc: func(index int, optional bool, name string) {
			assert.Equal(t, 2, index)
			assert.Equal(t, "pelle", name)
			assert.Equal(t, true, optional)
			invokes++
		},
	})

	assert.Equal(t, nil, err)
	assert.Equal(t, 3, invokes)

}

func TestComplexExpression(t *testing.T) {

	result := ""
	err := Parse("BLD#{cd}#TS#{?pg}#{p}#{d}", &ParserCallback{
		DivFunc: func(index int) {
			result += "#"
		},
		CompFunc: func(index int, component string) {
			result += component
		},
		IDFunc: func(index int, optional bool, name string) {

			if optional {
				result += "?:"
			}

			result += name

		},
	})

	assert.Equal(t, "BLD#cd#TS#?:pg#p#d", result)
	assert.Equal(t, nil, err)

}

func BenchmarkComplexExpression(t *testing.B) {

	expr := "BLD#{cd}#TS#{?pg}#{p}#{d}"

	pcb := &ParserCallback{
		DivFunc:  func(index int) {},
		CompFunc: func(index int, component string) {},
		IDFunc:   func(index int, optional bool, name string) {},
	}

	for i := 0; i < t.N; i++ {
		Parse(expr, pcb)
	}
}
