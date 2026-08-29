package sas

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestArchitectureRoundTrip(t *testing.T) {
	original := Architecture{
		Version: "0.1",
		Metadata: Metadata{
			Name:  "example-system",
			Owner: "platform-team",
		},
		Nodes: []Node{
			{
				ID:          "api",
				Kind:        NodeKindService,
				Name:        "API",
				Owner:       "platform-team",
				Criticality: CriticalityTier0,
				Technology:  &Technology{Provider: "aws", Service: "ecs"},
				Identity:    &Identity{Type: IdentityTypeWorkload, Mechanism: IdentityMechanismSPIFFE, Ref: "spiffe://acme/api"},
				Boundaries:  []string{"vpc-prod", "trust-application"},
				Assurance:   &Assurance{Tests: []string{"test://integration/api"}},
				Refs:        []ExternalRef{{Type: ExternalRefTypeOpenAPI, Ref: "openapi://api/v1"}},
				Extensions: &Extensions{
					Security: &SecurityExtension{TrustZone: "application"},
					SRE:      &SREExtension{AvailabilityTarget: "99.99%"},
				},
			},
			{
				ID:   "db",
				Kind: NodeKindDataDatabase,
				Name: "Orders DB",
			},
		},
		Relationships: []Relationship{
			{
				ID:   "api-to-db",
				From: "api",
				To:   "db",
				Kind: RelationKindDataAccess,
				Transport: &Transport{
					Protocol:            "tcp",
					Port:                5432,
					ApplicationProtocol: "postgresql",
					Encryption:          "tls",
				},
				Operations: []Operation{OperationRead, OperationCreate},
				Identity:   &IdentityRef{NodeID: "api"},
				Authorization: &Authorization{
					Entitlements: []Entitlement{
						{Subject: "api", Action: "orders.insert", Resource: "orders-table"},
					},
				},
				Data:              &DataFlow{Classifications: []string{"customer_data"}},
				CrossesBoundaries: []string{"vpc-prod"},
				CriticalPath:      true,
				Sync:              SyncModeSync,
			},
		},
		Boundaries: []Boundary{
			{
				ID:         "vpc-prod",
				Kind:       BoundaryKindNetwork,
				Name:       "Production VPC",
				Attributes: map[string]string{"vpc": "prod"},
			},
			{
				ID:   "trust-application",
				Kind: BoundaryKindTrust,
				Name: "Application Trust Zone",
				Compliance: &ComplianceBoundary{
					FedRAMPBoundary: true,
					Controls:        []string{"SC-7", "AC-3"},
				},
			},
		},
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var decoded Architecture
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	redata, err := json.Marshal(decoded)
	if err != nil {
		t.Fatalf("re-marshal: %v", err)
	}

	if string(data) != string(redata) {
		t.Fatalf("round-trip mismatch:\nfirst:  %s\nsecond: %s", data, redata)
	}

	if len(decoded.Nodes) != 2 {
		t.Fatalf("expected 2 nodes, got %d", len(decoded.Nodes))
	}
	if decoded.Nodes[0].Identity == nil || decoded.Nodes[0].Identity.Type != IdentityTypeWorkload {
		t.Fatalf("expected node 0 identity type workload, got %+v", decoded.Nodes[0].Identity)
	}
	if len(decoded.Boundaries) != 2 {
		t.Fatalf("expected 2 boundaries, got %d", len(decoded.Boundaries))
	}
	if len(decoded.Relationships) != 1 {
		t.Fatalf("expected 1 relationship, got %d", len(decoded.Relationships))
	}
	rel := decoded.Relationships[0]
	if rel.Transport == nil || rel.Transport.ApplicationProtocol != "postgresql" {
		t.Fatalf("expected relationship transport applicationProtocol postgresql, got %+v", rel.Transport)
	}
	if len(rel.Operations) != 2 || rel.Operations[0] != OperationRead || rel.Operations[1] != OperationCreate {
		t.Fatalf("expected operations [read create], got %v", rel.Operations)
	}
}

func TestNodeMinimal(t *testing.T) {
	// A minimal node (development-profile shape: just id/kind/name) must
	// marshal and unmarshal without requiring any of the optional fields.
	n := Node{ID: "web", Kind: NodeKindService, Name: "Web"}

	data, err := json.Marshal(n)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var decoded Node
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if !reflect.DeepEqual(decoded, n) {
		t.Fatalf("expected %+v, got %+v", n, decoded)
	}
}
