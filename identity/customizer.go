package identity

// IDCustomizer customizes `Identity` or object with a certain _cm_ identity.
//
// Some implementations allows for custom `IDCustomizer`. All implementations must
// natively translate a cm `Identity` mapping. If support for customization, this allows
// for a application to customize to deviate from standard cm mapping.
//
// For example a s3 customizer could be _1.0.0/{pk}/{year}/{model}_ and an example key could
// then be _1.0.0/SAAB/2001/9-5.json_. The postfix is determined by the implementation media
// type (in this case _JSON_).
type IDCustomizer interface {
}
