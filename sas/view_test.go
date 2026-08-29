package sas

import "testing"

func testSelectArchitecture() Architecture {
	return Architecture{
		Version: "0.1",
		Nodes: []Node{
			{ID: "web", Kind: NodeKindService, Name: "Web", Boundaries: []string{"vpc-prod"}},
			{ID: "api", Kind: NodeKindService, Name: "API", Boundaries: []string{"vpc-prod"}},
			{ID: "db", Kind: NodeKindDataDatabase, Name: "DB", Boundaries: []string{"vpc-prod"}},
			{ID: "external", Kind: NodeKindExternalService, Name: "External"},
		},
		Relationships: []Relationship{
			{ID: "web-to-api", From: "web", To: "api", Kind: RelationKindCalls},
			{ID: "api-to-db", From: "api", To: "db", Kind: RelationKindDataAccess},
			{ID: "api-to-external", From: "api", To: "external", Kind: RelationKindCalls},
		},
		Boundaries: []Boundary{
			{ID: "vpc-prod", Kind: BoundaryKindNetwork, Name: "Prod VPC"},
			{ID: "unused", Kind: BoundaryKindTrust, Name: "Unused"},
		},
	}
}

func TestSelect_IncludeKinds(t *testing.T) {
	arch := testSelectArchitecture()

	result := arch.Select(View{
		ID:           "databases",
		IncludeKinds: []string{string(NodeKindDataDatabase)},
	})

	if len(result.Nodes) != 1 || result.Nodes[0].ID != "db" {
		t.Fatalf("expected only db node, got %+v", result.Nodes)
	}
	// No relationship survives: api-to-db's From ("api") is excluded.
	if len(result.Relationships) != 0 {
		t.Fatalf("expected no relationships, got %+v", result.Relationships)
	}
}

func TestSelect_IncludeRelations(t *testing.T) {
	arch := testSelectArchitecture()

	result := arch.Select(View{
		ID:               "calls-only",
		IncludeRelations: []string{string(RelationKindCalls)},
	})

	if len(result.Relationships) != 2 {
		t.Fatalf("expected 2 calls relationships, got %d: %+v", len(result.Relationships), result.Relationships)
	}
	for _, r := range result.Relationships {
		if r.Kind != RelationKindCalls {
			t.Errorf("unexpected relationship kind %s in calls-only view", r.Kind)
		}
	}
	// All 4 nodes still present: IncludeRelations does not filter nodes.
	if len(result.Nodes) != 4 {
		t.Fatalf("expected all 4 nodes, got %d", len(result.Nodes))
	}
}

func TestSelect_BoundariesDefaultToReferenced(t *testing.T) {
	arch := testSelectArchitecture()

	result := arch.Select(View{ID: "all"})

	if len(result.Boundaries) != 1 || result.Boundaries[0].ID != "vpc-prod" {
		t.Fatalf("expected only vpc-prod (the unused boundary should be dropped), got %+v", result.Boundaries)
	}
}

func TestSelect_ExplicitIncludeBoundaries(t *testing.T) {
	arch := testSelectArchitecture()

	result := arch.Select(View{ID: "all", IncludeBoundaries: []string{"unused"}})

	if len(result.Boundaries) != 1 || result.Boundaries[0].ID != "unused" {
		t.Fatalf("expected only the explicitly included boundary, got %+v", result.Boundaries)
	}
}

func TestSelect_PreservesMetadataAndVersion(t *testing.T) {
	arch := testSelectArchitecture()
	arch.Metadata = Metadata{Name: "example"}

	result := arch.Select(View{ID: "all"})

	if result.Version != "0.1" || result.Metadata.Name != "example" {
		t.Fatalf("expected version/metadata preserved, got version=%q metadata=%+v", result.Version, result.Metadata)
	}
}
