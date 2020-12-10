package cmid

const (
	// IDStandardDivider is the standard component, id divider _#_.
	IDStandardDivider string = "#"
	// IDStandardCMTag is the standard tag that can be used to decorate
	// the struct fields (when using the idreflect package.)
	IDStandardCMTag string = "cm"
)

// Identity represents a single identity for a instance.
type Identity interface {
	// PartitionKey is the key that zooms into a certain partition or bucket.
	PartitionKey() string
	// SecondaryKey or SortKey is the key to uniquely identify the enity
	// within a partition. If no specific entity this key is empty.
	//
	// It is also possible to have a partial SecondaryKey when e.g. data is hierarchical
	// and thus points to a certain node in the hierarchy.
	SecondaryKey() string
}

// ComponentIdentity specifies that the _PK_ and _SK_ is only containing the
// component parts and not a complete instance identity.
//
// .Example ComponentIdentity
// [source,json]
// ....
// { "PK": "P", "SK": "S"}
// ....
//
// The above example denotes that the complete identity is _P#S_ and it may for example
// resolve to a certain go instance type in the `EntityFactory`.
type ComponentIdentity interface {
	Identity
}

// ID is a `Identity` implementation.
type ID struct {
	PK string
	SK string
}

// PartitionKey is the key that zooms into a certain partition or bucket.
func (id *ID) PartitionKey() string {
	return id.PK
}

// SecondaryKey or SortKey is the key to uniquely identify the enity
// within a partition. If no specific entity this key is empty.
//
// It is also possible to have a partial SecondaryKey when e.g. data is hierarchical
// and thus points to a certain node in the hierarchy.
func (id *ID) SecondaryKey() string {
	return id.SK
}
