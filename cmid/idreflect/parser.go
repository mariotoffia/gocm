package idreflect

import (
	"fmt"
	"reflect"

	"github.com/mariotoffia/gocm/cmid"
	"github.com/mariotoffia/gocm/cmid/idexpr"
	"github.com/mariotoffia/ssm/parser"
)

// MapperParser parser that uses _cm_ tags.
type MapperParser struct {
	tp      *parser.Parser
	mappers map[reflect.Type]*MapperImpl
	cache   []cmid.IDObjectMapper
	divider string
	tag     string
	err     error
	frozen  bool
}

// NewParser creates a new `MapperParser`
func NewParser() *MapperParser {

	return &MapperParser{
		tp: parser.New("", "", "").
			RegisterTagParser(cmid.IDStandardCMTag, parser.NewTagParser([]string{"name", "pk", "sk"})),
		mappers: map[reflect.Type]*MapperImpl{},
		divider: cmid.IDStandardDivider,
		tag:     cmid.IDStandardCMTag,
		cache:   []cmid.IDObjectMapper{},
	}

}

// Freeze will make registration of new mappings impossible.
func (p *MapperParser) Freeze() cmid.IDMapperRepository {
	return p
}

// IsFrozen returns `true` if the instance do not accept any more mapping registrations using `Add()`.
func (p *MapperParser) IsFrozen() bool {
	return p.frozen
}

// Mappers returns an array of currently supported mappings.
func (p *MapperParser) Mappers() []cmid.IDObjectMapper {
	return p.cache
}

// UseDivider changes the standard divider `cmid.IDStandardDivider` to _divider_.
//
// The already `Add()`ed `IDObjectMapper`s are using the previous and all new will
// use the newly set _divider_. Hence, it is possible to have different dividers in
// same `MapperParser`.
func (p *MapperParser) UseDivider(divider string) *MapperParser {
	p.divider = divider
	return p
}

// UseTag changes the standard tag name to control IDMapping.
//
// Default is `identityIDStandardCMTag`. This is quite expensive
// operation since it will allow a new tag parser with the tag name.
func (p *MapperParser) UseTag(tag string) *MapperParser {
	p.tag = tag

	p.tp = parser.New("", "", "").
		RegisterTagParser(tag,
			parser.NewTagParser([]string{"name", "pk", "sk"}),
		)

	return p
}

// Error returns any error state
func (p *MapperParser) Error() error {
	return p.err
}

// ClearError will clear any error state
func (p *MapperParser) ClearError() cmid.IDMapperRepository {
	p.err = nil
	return p
}

// Mapper gets the mapper for the parameter _v_ (as it was registered using `Add()`).
func (p *MapperParser) Mapper(v any) cmid.IDObjectMapper {

	if m, ok := p.mappers[reflect.TypeOf(v)]; ok {
		return m
	}

	return nil
}

// Add creates a new reflection based `IdMapper` and stores it in its internal hash.
//
// The parameter _v_ is expected to be a pointer to a type.
func (p *MapperParser) Add(v any) cmid.IDMapperRepository {

	if p.frozen {
		return p
	}

	reflect.TypeOf(v)
	sn, err := p.tp.Parse(reflect.ValueOf(v))

	if err != nil {

		p.err = err
		return p

	}

	mapper := MapperImpl{
		root:    sn,
		pk:      p.searchFieldWithExpression(sn.Childs, "pk"),
		sk:      p.searchFieldWithExpression(sn.Childs, "sk"),
		cache:   make(map[string]*parser.StructNode),
		divider: p.divider,
		div:     p.divider[0],
		tag:     p.tag,
	}

	if mapper.pk == nil {

		p.err = fmt.Errorf("could not find mandatory cm, pk expression on type %T", v)
		return p

	}

	id := &cmid.ID{PK: mapper.pk.Tag[p.tag].GetNamed()["pk"]}

	if mapper.sk != nil {
		id.SK = mapper.sk.Tag[p.tag].GetNamed()["sk"]
	}

	mapper.id = []cmid.ComponentIdentity{p.stripToComponents(id)}

	for i, c := range sn.Childs {
		if c.HasTag(p.tag) {
			if v, ok := c.Tag[p.tag].GetNamed()["name"]; ok {
				mapper.cache[v] = &sn.Childs[i]
			}
		}
	}

	p.mappers[reflect.TypeOf(v)] = &mapper
	p.cache = append(p.cache, &mapper)

	return p
}

// stripToComponents strips out everything except for the component parts and
// returns a new instance of `cmid.ID`
func (p *MapperParser) stripToComponents(id *cmid.ID) *cmid.ID {

	cid := cmid.ID{}
	pk := true

	pcb := &idexpr.ParserCallback{
		CompFunc: func(index int, component string) {

			if pk {
				if cid.PK == "" {
					cid.PK = component
				} else {
					cid.PK += p.divider + component
				}
				return
			}

			if cid.SK == "" {
				cid.SK = component
			} else {
				cid.SK += p.divider + component
			}
		},
		DivFunc: func(index int) {},
		IDFunc:  func(index int, optional bool, name string) {},
	}

	idexpr.Parse(p.divider, id.PK, pcb)

	if id.SK != "" {
		pk = false
		idexpr.Parse(p.divider, id.SK, pcb)
	}

	return &cid
}

// searchFieldWithExpression searches the cm tag if the _name_ do exists in the named parameter map.
func (p *MapperParser) searchFieldWithExpression(childs []parser.StructNode, name string) *parser.StructNode {

	for _, c := range childs {
		if c.HasTag(p.tag) {

			if _, ok := c.Tag[p.tag].GetNamed()[name]; ok {
				return &c
			}
		}

		if c.HasChildren() {
			return p.searchFieldWithExpression(c.Childs, name)
		}
	}

	return nil
}
