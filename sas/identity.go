package sas

// IdentityType classifies who or what an Identity represents.
type IdentityType string

const (
	IdentityTypeHuman    IdentityType = "human"
	IdentityTypeWorkload IdentityType = "workload"
	IdentityTypeService  IdentityType = "service"
	IdentityTypeAgent    IdentityType = "agent"
	IdentityTypeDevice   IdentityType = "device"
	IdentityTypeExternal IdentityType = "external"
)

// IdentityMechanism names the concrete authentication/authorization
// mechanism backing an Identity. This is metadata for humans and tooling;
// SAS does not implement or validate the mechanism itself.
type IdentityMechanism string

const (
	IdentityMechanismAWSIAMRole        IdentityMechanism = "aws_iam_role"
	IdentityMechanismSPIFFE            IdentityMechanism = "spiffe"
	IdentityMechanismOAuthClient       IdentityMechanism = "oauth_client"
	IdentityMechanismOIDC              IdentityMechanism = "oidc"
	IdentityMechanismAAuth             IdentityMechanism = "aauth"
	IdentityMechanismK8sServiceAccount IdentityMechanism = "k8s_service_account"
	IdentityMechanismX509              IdentityMechanism = "x509"
	IdentityMechanismAPIKey            IdentityMechanism = "api_key"
)

// Identity describes who or what acts as a node, or which identity a
// relationship's caller assumes. Identity is first-class rather than an
// attribute so security and entitlement analysis can rely on it being
// present and typed.
type Identity struct {
	// Type classifies the identity holder. Required.
	Type IdentityType `json:"type"`

	// Mechanism names the concrete authentication mechanism, e.g. "spiffe"
	// or "aws_iam_role". Optional — a design-time model may know the type
	// of identity before the mechanism is chosen.
	Mechanism IdentityMechanism `json:"mechanism,omitempty"`

	// Ref points at the concrete identity, e.g. a SPIFFE ID, an IAM role
	// ARN, or an OAuth client ID. Optional.
	Ref string `json:"ref,omitempty"`
}

// IdentityRef identifies which identity a relationship's caller acts as,
// by reference to a node's declared Identity or to a standalone identity
// description. Kept distinct from Identity so an edge can name "the
// identity node X uses" without repeating the full Identity payload.
type IdentityRef struct {
	// NodeID references the node whose Identity applies. Mutually
	// exclusive with Identity in practice, but both are optional so a
	// relationship can reference a node's identity or describe one inline.
	NodeID string `json:"nodeId,omitempty"`

	// Identity describes the identity inline when it is not simply "the
	// identity of the source node."
	Identity *Identity `json:"identity,omitempty"`
}

// Entitlement is a subject-action-resource authorization tuple, shaped to
// map cleanly onto ReBAC systems such as OpenFGA or SpiceDB (subject→user,
// action→relation, resource→object) without making SAS itself a ReBAC
// engine.
type Entitlement struct {
	// Subject is the identity or identity class the entitlement applies
	// to, e.g. a node ID or an identity ref. Required.
	Subject string `json:"subject"`

	// Action is the operation the subject is entitled to perform, e.g.
	// "read" or "orders.insert". Required.
	Action string `json:"action"`

	// Resource is what the action applies to, e.g. a node ID or a
	// resource locator. Required.
	Resource string `json:"resource"`

	// Conditions are human/agent-readable qualifiers narrowing the grant,
	// e.g. "production only". SAS does not evaluate conditions; they are
	// recorded for downstream policy engines and reviewers.
	Conditions []string `json:"conditions,omitempty"`
}

// Authorization carries the entitlements associated with a relationship,
// e.g. what a caller may do once it has crossed a boundary.
type Authorization struct {
	// Entitlements lists the specific grants exercised by this
	// relationship.
	Entitlements []Entitlement `json:"entitlements,omitempty"`
}
