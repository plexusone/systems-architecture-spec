package cli

import (
	"strings"
	"testing"
)

const architectureWithViewsJSON = `{
  "version": "0.1",
  "nodes": [
    {"id": "web", "kind": "service", "name": "Web", "boundaries": ["vpc"]},
    {"id": "db", "kind": "data.database", "name": "DB", "boundaries": ["vpc"]},
    {"id": "external", "kind": "external_service", "name": "External"}
  ],
  "relationships": [
    {"id": "web-to-db", "from": "web", "to": "db", "kind": "data.access"},
    {"id": "web-to-external", "from": "web", "to": "external", "kind": "calls"}
  ],
  "boundaries": [
    {"id": "vpc", "kind": "network", "name": "VPC"}
  ],
  "views": [
    {"id": "databases-only", "name": "Databases Only", "includeKinds": ["data.database"]}
  ]
}`

func TestView_NamedView(t *testing.T) {
	path := writeTempArchitecture(t, architectureWithViewsJSON)

	out, err := View(ViewOptions{ArchitecturePath: path, ViewID: "databases-only", Format: "mermaid"})
	if err != nil {
		t.Fatalf("View: %v", err)
	}
	if !strings.Contains(out, `db[("DB")]`) {
		t.Errorf("expected db node, got:\n%s", out)
	}
	if strings.Contains(out, `web[`) {
		t.Errorf("did not expect web node in databases-only view, got:\n%s", out)
	}
}

func TestView_UnknownNamedView(t *testing.T) {
	path := writeTempArchitecture(t, architectureWithViewsJSON)

	if _, err := View(ViewOptions{ArchitecturePath: path, ViewID: "bogus", Format: "mermaid"}); err == nil {
		t.Fatal("expected error for unknown view ID")
	}
}

func TestView_AdHocSelection(t *testing.T) {
	path := writeTempArchitecture(t, architectureWithViewsJSON)

	out, err := View(ViewOptions{
		ArchitecturePath: path,
		GroupBy:          "network",
		Format:           "d2",
	})
	if err != nil {
		t.Fatalf("View: %v", err)
	}
	if !strings.Contains(out, `vpc: "VPC" {`) {
		t.Errorf("expected vpc container from ad-hoc groupBy, got:\n%s", out)
	}
}

func TestView_AllThreeFormats(t *testing.T) {
	path := writeTempArchitecture(t, architectureWithViewsJSON)

	for _, format := range []string{"mermaid", "d2", "dot"} {
		out, err := View(ViewOptions{ArchitecturePath: path, Format: format})
		if err != nil {
			t.Fatalf("View(format=%s): %v", format, err)
		}
		if out == "" {
			t.Errorf("expected non-empty output for format %s", format)
		}
	}
}

func TestView_UnknownFormat(t *testing.T) {
	path := writeTempArchitecture(t, architectureWithViewsJSON)

	if _, err := View(ViewOptions{ArchitecturePath: path, Format: "svg"}); err == nil {
		t.Fatal("expected error for unknown format")
	}
}

func TestView_MissingFile(t *testing.T) {
	if _, err := View(ViewOptions{ArchitecturePath: "/nonexistent/architecture.json", Format: "mermaid"}); err == nil {
		t.Fatal("expected error for missing file")
	}
}
