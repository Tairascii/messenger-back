package chats

import _ "embed"

// Spec contains the OpenAPI v2 specification for Chats.
//
//go:embed docs/swagger.json
var Spec []byte
