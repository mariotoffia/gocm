package identity

// IDMapper maps a entity into a `Identity` or vice versa.
type IDMapper interface {
	// AssembleIdentity will create a new identity instance based of the _v_ parameter.
	//
	// This function do not alter the state of the _v_ parameter, opposed of `Map`
	// function.
	AssembleIdentity(v interface{}) (Identity, error)
	// ExtractIdentity will lookup the PK and SK in the _v_ parameter and create a
	// identity from that. This is the opposite from `AssembleIdentity`.
	ExtractIdentity(v interface{}) (Identity, error)
}

// IDObjectMapper maps or resolves a entity identity.
//
// This is a extension on the `IDMapper` interface.
type IDObjectMapper interface {
	// IDMapper is the base interface for this interface.
	IDMapper
	// Map will map the parameter _v_ PartitionKey and SortKey in same instance.
	//
	// It expects a pointer to the instance. This function may not be idempotent
	// since it alters the state inline.
	Map(v interface{}) error
	// Resolve will resolve any fields that are part of the PK and SK and set those from
	// the PK and SK attributes from the provided instance.
	//
	// This function is never idempotent since it will change the PK or SK. The other fields
	// is overwritten. But the PK and SK is always changed (if not static value).
	Resolve(v interface{}) error
}
