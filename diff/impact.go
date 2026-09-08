package diff

// ImpactKind classifies the security/compliance-relevant meaning of a
// Change — the semantic fact a human reviewer or a compliance profile
// actually cares about, as opposed to ChangeKind's structural fact (what
// was added/removed/modified).
type ImpactKind string

const (
	// ImpactKindNewExternalDependency: a new relationship to (or a new
	// node that is) an external_service.
	ImpactKindNewExternalDependency ImpactKind = "new_external_dependency"

	// ImpactKindAuthnMechanismChanged: the identity a relationship's
	// caller acts as changed.
	ImpactKindAuthnMechanismChanged ImpactKind = "authn_mechanism_changed"

	// ImpactKindEntitlementExpansion: a relationship gained a write
	// operation (create/update/delete/administer) it did not have
	// before, or gained new authorization entitlements.
	ImpactKindEntitlementExpansion ImpactKind = "entitlement_expansion"

	// ImpactKindNewBoundaryCrossing: a relationship now crosses a
	// boundary it did not cross before (including a brand new
	// relationship that crosses one from the start).
	ImpactKindNewBoundaryCrossing ImpactKind = "new_boundary_crossing"

	// ImpactKindDataClassificationChange: the data classifications
	// carried by a relationship changed.
	ImpactKindDataClassificationChange ImpactKind = "data_classification_change"

	// ImpactKindComputeModelChanged: a node's underlying Technology
	// changed, e.g. compute.instance moving from ecs to lambda.
	ImpactKindComputeModelChanged ImpactKind = "compute_model_changed"
)

// Impact is a single semantic classification of one Change, with enough
// detail for a human reviewer or compliance profile to act on without
// re-deriving it from the raw FieldDelta.
type Impact struct {
	Kind               ImpactKind  `json:"kind"`
	ElementType        ElementType `json:"elementType"`
	ElementID          string      `json:"elementId"`
	AffectedBoundaries []string    `json:"affectedBoundaries,omitempty"`
	Detail             string      `json:"detail"`
}

// ChangeImpact is every Impact derived from a ChangeSet, in the same
// stable order as the ChangeSet's own Changes.
type ChangeImpact struct {
	Impacts []Impact `json:"impacts"`
}
