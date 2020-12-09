package identity

// IDMediaMapper maps a `Identity` to a medium and vice versa.
//
// Some implementations allows for custom `IMediaMapper` and some have a built-in
// fixed mapper. All implementations must natively translate a cm `Identity`.
//
// For example a s3 mapper could be _1.0.0/{pk}/{year}/{model}_ and an example key could then be
// _1.0.0/SAAB/2001/9-5.json_. The postfix is determined by the implementation media type (in
// this case _JSON_).
type IDMediaMapper interface {
}
