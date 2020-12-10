package idreflect

import (
	"testing"

	"github.com/mariotoffia/gocm/cmid"
	"github.com/stretchr/testify/assert"
)

func TestSimplePKSKFieldProjection(t *testing.T) {

	type Car struct {
		Brand string `cm:"pk, pk=P#{pk}"`
		SK    string `cm:"sk, sk=S#{year}#{model}"`
		Model string `cm:"model"`
		Year  int    `cm:"year"`
	}

	p := NewProjectRepository().UseDivider("/")

	var test Car
	m := p.AddProjection(&test, &cmid.ID{
		PK: "1.0.0/{pk}/{year}/{model}",
	})

	assert.Equal(t, nil, p.Error())

	car := Car{
		Brand: "SAAB",
		Model: "9-5",
		Year:  2001,
	}

	id, err := m.Projection(&test).Project(&car)
	assert.Equal(t, nil, err)
	assert.Equal(t, "1.0.0/SAAB/2001/9-5", id.PartitionKey())
	assert.Equal(t, "", id.SecondaryKey())
}

func TestUnresolvedEntityWillFailIfPKPropertyIsIncorrect(t *testing.T) {

	type Car struct {
		Brand string `cm:"pk, pk=P#{pk}"`
		SK    string `cm:"sk, sk=S#{year}#{model}"`
		Model string `cm:"model"`
		Year  int    `cm:"year"`
	}

	car := Car{
		Brand: "SAAB",
		Model: "9-5",
		Year:  2001,
	}

	var test Car
	err := NewParser().Add(&test).Mapper(&test).Map(&car)
	assert.Equal(t, nil, err)

	p := NewProjectRepository().UseDivider("/")

	m := p.AddProjection(&test, &cmid.ID{
		PK: "1.0.0/{pk}/{year}/{model}",
	})

	assert.Equal(t, nil, p.Error())

	id, err := m.Projection(&test).Project(&car)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, "1.0.0/SAAB/2001/9-5", id.PartitionKey(), "since car is not resolved")
	assert.Equal(t, "", id.SecondaryKey())
}

func TestResolvedEntityCanProject(t *testing.T) {

	type Car struct {
		Brand string `cm:"pk, pk=P#{pk}"`
		SK    string `cm:"sk, sk=S#{year}#{model}"`
		Model string `cm:"model"`
		Year  int    `cm:"year"`
	}

	car := Car{
		Brand: "SAAB",
		Model: "9-5",
		Year:  2001,
	}

	var test Car
	m := NewParser().Add(&test).Mapper(&test)
	err := m.Map(&car)
	assert.Equal(t, nil, err)

	err = m.Resolve(&car)
	assert.Equal(t, nil, err)

	p := NewProjectRepository().UseDivider("/")

	prj := p.AddProjection(&test, &cmid.ID{
		PK: "1.0.0/{pk}/{year}/{model}",
	})

	assert.Equal(t, nil, p.Error())

	id, err := prj.Projection(&test).Project(&car)
	assert.Equal(t, nil, err)
	assert.Equal(t, "1.0.0/SAAB/2001/9-5", id.PartitionKey(), "since car is resolved")
	assert.Equal(t, "", id.SecondaryKey())
}

func BenchmarkSimplePKSKFieldProjection(t *testing.B) {

	type Car struct {
		Brand string `cm:"pk, pk=P#{pk}"`
		SK    string `cm:"sk, sk=S#{year}#{model}"`
		Model string `cm:"model"`
		Year  int    `cm:"year"`
	}

	p := NewProjectRepository().UseDivider("/")

	var test Car
	m := p.AddProjection(&test, &cmid.ID{
		PK: "1.0.0/{pk}/{year}/{model}",
	})

	assert.Equal(t, nil, p.Error())

	car := Car{
		Brand: "SAAB",
		Model: "9-5",
		Year:  2001,
	}

	prj := m.Projection(&test)

	t.ResetTimer()

	for i := 0; i < t.N; i++ {
		prj.Project(&car)
	}

}
