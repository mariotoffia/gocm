package idreflect

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/mariotoffia/gocm/cmid"
	"github.com/mariotoffia/gocm/cmid/idexpr"
	"github.com/mariotoffia/ssm/parser"
)

// ProjectRepositoryImpl implements the `cmid.IDProjectRepository`
type ProjectRepositoryImpl struct {
	tp      *parser.Parser
	mappers map[reflect.Type]*Projector
	err     error
	divider string
	tag     string
}

// Projector implements the `cmid.IDProjector`
type Projector struct {
	root    *parser.StructNode
	expr    cmid.Identity
	cache   map[string]*parser.StructNode
	divider string
	div     byte
	tag     string
}

// NewProjectRepository creates a new `ProjectRepositoryImpl`
func NewProjectRepository() *ProjectRepositoryImpl {

	return &ProjectRepositoryImpl{
		tp: parser.New("", "", "").
			RegisterTagParser(
				cmid.IDStandardCMTag, parser.NewTagParser([]string{"name", "pk", "sk"}),
			),
		mappers: map[reflect.Type]*Projector{},
		divider: cmid.IDStandardDivider,
		tag:     cmid.IDStandardCMTag,
	}

}

// Error returns any error state
func (pr *ProjectRepositoryImpl) Error() error {
	return pr.err
}

// ClearError will clear any error state
func (pr *ProjectRepositoryImpl) ClearError() *ProjectRepositoryImpl {
	pr.err = nil
	return pr
}

// UseDivider changes the standard divider `cmid.IDStandarDivider` to _divider_.
//
// The already `Add()`ed `IDObjectMapper`s are using the previous and all new will
// use the newly set _divider_. Hence, it is possible to have different dividers in
// same `MapperParser`.
func (pr *ProjectRepositoryImpl) UseDivider(divider string) *ProjectRepositoryImpl {
	pr.divider = divider
	return pr
}

// UseTag changes the standard tag name to control IDMapping.
//
// Default is `identityIDStandardCMTag`. This is quite expensive
// operation since it will allow a new tag parser with the tag name.
func (pr *ProjectRepositoryImpl) UseTag(tag string) *ProjectRepositoryImpl {
	pr.tag = tag

	pr.tp = parser.New("", "", "").
		RegisterTagParser(tag,
			parser.NewTagParser([]string{"name", "pk", "sk"}),
		)

	return pr
}

// AddProjection will add a projection using the cm attributes but uses the provided
// _PK_ and _SK_ instead.
func (pr *ProjectRepositoryImpl) AddProjection(
	v interface{},
	expr cmid.Identity) *ProjectRepositoryImpl {

	reflect.TypeOf(v)
	sn, err := pr.tp.Parse(reflect.ValueOf(v))

	if err != nil {

		pr.err = err
		return pr

	}

	mapper := Projector{
		root:    sn,
		cache:   make(map[string]*parser.StructNode),
		divider: pr.divider,
		div:     pr.divider[0],
		tag:     pr.tag,
		expr:    expr,
	}

	if expr.PartitionKey() == "" {

		pr.err = fmt.Errorf("could not find mandatory cm, PartitionKey expression on type %T", v)
		return pr

	}

	for i, c := range sn.Childs {
		if c.HasTag(pr.tag) {
			if v, ok := c.Tag[pr.tag].GetNamed()["name"]; ok {
				mapper.cache[v] = &sn.Childs[i]
			}
		}
	}

	pr.mappers[reflect.TypeOf(v)] = &mapper

	return pr
}

// Projector gets the projector for the parameter _v_ (as it was registered using `AddProjection()`).
func (pr *ProjectRepositoryImpl) Projector(v interface{}) *Projector {

	if m, ok := pr.mappers[reflect.TypeOf(v)]; ok {
		return m
	}

	return nil
}

// Project accepts a instance pointer _v_ that it will produce an `Identity` with
// the PK and SK projection expression.
//
// The returned `Identity` may be used by any implementation to use as _PK_ and _SK_ instead
// of the default `IDMapper` / `IDObjectMapper`.
func (p *Projector) Project(v interface{}) (*cmid.ID, error) {
	return p.assembleIdentity(reflect.ValueOf(v).Elem())
}

func (p *Projector) assembleIdentity(vval reflect.Value) (*cmid.ID, error) {

	id := cmid.ID{}
	pk := true
	skipdiv := false

	var expr *string
	var err error
	var sn *parser.StructNode

	pcb := &idexpr.ParserCallback{
		DivFunc: func(index int) {

			if skipdiv || err != nil {
				return
			}

			if pk {
				id.PK += p.divider
			} else {
				id.SK += p.divider
			}

		},
		CompFunc: func(index int, component string) {

			if pk {
				id.PK += component
			} else {
				id.SK += component
			}

		},
		IDFunc: func(index int, optional bool, name string) {

			sn = p.cache[name]

			if sn == nil {
				sn = p.searchFieldByCMName(p.root.Childs, name)
			}

			if nil == sn {

				if optional {
					skipdiv = true
				} else {
					err = fmt.Errorf("invalid expression %s, missing parameter named: %s", *expr, name)
				}

				return
			}

			val := vval.FieldByIndex(sn.Field.Index).Interface()

			var str string
			switch d := val.(type) {
			case string:
				str = d
			case int:
				str = strconv.FormatInt(int64(d), 10)
			case uint:
				str = strconv.FormatUint(uint64(d), 10)
			case int8:
				str = strconv.FormatInt(int64(d), 10)
			case int16:
				str = strconv.FormatInt(int64(d), 10)
			case int32:
				str = strconv.FormatInt(int64(d), 10)
			case int64:
				str = strconv.FormatInt(d, 10)
			case uint8:
				str = strconv.FormatUint(uint64(d), 10)
			case uint16:
				str = strconv.FormatUint(uint64(d), 10)
			case uint32:
				str = strconv.FormatUint(uint64(d), 10)
			case uint64:
				str = strconv.FormatUint(d, 10)
			default:
				str = fmt.Sprint(d)
			}

			if str == "" {

				if optional {
					skipdiv = true
				} else {

					err = fmt.Errorf(
						"could not map expression %s since field: %s is not optional and do not contain any value",
						*expr, name,
					)

				}

				return
			}

			if pk {
				id.PK += str
			} else {
				id.SK += str
			}
		},
	}

	e := p.expr.PartitionKey()
	expr = &e
	idexpr.Parse(p.divider, *expr, pcb)

	if p.expr.SecondaryKey() != "" {
		e := p.expr.SecondaryKey()
		expr = &e
		pk = false
		idexpr.Parse(p.divider, *expr, pcb)
	}

	pklen := len(id.PK) - 1
	if id.PK[pklen] == p.div {
		id.PK = id.PK[:pklen]
	}

	if id.SK != "" {
		sklen := len(id.SK) - 1
		if id.SK[sklen] == p.div {
			id.SK = id.SK[:sklen]
		}
	}

	return &id, err
}

// searchFieldByCMName searches the cm tag if the _name_ do match the named parameter name.
func (p *Projector) searchFieldByCMName(childs []parser.StructNode, name string) *parser.StructNode {

	for _, c := range childs {
		if c.HasTag(p.tag) {

			if v, ok := c.Tag[p.tag].GetNamed()["name"]; ok && v == name {
				return &c
			}
		}

		if c.HasChildren() {
			return p.searchFieldByCMName(c.Childs, name)
		}
	}

	return nil
}
