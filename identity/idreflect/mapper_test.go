package idreflect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimplePKSKFieldMapping(t *testing.T) {

	type Car struct {
		Brand string `cm:"pk, pk=P#{pk}"`
		SK    string `cm:"sk, sk=S#{year}#{model}"`
		Model string `cm:"model"`
		Year  int    `cm:"year"`
	}

	p := NewParser()

	var test Car
	m := p.Add(&test).Mapper(&test)

	assert.Equal(t, nil, p.Error())

	car := Car{
		Brand: "SAAB",
		Model: "9-5",
		Year:  2001,
	}

	id, err := m.Identity(&car)
	assert.Equal(t, nil, err)
	assert.Equal(t, "P#SAAB", id.PK)
	assert.Equal(t, "S#2001#9-5", id.SK)
}

func TestMapSimplePKSKFieldMapping(t *testing.T) {

	type Car struct {
		Brand string `cm:"pk, pk=P#{pk}"`
		SK    string `cm:"sk, sk=S#{year}#{model}"`
		Model string `cm:"model"`
		Year  int    `cm:"year"`
	}

	p := NewParser()

	var test Car
	m := p.Add(&test).Mapper(&test)

	assert.Equal(t, nil, p.Error())

	car := Car{
		Brand: "SAAB",
		Model: "9-5",
		Year:  2001,
	}

	err := m.Map(&car)
	assert.Equal(t, nil, err)
	assert.Equal(t, "P#SAAB", car.Brand)
	assert.Equal(t, "S#2001#9-5", car.SK)
}

func TestSimpleMappingWithEndingOptionalID(t *testing.T) {

	type Car struct {
		Brand string `cm:"pk, pk=P#{pk}"`
		SK    string `cm:"sk, sk=S#{year}#{?model}"`
		Model string `cm:"model"`
		Year  int    `cm:"year"`
	}

	p := NewParser()

	var test Car
	m := p.Add(&test).Mapper(&test)

	assert.Equal(t, nil, p.Error())

	car := Car{
		Brand: "SAAB",
		Model: "",
		Year:  2001,
	}

	id, err := m.Identity(&car)
	assert.Equal(t, nil, err)
	assert.Equal(t, "P#SAAB", id.PK)
	assert.Equal(t, "S#2001", id.SK)
}

func TestSimpleMappingWithMiddleOptionalID(t *testing.T) {

	type Car struct {
		Brand string `cm:"pk, pk=P#{pk}"`
		SK    string `cm:"sk, sk=S#{?model}#{year}"`
		Model string `cm:"model"`
		Year  int    `cm:"year"`
	}

	p := NewParser()

	var test Car
	m := p.Add(&test).Mapper(&test)

	assert.Equal(t, nil, p.Error())

	car := Car{
		Brand: "SAAB",
		Model: "",
		Year:  2001,
	}

	id, err := m.Identity(&car)
	assert.Equal(t, nil, err)
	assert.Equal(t, "P#SAAB", id.PK)
	assert.Equal(t, "S#2001", id.SK)
}

func BenchmarkSimplePKSKFieldMapping(t *testing.B) {

	type Car struct {
		Brand string `cm:"pk, pk=P#{pk}"`
		SK    string `cm:"sk, sk=S#{year}#{model}"`
		Model string `cm:"model"`
		Year  int    `cm:"year"`
	}

	p := NewParser()

	var test Car
	m := p.Add(&test).Mapper(&test)

	assert.Equal(t, nil, p.Error())

	test = Car{
		Brand: "SAAB",
		Model: "9-5",
		Year:  2001,
	}

	t.ResetTimer()

	for i := 0; i < t.N; i++ {
		m.Identity(&test)
	}
}

func BenchmarkMapSimplePKSKFieldMapping(t *testing.B) {

	type Car struct {
		Brand string `cm:"pk, pk=P#{pk}"`
		SK    string `cm:"sk, sk=S#{year}#{model}"`
		Model string `cm:"model"`
		Year  int    `cm:"year"`
	}

	p := NewParser()

	var test Car
	m := p.Add(&test).Mapper(&test)

	assert.Equal(t, nil, p.Error())

	t.ResetTimer()

	for i := 0; i < t.N; i++ {

		car := Car{
			Brand: "SAAB",
			Model: "9-5",
			Year:  2001,
		}

		m.Map(&car)
	}
}
