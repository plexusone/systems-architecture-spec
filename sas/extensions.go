package sas

// ExternalRefType names the kind of external artifact an ExternalRef
// points at.
type ExternalRefType string

const (
	ExternalRefTypeOpenAPI         ExternalRefType = "openapi"
	ExternalRefTypePIDL            ExternalRefType = "pidl"
	ExternalRefTypeThreatModelSpec ExternalRefType = "threat-model-spec"
	ExternalRefTypeMultiAgentSpec  ExternalRefType = "multi-agent-spec"
	ExternalRefTypeTerraform       ExternalRefType = "terraform"
	ExternalRefTypeOTel            ExternalRefType = "otel"
)

// ExternalRef points from a node or relationship to detail owned by
// another specification, so SAS can reference OpenAPI operations, PIDL
// protocol definitions, threat models, agent definitions, and IaC without
// duplicating their content.
type ExternalRef struct {
	// Type identifies which external specification Ref points into.
	// Required.
	Type ExternalRefType `json:"type"`

	// Ref is the locator within that specification, e.g. an OpenAPI
	// operationId, a PIDL protocol URI, or a Terraform resource address.
	// Required.
	Ref string `json:"ref"`
}

// SecurityExtension carries security-profile facts about a node or
// relationship that do not belong in the core graph shape.
type SecurityExtension struct {
	// TrustZone names the security trust zone the element belongs to,
	// e.g. "restricted" or "application".
	TrustZone string `json:"trustZone,omitempty"`
}

// SREExtension carries site-reliability facts about a node.
type SREExtension struct {
	// AvailabilityTarget is the declared SLO, e.g. "99.99%".
	AvailabilityTarget string `json:"availabilityTarget,omitempty"`
}

// ComplianceExtension marks a node or relationship's involvement in a
// compliance scope. Boundary-level control mappings live on
// Boundary.Compliance; this is for marking individual elements.
type ComplianceExtension struct {
	// FedRAMPBoundary marks that this element falls within the FedRAMP
	// authorization boundary.
	FedRAMPBoundary bool `json:"fedrampBoundary,omitempty"`
}

// AgentExtension marks a node as an agent runtime for the agentic-systems
// use case, without pulling agent governance concepts into the core graph.
type AgentExtension struct {
	// Runtime marks that this node is (or hosts) an agent runtime.
	Runtime bool `json:"runtime,omitempty"`
}

// Extensions carries namespaced, typed extension data on a node,
// relationship, or boundary. Each namespace is an explicit, statically
// typed struct field — never a generic map — so the static-type profile
// holds even as new concerns (security, SRE, compliance, agent, and
// others as they arise) attach data to the core graph. Adding a new
// extension namespace means adding a new named field and type here, not
// widening a dynamic bag.
type Extensions struct {
	Security   *SecurityExtension   `json:"security,omitempty"`
	SRE        *SREExtension        `json:"sre,omitempty"`
	Compliance *ComplianceExtension `json:"compliance,omitempty"`
	Agent      *AgentExtension      `json:"agent,omitempty"`
}
