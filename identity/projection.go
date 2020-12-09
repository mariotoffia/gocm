package identity

// IDProjectRepository registers and uses `IDProjector`s.
//
// Some implementations allows for custom `IDProjector`. All implementations must
// natively translate a cm `Identity` mapping. If support for customization, this allows
// for a application to customize to deviate from standard cm mapping.
//
// For example a s3 customizer could be _1.0.0/{pk}/{year}/{model}_ and an example key could
// then be _1.0.0/SAAB/2001/9-5.json_. The postfix is determined by the implementation media
// type (in this case _JSON_).
type IDProjectRepository interface {
	// AddProjection will add the type _v_ with the PK and, optionally SK, expressions that it uses
	// instead of the default cm PK, SK expressions.
	//
	// If a _v_ type is registered twice, the last one is the one that will be used. If you
	// wish to have different expressions on same type, you need to create more than one
	// `IDProjectRepository`.
	AddProjection(v interface{}, expr Identity) IDProjectRepository
}

// IDProjector is mapping a entity to a Identity using the registered
// PK, and SK projections (_see `IDProjectRepository` for more information_).
//
// If a implementation do support projection, it will accept this mapping instead
// of the default cm mapping based on `IDMapper` or `IDObjectMapper`. Since this is a
// projection it will not alter the actual instance as with `IDObjectMapper`, instead
// it offers a `Identity` compatible struct that may be used to override the cm managed
// _PK_ and _SK_.
type IDProjector interface {
	// Project accepts a instance pointer _v_ that it will produce an `Identity` with
	// the PK and SK projection expression.
	//
	// The returned `Identity` may be used by any implementation to use as _PK_ and _SK_ instead
	// of the default `IDMapper` / `IDObjectMapper`.
	//
	// The _v_ entity passed must be in it's resolved statet to be guaranteed to succeed, otherwise
	// it may not contain correct information to do the projection.
	Project(v interface{}) (*ID, error)
}
