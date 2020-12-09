package idreflect

import (
	"reflect"

	"github.com/mariotoffia/gocm/identity"
	"github.com/mariotoffia/ssm/parser"
)

// MapperParser parser that uses _cm_ tags.
type MapperParser struct {
	tp      *parser.Parser
	mappers map[reflect.Type]*MapperImpl
	divider string
	err     error
}

// NewParser creates a new `MapperParser`
func NewParser() *MapperParser {

	return &MapperParser{
		tp: parser.New("", "", "").
			RegisterTagParser("cm", parser.NewTagParser([]string{"name", "pk", "sk"})),
		mappers: map[reflect.Type]*MapperImpl{},
		divider: identity.IDStandardDivider,
	}

}

// UseDivider changes the standard divider `identity.IDStandarDivider` to _divider_.
//
// The already `Add()`ed `IDObjectMapper`s are using the previous and all new will
// use the newly set _divider_. Hence, it is possible to have different dividers in
// same `MapperParser`.
func (p *MapperParser) UseDivider(divider string) *MapperParser {
	p.divider = divider
	return p
}

// Error returns any error state
func (p *MapperParser) Error() error {
	return p.err
}

// ClearError will clear any error state
func (p *MapperParser) ClearError() *MapperParser {
	p.err = nil
	return p
}

// Mapper gets the mapper for the parameter _v_ (as it was registered using `Add()`).
func (p *MapperParser) Mapper(v interface{}) *MapperImpl {

	if m, ok := p.mappers[reflect.TypeOf(v)]; ok {
		return m
	}

	return nil
}

// Add creates a new reflection based `IdMapper` and stores it in its internal hash.
//
// The parameter _v_ is expected to be a pointer to a type.
func (p *MapperParser) Add(v interface{}) *MapperParser {

	reflect.TypeOf(v)
	sn, err := p.tp.Parse(reflect.ValueOf(v))

	if err != nil {

		p.err = err
		return p

	}

	mapper := MapperImpl{
		root:    sn,
		pk:      searchFieldWithExpression(sn.Childs, "pk"),
		sk:      searchFieldWithExpression(sn.Childs, "sk"),
		cache:   make(map[string]*parser.StructNode),
		divider: p.divider,
	}

	for i, c := range sn.Childs {
		if c.HasTag("cm") {
			if v, ok := c.Tag["cm"].GetNamed()["name"]; ok {
				mapper.cache[v] = &sn.Childs[i]
			}
		}
	}

	p.mappers[reflect.TypeOf(v)] = &mapper

	return p
}

// searchFieldWithExpression searches the cm tag if the _name_ do exists in the named parameter map.
func searchFieldWithExpression(childs []parser.StructNode, name string) *parser.StructNode {

	for _, c := range childs {
		if c.HasTag("cm") {

			if _, ok := c.Tag["cm"].GetNamed()[name]; ok {
				return &c
			}
		}

		if c.HasChildren() {
			return searchFieldWithExpression(c.Childs, name)
		}
	}

	return nil
}
