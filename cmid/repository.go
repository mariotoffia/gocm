package cmid

// IDMapperRepository is providing functionality to register entities
// and create mappers for those.
//
// The repository itself may not use tags or reflection to provide
// its functionality.
type IDMapperRepository interface {
	// Error returns any error state
	Error() error
	// ClearError will clear any error state
	ClearError() IDMapperRepository
	// Mapper gets the mapper for the parameter _v_ (as it was registered using `Add()`).
	//
	// Some implementations has default implementations and thus do not need to do `Add()`.
	//
	// If none is registered for type _v_, `nil` is returned.
	Mapper(v any) IDObjectMapper
	// Add creates a new `IdMapper` and stores it in its internal hash.
	//
	// The parameter _v_ is expected to be a pointer to a type. If the repository is `IsFrozen()`
	// it will silently ignore the `Add()` request.
	Add(v any) IDMapperRepository
	// Mappers returns an array of currently supported mappings.
	Mappers() []IDObjectMapper
	// Freeze will make registration of new mappings impossible.
	Freeze() IDMapperRepository
	// IsFrozen returns `true` if the instance do not accept any more mapping regisrations using `Add()`.
	IsFrozen() bool
}
