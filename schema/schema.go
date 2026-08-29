// Package schema embeds the JSON Schema document generated from the
// sas.Architecture Go type. The Go structs are the source of truth; this
// schema is a generated, embeddable artifact for non-Go consumers and for
// validating hand-authored architecture documents.
package schema

import _ "embed"

//go:generate go run gen/main.go

//go:embed architecture.schema.json
var ArchitectureJSON []byte
