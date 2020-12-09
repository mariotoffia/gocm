package idreflect

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/mariotoffia/gocm/cmid"
	"github.com/mariotoffia/gocm/cmid/idexpr"
	"github.com/mariotoffia/ssm/parser"
)

// MapperImpl is a reflection based `cmid.IDMapper`
type MapperImpl struct {
	root    *parser.StructNode
	pk      *parser.StructNode
	sk      *parser.StructNode
	cache   map[string]*parser.StructNode
	divider string
	div     byte
	tag     string
}

// Resolve will resolve any fields that are part of the PK and SK and set those from
// the PK and SK attributes from the provided instance.
//
// This function is never idempotent since it will change the PK or SK. The other fields
// is overwritten. But the PK and SK is always changed (if not static value).
func (rm *MapperImpl) Resolve(v interface{}) error {

	vval := reflect.ValueOf(v).Elem()
	id, err := rm.extractIdentity(vval)

	if id.PK == "" {
		return fmt.Errorf("No PartitionKey detected in instance %T", v)
	}

	if err != nil {
		return err
	}

	var sn *parser.StructNode

	pk := true
	idx := 0
	optionalIncluded := true
	pkv := strings.Split(id.PK, rm.divider)
	skv := strings.Split(id.SK, rm.divider)

	pcb := &idexpr.ParserCallback{
		DivFunc: func(index int) {
		},
		CompFunc: func(index int, component string) {
			idx++
		},
		IDFunc: func(index int, optional bool, name string) {

			if err != nil {
				return
			}

			if optional && !optionalIncluded {
				return
			}

			sn = rm.cache[name]

			if sn == nil {
				sn = rm.searchFieldByCMName(rm.root.Childs, name)
			}

			var val string
			if pk {
				val = pkv[idx]
			} else {
				val = skv[idx]
			}

			idx++

			err = reflectSetValue(sn, vval, val)
		},
	}

	expr := rm.pk.Tag[rm.tag].GetNamed()["pk"]

	optionalIncluded = len(strings.Split(id.PK, rm.divider)) == len(strings.Split(expr, rm.divider))

	idexpr.Parse(rm.divider, expr, pcb)

	if rm.sk != nil {

		expr = rm.sk.Tag[rm.tag].GetNamed()["sk"]
		pk = false
		idx = 0

		optionalIncluded = len(strings.Split(id.SK, rm.divider)) == len(strings.Split(expr, rm.divider))

		idexpr.Parse(rm.divider, expr, pcb)
	}

	return err
}

func reflectSetValue(sn *parser.StructNode, vval reflect.Value, val string) error {

	switch sn.Field.Type.Kind() {
	case reflect.String:
		vval.FieldByIndex(sn.Field.Index).SetString(val)
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:

		i, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return err
		}

		vval.FieldByIndex(sn.Field.Index).SetInt(i)

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64:

		i, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			return err
		}

		vval.FieldByIndex(sn.Field.Index).SetUint(i)

	case reflect.Float32, reflect.Float64:
		i, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return err
		}

		vval.FieldByIndex(sn.Field.Index).SetFloat(i)
	}

	return nil
}

// Map will map the sub-properties in the instance to form a PK and
// possibly a SK. This is done one the existing instance, so it should
// not be used anything else than CM communication.
//
// This function may not be idempotent, use a fresh new instance each time
// invoking this function.
func (rm *MapperImpl) Map(v interface{}) error {

	vval := reflect.ValueOf(v).Elem()
	id, err := rm.assembleIdentity(vval)

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

// ExtractIdentity will extract the PK and SK from the _v_ parameter.
func (rm *MapperImpl) ExtractIdentity(v interface{}) (*cmid.ID, error) {
	return rm.extractIdentity(reflect.ValueOf(v).Elem())
}

// AssembleIdentity will create a new identity from the _v_ instance data.
func (rm *MapperImpl) AssembleIdentity(v interface{}) (*cmid.ID, error) {
	return rm.assembleIdentity(reflect.ValueOf(v).Elem())
}

func (rm *MapperImpl) extractIdentity(vval reflect.Value) (*cmid.ID, error) {

	id := cmid.ID{}

	if rm.pk != nil {
		id.PK = vval.FieldByIndex(rm.pk.Field.Index).Interface().(string)
	}

	if rm.sk != nil {
		id.SK = vval.FieldByIndex(rm.sk.Field.Index).Interface().(string)
	}

	return &id, nil
}

func (rm *MapperImpl) assembleIdentity(vval reflect.Value) (*cmid.ID, error) {

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
				id.PK += rm.divider
			} else {
				id.SK += rm.divider
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
				sn = rm.searchFieldByCMName(rm.root.Childs, name)
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
		e := rm.pk.Tag[rm.tag].GetNamed()["pk"]
		expr = &e
		idexpr.Parse(rm.divider, *expr, pcb)
	}

	if rm.sk != nil {
		e := rm.sk.Tag[rm.tag].GetNamed()["sk"]
		expr = &e
		pk = false
		idexpr.Parse(rm.divider, *expr, pcb)
	}

	pklen := len(id.PK) - 1
	if id.PK[pklen] == rm.div {
		id.PK = id.PK[:pklen]
	}

	if rm.sk != nil {
		sklen := len(id.SK) - 1
		if id.SK[sklen] == rm.div {
			id.SK = id.SK[:sklen]
		}
	}

	return &id, err
}

// searchFieldByCMName searches the cm tag if the _name_ do match the named parameter name.
func (rm *MapperImpl) searchFieldByCMName(childs []parser.StructNode, name string) *parser.StructNode {

	for _, c := range childs {
		if c.HasTag(rm.tag) {

			if v, ok := c.Tag[rm.tag].GetNamed()["name"]; ok && v == name {
				return &c
			}
		}

		if c.HasChildren() {
			return rm.searchFieldByCMName(c.Childs, name)
		}
	}

	return nil
}
