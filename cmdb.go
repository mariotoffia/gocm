package gocm

import "github.com/mariotoffia/gocm/cmmgr"

// EntityManagerAware is implemented when it is possible to externally set or
// interact with a `cmmgr.EntityManager`.
type EntityManagerAware interface {
	// Manager gets the `cmmgr.EntityManager`.
	//
	// If none is set, it shall return a default `cmmgr.EntityManager`.
	Manager() cmmgr.EntityManager
	// SetManager sets a manager if none has been set before.
	//
	// If a `Manager()` call has been made, it is not possible to set a new manager
	// and the function will return an `error`. This function is optional, and if
	// it is not supported an error is returned.
	SetManager(mgr cmmgr.EntityManager) error
}

// CMDb is the main functionality to communicate with one or more CM repositories.
// If the `CMDb` implementation suports a `EntityManager` it will in addition to this
// implement the `EntityManagerAware` interface.
//
// Since each type of implementation differs, they often cannot share entity manager. However,
// all is _REQUIRED_ to interpret the standard _cm_ based `Identity` and the convention around
// the _PK_ and _SK_ format.
type CMDb interface {
	// Repositories returns the underlying repositories in this instance.
	//
	// This is often used when a aggregate `CMDb` implementation do delegate to others to perform
	// its work.
	Repositories() []CMDb
}
