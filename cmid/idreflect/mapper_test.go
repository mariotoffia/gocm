package idreflect

import (
	"fmt"
	"testing"

	"github.com/mariotoffia/gocm/cmid"
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

	id, err := m.AssembleIdentity(&car)
	assert.Equal(t, nil, err)
	assert.Equal(t, "P#SAAB", id.PartitionKey())
	assert.Equal(t, "S#2001#9-5", id.SecondaryKey())
}

func TestExtractIdentitySimplePKSKFieldMapping(t *testing.T) {

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

	rid, err := m.ExtractIdentity(&car)

	assert.Equal(t, nil, err)
	assert.Equal(t, "P#SAAB", rid.PartitionKey())
	assert.Equal(t, "S#2001#9-5", rid.SecondaryKey())
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

func TestMapIdentityCustomTagsAndSeparator(t *testing.T) {

	type Car struct {
		Brand string `s3:"pk, pk=1.0.0/{pk}/{year}/{model}"`
		SK    string
		Model string `s3:"model"`
		Year  int    `s3:"year"`
	}

	p := NewParser().UseDivider("/").UseTag("s3")

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
	assert.Equal(t, "1.0.0/SAAB/2001/9-5", car.Brand)
	assert.Equal(t, "", car.SK)
}

func TestResolveSimplePKSKFieldMapping(t *testing.T) {

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
	car.Model = ""
	car.Year = 0

	// Now to the actual test
	err = m.Resolve(&car)
	assert.Equal(t, nil, err)
	assert.Equal(t, "SAAB", car.Brand)
	assert.Equal(t, "9-5", car.Model)
	assert.Equal(t, 2001, car.Year)
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

	id, err := m.AssembleIdentity(&car)
	assert.Equal(t, nil, err)
	assert.Equal(t, "P#SAAB", id.PartitionKey())
	assert.Equal(t, "S#2001", id.SecondaryKey())
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

	id, err := m.AssembleIdentity(&car)
	assert.Equal(t, nil, err)
	assert.Equal(t, "P#SAAB", id.PartitionKey())
	assert.Equal(t, "S#2001", id.SecondaryKey())
}

func TestResolveSimpleMappingWithMiddleOptionalIDMissing(t *testing.T) {

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

	err := m.Map(&car)
	assert.Equal(t, nil, err)
	car.Model = ""
	car.Year = 0

	err = m.Resolve(&car)
	assert.Equal(t, nil, err)
	assert.Equal(t, "SAAB", car.Brand)
	assert.Equal(t, "", car.Model)
	assert.Equal(t, 2001, car.Year)

}

func TestResolveSimpleMappingWithMiddleOptionalID(t *testing.T) {

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
		Model: "9-5",
		Year:  2001,
	}

	err := m.Map(&car)
	assert.Equal(t, nil, err)
	car.Model = ""
	car.Year = 0

	err = m.Resolve(&car)
	assert.Equal(t, nil, err)
	assert.Equal(t, "SAAB", car.Brand)
	assert.Equal(t, "9-5", car.Model)
	assert.Equal(t, 2001, car.Year)

}

func TestEntityFactoryIDComponentShallMatchPKSKInObject(t *testing.T) {

	type Car struct {
		Brand string `cm:"pk, pk=P#{pk}"`
		SK    string `cm:"sk, sk=S#{year}#{model}"`
		Model string `cm:"model"`
		Year  int    `cm:"year"`
	}

	p := NewParser()

	var test Car
	m := p.Add(&test).Mapper(&test)

	assert.Equal(t, "P", m.Identities()[0].PartitionKey())
	assert.Equal(t, "S", m.Identities()[0].SecondaryKey())

}

func TestEntityFactoryShallReturnInstanceOnCorrectPKSK(t *testing.T) {

	type Car struct {
		Brand string `cm:"pk, pk=P#{pk}"`
		SK    string `cm:"sk, sk=S#{year}#{model}"`
		Model string `cm:"model"`
		Year  int    `cm:"year"`
	}

	p := NewParser()

	var test Car
	m := p.Add(&test).Mapper(&test)

	i := m.Create(&cmid.ID{PK: "P", SK: "S"})

	assert.NotNil(t, i)
	assert.Equal(t, "*idreflect.Car", fmt.Sprintf("%T", i))

	car := i.(*Car)
	car.Brand = "SAAB"
	car.Model = "9-5"
	car.Year = 2001

	assert.Equal(t, "SAAB", car.Brand, "Just to make sure it is a valid pointer")
}

func TestEntityFactoryShallReturnNilOnIncorrectPKSK(t *testing.T) {

	type Car struct {
		Brand string `cm:"pk, pk=P#{pk}"`
		SK    string `cm:"sk, sk=S#{year}#{model}"`
		Model string `cm:"model"`
		Year  int    `cm:"year"`
	}

	p := NewParser()

	var test Car
	m := p.Add(&test).Mapper(&test)

	i := m.Create(&cmid.ID{PK: "P", SK: "A"})

	assert.Nil(t, i)
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
		m.AssembleIdentity(&test)
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

func BenchmarkResolveSimplePKSKFieldMapping(t *testing.B) {

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
	pk := car.Brand

	t.ResetTimer()

	for i := 0; i < t.N; i++ {
		m.Resolve(&car)

		car.Brand = pk
	}
}

func BenchmarkSimpleEntityFactoryCreate(t *testing.B) {

	type Car struct {
		Brand string `cm:"pk, pk=P#{pk}"`
		SK    string `cm:"sk, sk=S#{year}#{model}"`
		Model string `cm:"model"`
		Year  int    `cm:"year"`
	}

	p := NewParser()

	var test Car
	m := p.Add(&test).Mapper(&test)
	id := cmid.ID{PK: "P", SK: "S"}

	t.ResetTimer()

	for i := 0; i < t.N; i++ {
		m.Create(&id)
	}
}
