package sas

import (
	"reflect"
	"testing"
)

func testCrossingArchitecture() Architecture {
	return Architecture{
		Version: "0.1",
		Nodes: []Node{
			{ID: "api", Kind: NodeKindService, Name: "API", Boundaries: []string{"vpc-prod", "trust-app"}},
			{ID: "db", Kind: NodeKindDataDatabase, Name: "DB", Boundaries: []string{"vpc-prod"}},
			{ID: "external", Kind: NodeKindExternalService, Name: "External"},
		},
		Relationships: []Relationship{
			{ID: "api-to-db", From: "api", To: "db", Kind: RelationKindDataAccess},
			{ID: "api-to-external", From: "api", To: "external", Kind: RelationKindCalls},
			{ID: "api-to-unknown", From: "api", To: "does-not-exist", Kind: RelationKindCalls},
		},
		Boundaries: []Boundary{
			{ID: "vpc-prod", Kind: BoundaryKindNetwork, Name: "Prod VPC"},
			{ID: "trust-app", Kind: BoundaryKindTrust, Name: "App Trust Zone"},
		},
	}
}

func TestNodeByID(t *testing.T) {
	arch := testCrossingArchitecture()

	n, ok := arch.NodeByID("api")
	if !ok || n.Name != "API" {
		t.Fatalf("expected to find node api, got %+v ok=%v", n, ok)
	}

	if _, ok := arch.NodeByID("missing"); ok {
		t.Fatal("expected missing node lookup to fail")
	}
}

func TestBoundaryByID(t *testing.T) {
	arch := testCrossingArchitecture()

	b, ok := arch.BoundaryByID("vpc-prod")
	if !ok || b.Name != "Prod VPC" {
		t.Fatalf("expected to find boundary vpc-prod, got %+v ok=%v", b, ok)
	}

	if _, ok := arch.BoundaryByID("missing"); ok {
		t.Fatal("expected missing boundary lookup to fail")
	}
}

func TestViewByID(t *testing.T) {
	arch := testCrossingArchitecture()
	arch.Views = []View{{ID: "context", Name: "Context View"}}

	v, ok := arch.ViewByID("context")
	if !ok || v.Name != "Context View" {
		t.Fatalf("expected to find view context, got %+v ok=%v", v, ok)
	}

	if _, ok := arch.ViewByID("missing"); ok {
		t.Fatal("expected missing view lookup to fail")
	}
}

func TestBindingsForProtocol(t *testing.T) {
	arch := testCrossingArchitecture()
	arch.Bindings = []ProtocolBinding{
		{ID: "b1", ProtocolRef: "pidl://oauth2-authorization-code", Participants: map[string]string{"client": "api"}},
		{ID: "b2", ProtocolRef: "pidl://oauth2-authorization-code", Participants: map[string]string{"client": "db"}},
		{ID: "b3", ProtocolRef: "pidl://mcp-tool-call", Participants: map[string]string{"client": "api"}},
	}

	matches := arch.BindingsForProtocol("pidl://oauth2-authorization-code")
	if len(matches) != 2 || matches[0].ID != "b1" || matches[1].ID != "b2" {
		t.Fatalf("expected [b1 b2], got %+v", matches)
	}

	if matches := arch.BindingsForProtocol("pidl://nonexistent"); matches != nil {
		t.Fatalf("expected no matches, got %+v", matches)
	}
}

func TestCrossedBoundaries(t *testing.T) {
	arch := testCrossingArchitecture()

	tests := []struct {
		name string
		rel  Relationship
		want []string
	}{
		{
			name: "api to db shares vpc-prod, differs on trust-app",
			rel:  arch.Relationships[0],
			want: []string{"trust-app"},
		},
		{
			name: "api to external: external is in no boundaries, so api's boundaries all cross",
			rel:  arch.Relationships[1],
			want: []string{"trust-app", "vpc-prod"},
		},
		{
			name: "unresolvable target returns nil",
			rel:  arch.Relationships[2],
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := arch.CrossedBoundaries(tt.rel)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CrossedBoundaries() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCrossedBoundaries_SameMembership(t *testing.T) {
	arch := Architecture{
		Nodes: []Node{
			{ID: "a", Boundaries: []string{"vpc-prod"}},
			{ID: "b", Boundaries: []string{"vpc-prod"}},
		},
		Relationships: []Relationship{
			{ID: "a-to-b", From: "a", To: "b"},
		},
	}
	got := arch.CrossedBoundaries(arch.Relationships[0])
	if got != nil {
		t.Errorf("expected no crossings for identical membership, got %v", got)
	}
}
