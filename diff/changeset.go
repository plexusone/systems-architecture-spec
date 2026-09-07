package diff

// ElementType names which part of the graph a Change applies to.
type ElementType string

const (
	ElementTypeNode         ElementType = "node"
	ElementTypeRelationship ElementType = "relationship"
	ElementTypeBoundary     ElementType = "boundary"
)

// ChangeKind classifies what happened to an element between the base and
// proposed architecture.
type ChangeKind string

const (
	ChangeKindNodeAdded    ChangeKind = "node_added"
	ChangeKindNodeRemoved  ChangeKind = "node_removed"
	ChangeKindNodeModified ChangeKind = "node_modified"

	// ChangeKindBoundaryMembershipChanged is reported separately from
	// ChangeKindNodeModified even though membership lives on Node, since
	// "this node now belongs to a new boundary" is the specific fact a
	// change-impact classifier (new boundary crossing) needs to find
	// without parsing a generic field diff.
	ChangeKindBoundaryMembershipChanged ChangeKind = "boundary_membership_changed"

	ChangeKindRelationshipAdded    ChangeKind = "relationship_added"
	ChangeKindRelationshipRemoved  ChangeKind = "relationship_removed"
	ChangeKindRelationshipModified ChangeKind = "relationship_modified"

	ChangeKindBoundaryAdded    ChangeKind = "boundary_added"
	ChangeKindBoundaryRemoved  ChangeKind = "boundary_removed"
	ChangeKindBoundaryModified ChangeKind = "boundary_modified"
)

// FieldDelta names a single field that differs between the base and
// proposed element, with human-readable before/after values. Field is a
// dotted path (e.g. "transport.encryption") so callers can filter or
// match on the specific fact that changed rather than parsing prose.
// Before/After are omitted (empty string) when the field was unset on
// that side.
type FieldDelta struct {
	Field  string `json:"field"`
	Before string `json:"before,omitempty"`
	After  string `json:"after,omitempty"`
}

// Change is a single typed fact about how the architecture changed: one
// element, one kind of change, and (for modifications) the specific
// fields that differ.
type Change struct {
	Kind        ChangeKind   `json:"kind"`
	ElementType ElementType  `json:"elementType"`
	ElementID   string       `json:"elementId"`
	Deltas      []FieldDelta `json:"deltas,omitempty"`
}

// ChangeSet is every Change between a base and proposed Architecture, in
// a stable order (nodes, then relationships, then boundaries; within each,
// sorted by element ID).
type ChangeSet struct {
	Changes []Change `json:"changes"`
}
