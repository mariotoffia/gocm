package cment

import "github.com/mariotoffia/gocm/cmid"

// EntityFactory is a implementation that is able to create
// entities bases on their _PK_ and _SK_.
//
// Multiple `EntityFactory` instance may be at use at the same time.
type EntityFactory interface {
	// Create uses the _id_ parameter to determine the type
	// of entity to be created if the `EntityFactory` supports it.
	//
	// If the entity could not be created, `nil` is returned.
	Create(id cmid.ComponentIdentity) interface{}
	// Identities is an array of all types that the factory may
	// be able to create.
	Identities() []cmid.ComponentIdentity
}

// EntityFactoryImpl is a default implementation of a `EntityFactory` for code
// based factory methods.
type EntityFactoryImpl struct {
	// Identities is an array of all types that the factory may
	// be able to create.
	ID []cmid.ComponentIdentity
	// createEntity is a function exposed as `EntityFactory.Create()`
	createEntity func(id cmid.ComponentIdentity) interface{}
}

// Create uses the _id_ parameter to determine the type
// of entity to be created if the `EntityFactory` supports it.
//
// If the entity could not be created, `nil` is returned.
func (efi *EntityFactoryImpl) Create(id cmid.ComponentIdentity) interface{} {
	return efi.createEntity(id)
}

// Identities is an array of all types that the factory may
// be able to create.
func (efi *EntityFactoryImpl) Identities() []cmid.ComponentIdentity {
	return efi.ID
}

// NewFactory creates a entity factory for a one or more scalar or composite id. Since one entity factory
// may handle multiple entities it uses a vardict id.
func NewFactory(create func(id cmid.ComponentIdentity) interface{}, id ...cmid.ComponentIdentity) *EntityFactoryImpl {

	return &EntityFactoryImpl{
		ID:           id,
		createEntity: create,
	}

}
