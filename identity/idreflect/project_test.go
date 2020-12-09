package idreflect

import (
	"testing"

	"github.com/mariotoffia/gocm/identity"
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
	m := p.AddProjection(&test, &identity.ID{
		PK: "1.0.0/{pk}/{year}/{model}",
	})

	assert.Equal(t, nil, p.Error())

	car := Car{
		Brand: "SAAB",
		Model: "9-5",
		Year:  2001,
	}

	id, err := m.Projector(&test).Project(&car)
	assert.Equal(t, nil, err)
	assert.Equal(t, "1.0.0/SAAB/2001/9-5", id.PK)
	assert.Equal(t, "", id.SK)
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

	m := p.AddProjection(&test, &identity.ID{
		PK: "1.0.0/{pk}/{year}/{model}",
	})

	assert.Equal(t, nil, p.Error())

	id, err := m.Projector(&test).Project(&car)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, "1.0.0/SAAB/2001/9-5", id.PK, "since car is not resolved")
	assert.Equal(t, "", id.SK)
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

	prj := p.AddProjection(&test, &identity.ID{
		PK: "1.0.0/{pk}/{year}/{model}",
	})

	assert.Equal(t, nil, p.Error())

	id, err := prj.Projector(&test).Project(&car)
	assert.Equal(t, nil, err)
	assert.Equal(t, "1.0.0/SAAB/2001/9-5", id.PK, "since car is resolved")
	assert.Equal(t, "", id.SK)
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
	m := p.AddProjection(&test, &identity.ID{
		PK: "1.0.0/{pk}/{year}/{model}",
	})

	assert.Equal(t, nil, p.Error())

	car := Car{
		Brand: "SAAB",
		Model: "9-5",
		Year:  2001,
	}

	prj := m.Projector(&test)

	t.ResetTimer()

	for i := 0; i < t.N; i++ {
		prj.Project(&car)
	}

}
