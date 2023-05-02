package cmid

// IDMapper maps a entity into a `Identity` or vice versa.
type IDMapper interface {
	// AssembleIdentity will create a new identity instance based of the _v_ parameter.
	//
	// This function do not alter the state of the _v_ parameter. The returned `Identity`
	// is formatted as the mapper is configured and the values are extracted from the
	// _v_ parameter.
	AssembleIdentity(v any) (Identity, error)
	// ExtractIdentity will lookup the PK and SK in the _v_ parameter and create a
	// identity from that. Hence, only the PK and SK fields are used in the _v_.
	ExtractIdentity(v any) (Identity, error)
}

// IDObjectMapper maps or resolves a entity identity.
//
// This is a extension on the `IDMapper` interface.
type IDObjectMapper interface {
	// IDMapper is the base interface for this interface.
	IDMapper
	// Map will map the parameter _v_ (pointer) PartitionKey and SortKey in the instance _v_.
	//
	// This will write to the _PK_ and _SK_ (if present) with the expressions declared.
	Map(v any) error
	// Resolve will resolve any fields that are part of the PK and SK and set those from
	// the PK and SK attributes from the provided instance.
	//
	// This function is never idempotent since it will change the PK or SK. The other fields
	// is overwritten. But the PK and SK is always changed (if not static value).
	Resolve(v any) error
}
