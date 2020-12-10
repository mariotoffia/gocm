package cmmgr

import (
	"github.com/mariotoffia/gocm/cment"
	"github.com/mariotoffia/gocm/cmid"
)

// EntityManager handles factories, mapping, and projections
// for the `CMDb`.
//
// One single entity manager may be associated with a single `CMDb` instance
// and acts as a way for the `CMDb` to create new entities when marshal data
// as well as identity mapping and resolving.
//
// For those `CMDb` implementations that do support projections uses the `Projection()`
// repository to get `IDProjector` to project data paths (or keys) onto the media.
type EntityManager interface {
	// Factories exposes the repository where all `EntityFactory`
	// instances reside.
	//
	// NOTE: When it is _frozen_ it is not possible to register more
	// factory instances.
	Factories() cment.EntityFactoryRepository
	// Mappers exposes the mapper repository.
	//
	// NOTE: When it is _frozen_ it is not possible to register more
	// mapper instances.
	Mappers() cmid.IDMapperRepository
	// Projections exposes any projections that the `CMDb` shall do when
	// storing the data (if supported).
	//
	// NOTE: When it is _frozen_ it is not possible to register more
	// `cmid.IDProjector` instances.
	Projections() cmid.IDProjectRepository
}
