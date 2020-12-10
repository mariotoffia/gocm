package reflectmgr

import (
	"github.com/mariotoffia/gocm/cment"
	"github.com/mariotoffia/gocm/cmid"
	"github.com/mariotoffia/gocm/cmid/idreflect"
	"github.com/mariotoffia/gocm/cmmgr"
)

// ReflectiveEntityManagerImpl is a reflective only manager that aids in registering
// entities to do mapping, factory and projection operations.
type ReflectiveEntityManagerImpl struct {
	cmmgr.EntityManagerImpl
}

// NewReflectiveEntityManager creates a entity manager that factories, mappers and projections
// are all reflection based.
//
// When done configuring factories, mappers and projection call `Freeze()` and start using the manager.
func NewReflectiveEntityManager() *ReflectiveEntityManagerImpl {

	r := ReflectiveEntityManagerImpl{}

	r.UseMappers(idreflect.NewParser()).
		UseProjections(idreflect.NewProjectRepository()).
		UseFactories(cment.NewFactories())

	return &r

}

// Register will register the entity _v_ for mapping, factory purposes.
func (rm *ReflectiveEntityManagerImpl) Register(v interface{}) *ReflectiveEntityManagerImpl {

	factory := rm.Mappers().Add(v).Mapper(v).(cment.EntityFactory)

	rm.Factories().Register(factory)

	return rm
}

// AddProjection will add the type _v_ with the PK and, optionally SK, expressions that it uses
// instead of the default cm PK, SK expressions.
//
// If a _v_ type is registered twice, the last one is the one that will be used. If you
// wish to have different expressions on same type, you need to create more than one
// `IDProjectRepository`.
func (rm *ReflectiveEntityManagerImpl) AddProjection(v interface{}, expr cmid.Identity) *ReflectiveEntityManagerImpl {

	rm.Projections().AddProjection(v, expr)
	return rm

}
