package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// A local copy of the real oauth2-authorization-code entity list from
// github.com/grokify/pidl's own example (examples/oauth2_authorization_code.json),
// trimmed to only the fields Bind reads, so this test doesn't take a
// module dependency on PIDL while still exercising its real shape.
const oauth2PIDLJSON = `{
  "protocol": {
    "id": "oauth2-authorization-code",
    "name": "OAuth 2.0 Authorization Code Flow"
  },
  "entities": [
    {"id": "user", "name": "Resource Owner", "type": "user"},
    {"id": "browser", "name": "User Agent", "type": "browser"},
    {"id": "client", "name": "Client Application", "type": "client"},
    {"id": "auth_server", "name": "Authorization Server", "type": "authorization_server"},
    {"id": "resource_server", "name": "Resource Server", "type": "resource_server"}
  ]
}`

func writeTempFile(t *testing.T, name, contents string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), name)
	if err := os.WriteFile(path, []byte(contents), 0o600); err != nil {
		t.Fatalf("write temp file: %v", err)
	}
	return path
}

func TestBind_FullyBound(t *testing.T) {
	pidlPath := writeTempFile(t, "oauth2.json", oauth2PIDLJSON)
	archPath := writeTempFile(t, "arch.json", `{
    "version": "0.1",
    "nodes": [
      {"id": "web", "kind": "service", "name": "Web"},
      {"id": "api", "kind": "service", "name": "API"},
      {"id": "google-oauth", "kind": "external_service", "name": "Google OAuth"}
    ],
    "bindings": [
      {
        "id": "web-google-oauth",
        "protocolRef": "pidl://oauth2-authorization-code",
        "participants": {
          "user": "web",
          "browser": "web",
          "client": "api",
          "auth_server": "google-oauth",
          "resource_server": "google-oauth"
        }
      }
    ]
  }`)

	result, err := Bind(BindOptions{ArchitecturePath: archPath, PIDLPath: pidlPath})
	if err != nil {
		t.Fatalf("Bind: %v", err)
	}
	if result.ProtocolID != "oauth2-authorization-code" {
		t.Errorf("expected protocol ID oauth2-authorization-code, got %q", result.ProtocolID)
	}
	if len(result.Bindings) != 1 {
		t.Fatalf("expected 1 binding check, got %d", len(result.Bindings))
	}
	if !result.OK() {
		t.Errorf("expected fully bound binding to be OK, got %+v", result.Bindings[0])
	}
}

func TestBind_UnknownEntity(t *testing.T) {
	pidlPath := writeTempFile(t, "oauth2.json", oauth2PIDLJSON)
	archPath := writeTempFile(t, "arch.json", `{
    "version": "0.1",
    "nodes": [{"id": "api", "kind": "service", "name": "API"}],
    "bindings": [
      {
        "id": "binding-1",
        "protocolRef": "pidl://oauth2-authorization-code",
        "participants": {"typo_entity": "api"}
      }
    ]
  }`)

	result, err := Bind(BindOptions{ArchitecturePath: archPath, PIDLPath: pidlPath})
	if err != nil {
		t.Fatalf("Bind: %v", err)
	}
	if result.OK() {
		t.Fatal("expected unknown entity to fail binding")
	}
	if len(result.Bindings[0].UnknownEntities) != 1 || result.Bindings[0].UnknownEntities[0] != "typo_entity" {
		t.Errorf("expected UnknownEntities [typo_entity], got %+v", result.Bindings[0].UnknownEntities)
	}
}

