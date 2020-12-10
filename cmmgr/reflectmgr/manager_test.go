package reflectmgr

import (
	"testing"

	"github.com/mariotoffia/gocm/cmid"
	"github.com/stretchr/testify/assert"
)

func TestSimpleFactoryAndMappingResolving(t *testing.T) {

	type Car struct {
		Brand string `cm:"pk, pk=P#{pk}"`
		SK    string `cm:"sk, sk=S#{year}#{model}"`
		Model string `cm:"model"`
		Year  int    `cm:"year"`
	}

	var test Car
	mgr := NewReflectiveEntityManager().Register(&test)
	mgr.Freeze()

	car := Car{
		Brand: "SAAB",
		Model: "9-5",
		Year:  2001,
	}

	// Map
	err := mgr.Mappers().Mapper(&test).Map(&car)
	assert.Equal(t, nil, err)

	assert.Equal(t, "P#SAAB", car.Brand)
	assert.Equal(t, "S#2001#9-5", car.SK)

	// Resolve
	car.Model = ""
	car.Year = 0

	err = mgr.Mappers().Mapper(&test).Resolve(&car)
	assert.Equal(t, nil, err)
	assert.Equal(t, "SAAB", car.Brand)
	assert.Equal(t, "9-5", car.Model)
	assert.Equal(t, 2001, car.Year)

	// Factory
	i := mgr.Factories().Create(&cmid.ID{PK: "P", SK: "S"}).(*Car)
	assert.NotNil(t, i)

	i.Brand = "SAAB"
	assert.Equal(t, "SAAB", i.Brand, "Make sure we have a valid pointer")
}

func BenchmarkSimpleFactoryAndMappingResolving(t *testing.B) {

	type Car struct {
		Brand string `cm:"pk, pk=P#{pk}"`
		SK    string `cm:"sk, sk=S#{year}#{model}"`
		Model string `cm:"model"`
		Year  int    `cm:"year"`
	}

	var test Car
	mgr := NewReflectiveEntityManager().Register(&test)
	mgr.Freeze()

	car := Car{
		Brand: "SAAB",
		Model: "9-5",
		Year:  2001,
	}

	id := &cmid.ID{PK: "P", SK: "S"}

	for i := 0; i < t.N; i++ {
		mgr.Mappers().Mapper(&test).Map(&car)
		mgr.Mappers().Mapper(&test).Resolve(&car)
		mgr.Factories().Create(id)
	}
}
