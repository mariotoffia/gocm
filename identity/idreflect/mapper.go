package idreflect

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/mariotoffia/gocm/identity"
	"github.com/mariotoffia/gocm/identity/idexpr"
	"github.com/mariotoffia/ssm/parser"
)

// MapperImpl is a reflection based `identity.IDMapper`
type MapperImpl struct {
	root  *parser.StructNode
	pk    *parser.StructNode
	sk    *parser.StructNode
	cache map[string]*parser.StructNode
}

// Map will map the sub-properties in the instance to form a PK and
// possibly a SK. This is done one the existing instance, so it should
// not be used anything else than CM communication.
//
// This function may not be idempotent, use a fresh new instance each time
// invoking this function.
func (rm *MapperImpl) Map(v interface{}) error {

	vval := reflect.ValueOf(v).Elem()
	id, err := rm.identity(vval)

	if err != nil {
		return err
	}

	if id.PK != "" {
		vval.FieldByIndex(rm.pk.Field.Index).SetString(id.PK)
	}

	if id.SK != "" {
		vval.FieldByIndex(rm.sk.Field.Index).SetString(id.SK)
	}

	return nil
}

// Identity will create a new identity from the _v_ instance data.
func (rm *MapperImpl) Identity(v interface{}) (*identity.ID, error) {
	return rm.identity(reflect.ValueOf(v).Elem())
}

func (rm *MapperImpl) identity(vval reflect.Value) (*identity.ID, error) {

	id := identity.ID{}
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
				id.PK += identity.IDStandardDivider
			} else {
				id.SK += identity.IDStandardDivider
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

			sn = rm.cache[name]

			if sn == nil {
				sn = searchFieldByCMName(rm.root.Childs, name)
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

	if rm.pk != nil {
		e := rm.pk.Tag["cm"].GetNamed()["pk"]
		expr = &e
		idexpr.Parse(*expr, pcb)
	}

	if rm.sk != nil {
		e := rm.sk.Tag["cm"].GetNamed()["sk"]
		expr = &e
		pk = false
		idexpr.Parse(*expr, pcb)
	}

	pklen := len(id.PK) - 1
	if id.PK[pklen] == '#' {
		id.PK = id.PK[:pklen]
	}

	sklen := len(id.SK) - 1
	if id.SK[sklen] == '#' {
		id.SK = id.SK[:sklen]
	}

	return &id, err
}

// searchFieldByCMName searches the cm tag if the _name_ do match the named parameter name.
func searchFieldByCMName(childs []parser.StructNode, name string) *parser.StructNode {

	for _, c := range childs {
		if c.HasTag("cm") {

			if v, ok := c.Tag["cm"].GetNamed()["name"]; ok && v == name {
				return &c
			}
		}

		if c.HasChildren() {
			return searchFieldWithExpression(c.Childs, name)
		}
	}

	return nil
}
