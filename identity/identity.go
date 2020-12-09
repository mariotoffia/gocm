package identity

const (
	// IDStandardDivider is the standard component, id divider _#_.
	IDStandardDivider string = "#"
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

// ID is a `Identity`
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

// IDMapper maps a entity into a `Identity`
type IDMapper interface {
	// Map will map the parameter _v_ PartitionKey and SortKey in same instance.
	//
	// It expects a pointer to the instance. This function may not be idempotent
	// since it alters the state inline.
	Map(v interface{}) error
	// Indentity will create a new identity instance based of the _v_ parameter.
	//
	// This function do not alter the state of the _v_ parameter, opposed of `Map`
	// function.
	Identity(v interface{}) (Identity, error)
}