func TestBind_UnresolvedNode(t *testing.T) {
	pidlPath := writeTempFile(t, "oauth2.json", oauth2PIDLJSON)
	archPath := writeTempFile(t, "arch.json", `{
    "version": "0.1",
    "nodes": [{"id": "api", "kind": "service", "name": "API"}],
    "bindings": [
      {
        "id": "binding-1",
        "protocolRef": "pidl://oauth2-authorization-code",
        "participants": {"client": "does-not-exist"}
      }
    ]
  }`)

	result, err := Bind(BindOptions{ArchitecturePath: archPath, PIDLPath: pidlPath})
	if err != nil {
		t.Fatalf("Bind: %v", err)
	}
	if result.OK() {
		t.Fatal("expected unresolved node to fail binding")
	}
	if len(result.Bindings[0].UnresolvedNodes) != 1 || result.Bindings[0].UnresolvedNodes[0] != "does-not-exist" {
		t.Errorf("expected UnresolvedNodes [does-not-exist], got %+v", result.Bindings[0].UnresolvedNodes)
	}
}

func TestBind_UnboundEntitiesAreInformationalNotFailures(t *testing.T) {
	pidlPath := writeTempFile(t, "oauth2.json", oauth2PIDLJSON)
	archPath := writeTempFile(t, "arch.json", `{
    "version": "0.1",
    "nodes": [{"id": "api", "kind": "service", "name": "API"}],
    "bindings": [
      {
        "id": "binding-1",
        "protocolRef": "pidl://oauth2-authorization-code",
        "participants": {"client": "api"}
      }
    ]
  }`)

	result, err := Bind(BindOptions{ArchitecturePath: archPath, PIDLPath: pidlPath})
	if err != nil {
		t.Fatalf("Bind: %v", err)
	}
	if !result.OK() {
		t.Fatalf("expected partially bound binding with no errors to still be OK, got %+v", result.Bindings[0])
	}
	if len(result.Bindings[0].UnboundEntities) != 4 {
		t.Errorf("expected 4 unbound entities (user, browser, auth_server, resource_server), got %+v", result.Bindings[0].UnboundEntities)
	}
}

func TestBind_NoMatchingBindings(t *testing.T) {
	pidlPath := writeTempFile(t, "oauth2.json", oauth2PIDLJSON)
	archPath := writeTempFile(t, "arch.json", `{"version": "0.1", "nodes": [{"id": "api", "kind": "service", "name": "API"}]}`)

	result, err := Bind(BindOptions{ArchitecturePath: archPath, PIDLPath: pidlPath})
	if err != nil {
		t.Fatalf("Bind: %v", err)
	}
	if !result.OK() {
		t.Fatal("expected vacuous result (no matching bindings) to be OK")
	}
	if len(result.Bindings) != 0 {
		t.Errorf("expected no binding checks, got %+v", result.Bindings)
	}
}

func TestBind_MissingProtocolID(t *testing.T) {
	pidlPath := writeTempFile(t, "bad.json", `{"entities": []}`)
	archPath := writeTempFile(t, "arch.json", `{"version": "0.1", "nodes": []}`)

	if _, err := Bind(BindOptions{ArchitecturePath: archPath, PIDLPath: pidlPath}); err == nil {
		t.Fatal("expected error for PIDL document with no protocol.id")
	}
}

func TestFormatBindReport_Console(t *testing.T) {
	result := BindResult{
		ProtocolID: "oauth2-authorization-code",
		Bindings: []BindingCheck{
			{BindingID: "b1", UnknownEntities: []string{"typo"}},
		},
	}
	out, err := FormatBindReport(result, "console")
	if err != nil {
		t.Fatalf("FormatBindReport: %v", err)
	}
	if !strings.Contains(out, "FAILED") || !strings.Contains(out, "typo") {
		t.Errorf("expected failure and entity name in output, got %q", out)
	}
}

func TestFormatBindReport_JSON(t *testing.T) {
	result := BindResult{ProtocolID: "oauth2-authorization-code"}
	out, err := FormatBindReport(result, "json")
	if err != nil {
		t.Fatalf("FormatBindReport: %v", err)
	}
	if !strings.Contains(out, `"protocolId": "oauth2-authorization-code"`) {
		t.Errorf("expected protocolId in JSON output, got %q", out)
	}
}
