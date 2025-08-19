package user

import _ "embed"

// Spec contains the OpenAPI v2 specification for User.
//
//go:embed docs/swagger.json
var Spec []byte
