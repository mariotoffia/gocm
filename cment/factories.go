package cment

import "github.com/mariotoffia/gocm/cmid"

// EntityFactories is a collection of `EntityFactory` instances.
type EntityFactories interface {
	// EntityFactories is a aggregate `EntityFactory`
	EntityFactory
	// Register will register one or more `EntityFactory` instances.
	//
	// NOTE: The registration order matter! The last registered factory may overwrite
	// a `cmid.ComponentIdentity` and hence take precedence when creating a entity.
	Register(ef ...EntityFactory) EntityFactories
	// Freeze will make registration of new `EntityFactory` instances impossible.
	Freeze() EntityFactories
	// IsFrozen returns `true` if the instance do not accept any more factory regisrations.
	IsFrozen() bool
}

// EntityFactoriesImpl is a default implementation of `EntityFactories`
type EntityFactoriesImpl struct {
	// ef contains all `EntityFactory` instances.
	//
	// The key is the rendered `cmid.ComponentIdentity`.
	ef map[string]EntityFactory
	// comp is all components that aggregate may handle.
	comp   []cmid.ComponentIdentity
	frozen bool
}

// NewFactories creates a new empty factories instance
func NewFactories() *EntityFactoriesImpl {

	return &EntityFactoriesImpl{
		ef:   map[string]EntityFactory{},
		comp: []cmid.ComponentIdentity{},
	}

}

// Freeze will make registration of new `EntityFactory` instances impossible.
func (ef *EntityFactoriesImpl) Freeze() EntityFactories {
	ef.frozen = true
	return ef
}

// IsFrozen returns `true` if the instance do not accept any more factory regisrations.
func (ef *EntityFactoriesImpl) IsFrozen() bool {
	return ef.frozen
}

// Register will register one or more `EntityFactory` instances.
func (ef *EntityFactoriesImpl) Register(f ...EntityFactory) EntityFactories {

	if ef.frozen {
		return ef
	}

	for _, fact := range f {

		for _, id := range fact.Identities() {

			component := renderComponents(id)

			if _, ok := ef.ef[component]; !ok {
				ef.comp = append(ef.comp, id)
			}

			ef.ef[component] = fact

		}

	}

	return ef
}

// Create uses the _id_ parameter to determine the type
// of entity to be created if the `EntityFactory` supports it.
//
// If the entity could not be created, `nil` is returned. Since this is an aggregated
// `EntityFactory` it will search for the last registered factory that do support this
// `cmid.ComponentIdentity`.
func (ef *EntityFactoriesImpl) Create(id cmid.ComponentIdentity) interface{} {

	component := renderComponents(id)

	if fact, ok := ef.ef[component]; ok {
		return fact.Create(id)
	}

	return nil
}

// Identities is an array of all types that the factory may
// be able to create.
func (ef *EntityFactoriesImpl) Identities() []cmid.ComponentIdentity {
	return ef.comp
}

// renderComponents will concatenate _PK_ and _SK_ using `cmid.IDStandardDivider`.
//
// If _SK_ is empty only _PK_ is returned.
func renderComponents(id cmid.ComponentIdentity) string {

	sk := id.SecondaryKey()
	if sk == "" {
		return id.PartitionKey()
	}

	return id.PartitionKey() + cmid.IDStandardDivider + sk
}
