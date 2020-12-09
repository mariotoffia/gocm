package idexpr

import "fmt"

// ExpressionParserCallback is invoked when parsing the expression.
type ExpressionParserCallback interface {
	// Divider is called when a _divider_ is encountered.
	//
	// The _index_ is the ordinal of the expression, not
	// the string index.
	Divider(index int)
	// Component is called when a component has been encountered.
	Component(index int, component string)
	// ID is invoked when a id has been encountered.
	ID(index int, optional bool, name string)
}

// ParserCallback is a type where functions may be attached to
// do a ad-hoc `ExpressionParserCallback` object.
type ParserCallback struct {
	DivFunc  func(index int)
	CompFunc func(index int, component string)
	IDFunc   func(index int, optional bool, name string)
}

// Divider calls the `DivFunc`
func (pc *ParserCallback) Divider(index int) {
	pc.DivFunc(index)
}

// Component calls the `CompFunc`
func (pc *ParserCallback) Component(index int, component string) {
	pc.CompFunc(index, component)
}

// ID calls the `IDFunc`
func (pc *ParserCallback) ID(index int, optional bool, name string) {
	pc.IDFunc(index, optional, name)
}

// Parse will parse the _expr_ and do callback on the interface.
//
// Format COMPONENT#{id}#{id}#COMPONENT#{id} where _COMPONENT_
// is a constant keyword name to denote a type or hieachy of types.
// The _{id}_ is a identifier that must be substituted with a runtime
// value. When a {?id} is encountered, it is optional and if empty and
// has subsequent divider (#), it will be omitted.
//
// If a empty _divider_ it will default to use '#' as divider. A divider
// may never be more than one rune in length.
//
// NOTE: Only one optional parameter is allowed since it is not possible
// reverse engineer which of the optional parameter is not present if multiple.
func Parse(divider, expr string, cb ExpressionParserCallback) error {

	if divider == "" {
		divider = "#"
	}

	if expr == "" || expr == divider {
		return nil
	}

	divskip := false
	open := -1
	component := -1
	index := 0
	optcnt := 0
	div := rune(divider[0])

	var i int
	var c rune

	for i, c = range expr {

		if c == '?' {

			optcnt++

			if open == -1 {
				return fmt.Errorf("invalid syntax, expression: %s", expr)
			}

			divskip = true
		}

		if c == div {

			if open != -1 {
				return fmt.Errorf("invalid syntax, expression: %s", expr)
			}

			if component != -1 {
				cb.Component(index, expr[component:i])
				index++
				component = -1
			}

			cb.Divider(index)
			index++

			continue

		}

		if c == '{' {

			if open != -1 {
				return fmt.Errorf("invalid syntax, expression: %s", expr)
			}

			open = i
			continue
		}

		if c == '}' {

			if open == -1 {
				return fmt.Errorf("invalid syntax, expression: %s", expr)
			}

			if divskip {
				cb.ID(index, divskip, expr[open+2:i])
			} else {
				cb.ID(index, divskip, expr[open+1:i])
			}

			index++
			divskip = false
			open = -1
			continue
		}

		if open == -1 && component == -1 {
			component = i
		}
	}

	if open != -1 {
		return fmt.Errorf("invalid syntax, expression: %s", expr)
	}

	if component != -1 {
		cb.Component(index, expr[component:i+1])
		index++
		component = -1
	}

	if optcnt > 1 {
		return fmt.Errorf("invalid syntax, expression: %s, only one optional is allowed", expr)
	}

	return nil
}
