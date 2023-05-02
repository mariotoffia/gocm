package cmexpr

// AllAttributes is the wildcard for all attributes and cannot have a alias.
const AllAttributes = "*"

type Projection interface {
	Expression
	Attributes() []ProjectionPair
}

type ProjectionPair struct {
	Alias     string
	Attribute string
}

func (pp *ProjectionPair) IsWildCard() bool {
	return pp.Attribute == "*"
}

type ProjectionImpl struct {
	ExpressionImpl
	attributes []ProjectionPair
}

func (p *ProjectionImpl) Attributes() []ProjectionPair {
	return p.attributes
}

func (p *ProjectionImpl) Attribute(name string, alias ...string) *ProjectionImpl {

	pp := ProjectionPair{
		Attribute: name,
	}

	if len(alias) > 0 {

		if name == AllAttributes {
			panic("cannot alias wildcard attribute")
		}

		pp.Alias = alias[0]
	}

	p.attributes = append(p.attributes, pp)

	return p
}
