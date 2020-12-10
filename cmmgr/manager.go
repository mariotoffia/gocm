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
	// Freeze will freeze all registrations in all sub repositories as well.
	Freeze() EntityManager
	// IsFrozen returns `true` if the manager is completely frozen, including sub-repositories.
	IsFrozen() bool
}

// EntityManagerImpl is a default implementation of `EntityManager`.
type EntityManagerImpl struct {
	factories   cment.EntityFactoryRepository
	mappers     cmid.IDMapperRepository
	projections cmid.IDProjectRepository
	frozen      bool
}

// NewEntityManager creates a default entity manager where it is possible to
// register factories, mappers, and projections on.
func NewEntityManager() *EntityManagerImpl {
	return &EntityManagerImpl{}
}

// Freeze will freeze all registrations in all sub repositories as well.
func (em *EntityManagerImpl) Freeze() EntityManager {

	if em.factories != nil {
		em.factories.Freeze()
	}

	if em.mappers != nil {
		em.mappers.Freeze()
	}

	if em.projections != nil {
		em.projections.Freeze()
	}

	em.frozen = true

	return em
}

// IsFrozen returns `true` if the manager is completely frozen, including sub-repositories.
func (em *EntityManagerImpl) IsFrozen() bool {
	return em.frozen
}

// UseFactories sets the current factory
func (em *EntityManagerImpl) UseFactories(factories cment.EntityFactoryRepository) *EntityManagerImpl {
	em.factories = factories
	return em
}

// Factories exposes the repository where all `EntityFactory`
// instances reside.
//
// NOTE: When it is _frozen_ it is not possible to register more
// factory instances.
func (em *EntityManagerImpl) Factories() cment.EntityFactoryRepository {
	return em.factories
}

// UseMappers sets the current mappers
func (em *EntityManagerImpl) UseMappers(mappers cmid.IDMapperRepository) *EntityManagerImpl {
	em.mappers = mappers
	return em
}

// Mappers exposes the mapper repository.
//
// NOTE: When it is _frozen_ it is not possible to register more
// mapper instances.
func (em *EntityManagerImpl) Mappers() cmid.IDMapperRepository {
	return em.mappers
}

// UseProjections will set the current projections.
func (em *EntityManagerImpl) UseProjections(projections cmid.IDProjectRepository) *EntityManagerImpl {
	em.projections = projections
	return em
}

// Projections exposes any projections that the `CMDb` shall do when
// storing the data (if supported).
//
// NOTE: When it is _frozen_ it is not possible to register more
// `cmid.IDProjector` instances.
func (em *EntityManagerImpl) Projections() cmid.IDProjectRepository {
	return em.projections
}
