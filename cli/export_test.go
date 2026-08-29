package cli

import (
	"strings"
	"testing"
)

const architectureForExportJSON = `{
  "version": "0.1",
  "metadata": {"name": "Example"},
  "nodes": [
    {"id": "api", "kind": "service", "name": "API"},
    {"id": "db", "kind": "data.database", "name": "DB"}
  ],
  "relationships": [
    {"id": "api-to-db", "from": "api", "to": "db", "kind": "data.access"}
  ]
}`

func TestExportThreatModel(t *testing.T) {
	path := writeTempArchitecture(t, architectureForExportJSON)

	out, err := ExportThreatModel(ExportThreatModelOptions{ArchitecturePath: path})
	if err != nil {
		t.Fatalf("ExportThreatModel: %v", err)
	}
	if !strings.Contains(out, `"type": "dfd"`) {
		t.Errorf("expected dfd type in output, got %q", out)
	}
	if !strings.Contains(out, `"id": "api"`) || !strings.Contains(out, `"id": "db"`) {
		t.Errorf("expected both nodes exported, got %q", out)
	}
}

func TestExportThreatModel_MissingFile(t *testing.T) {
	if _, err := ExportThreatModel(ExportThreatModelOptions{ArchitecturePath: "/nonexistent/architecture.json"}); err == nil {
		t.Fatal("expected error for missing file")
	}
}
