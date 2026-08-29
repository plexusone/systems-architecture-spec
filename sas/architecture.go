package sas

// Metadata carries descriptive facts about an Architecture document that
// are not part of the graph itself.
type Metadata struct {
	// Name is the human-readable name of the system being described.
	Name string `json:"name,omitempty"`

	// Description explains what the system is and does.
	Description string `json:"description,omitempty"`

	// Owner identifies who is responsible for this architecture document.
	Owner string `json:"owner,omitempty"`
}

// Architecture is the canonical, top-level SAS document: a typed graph of
// Nodes and Boundaries (and, as later RMIs land, Relationships, protocol
// Bindings, and Views). It is the semantic model — the thing this project
// exists to define. Diagrams, threat models, and change reports are all
// computed from an Architecture; none of them are the Architecture.
type Architecture struct {
	// Version is the SAS document schema version, e.g. "0.1".
	Version string `json:"version"`

	// Metadata carries descriptive facts about the system.
	Metadata Metadata `json:"metadata,omitempty"`

	// Nodes lists every system, service, component, and actor in the
	// architecture.
	Nodes []Node `json:"nodes"`

	// Relationships lists every directed edge between nodes.
	Relationships []Relationship `json:"relationships,omitempty"`

	// Boundaries lists every network, trust, environment, account,
	// region, compliance, and organization boundary. Nodes reference
	// boundaries by ID; membership is many-to-many.
	Boundaries []Boundary `json:"boundaries,omitempty"`
}
