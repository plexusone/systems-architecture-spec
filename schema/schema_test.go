package schema

import (
	"encoding/json"
	"testing"
)

func TestArchitectureJSON_IsValidJSONSchema(t *testing.T) {
	if len(ArchitectureJSON) == 0 {
		t.Fatal("ArchitectureJSON is empty; run `go generate ./schema/...`")
	}
	var doc map[string]any
	if err := json.Unmarshal(ArchitectureJSON, &doc); err != nil {
		t.Fatalf("ArchitectureJSON is not valid JSON: %v", err)
	}
	if doc["$schema"] == nil {
		t.Error("ArchitectureJSON missing $schema")
	}
	required, _ := doc["required"].([]any)
	if !containsString(required, "version") || !containsString(required, "nodes") {
		t.Errorf("ArchitectureJSON required = %v, want version and nodes", required)
	}
}

func containsString(items []any, want string) bool {
	for _, item := range items {
		if s, ok := item.(string); ok && s == want {
			return true
		}
	}
	return false
}
