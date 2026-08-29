package schema

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/santhosh-tekuri/jsonschema/v6"

	"github.com/plexusone/systems-architecture-spec/sas"
)

// TestValidFixtures proves the shared fixture corpus round-trips through
// the full Go/JSON Schema conformance chain: every file under
// examples/fixtures/valid must validate against the generated JSON
// Schema and unmarshal cleanly into sas.Architecture. This is the
// behavioral conformance contract referenced by the TypeScript/Zod side
// (ts/), which validates the same fixtures.
func TestValidFixtures(t *testing.T) {
	compiler := jsonschema.NewCompiler()

	var schemaDoc any
	if err := json.Unmarshal(ArchitectureJSON, &schemaDoc); err != nil {
		t.Fatalf("unmarshal embedded schema: %v", err)
	}
	if err := compiler.AddResource("architecture.schema.json", schemaDoc); err != nil {
		t.Fatalf("add schema resource: %v", err)
	}
	compiled, err := compiler.Compile("architecture.schema.json")
	if err != nil {
		t.Fatalf("compile schema: %v", err)
	}

	fixtures, err := filepath.Glob(filepath.Join("..", "examples", "fixtures", "valid", "*.json"))
	if err != nil {
		t.Fatalf("glob fixtures: %v", err)
	}
	if len(fixtures) == 0 {
		t.Fatal("no valid fixtures found under examples/fixtures/valid")
	}

	for _, path := range fixtures {
		path := path
		t.Run(filepath.Base(path), func(t *testing.T) {
			data, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("read fixture: %v", err)
			}

			var doc any
			if err := json.Unmarshal(data, &doc); err != nil {
				t.Fatalf("fixture is not valid JSON: %v", err)
			}
			if err := compiled.Validate(doc); err != nil {
				t.Fatalf("fixture does not validate against architecture.schema.json: %v", err)
			}

			var arch sas.Architecture
			if err := json.Unmarshal(data, &arch); err != nil {
				t.Fatalf("fixture does not unmarshal into sas.Architecture: %v", err)
			}
			if arch.Version == "" {
				t.Error("expected non-empty Version after unmarshal")
			}
			if len(arch.Nodes) == 0 {
				t.Error("expected at least one node after unmarshal")
			}
		})
	}
}
