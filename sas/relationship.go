package sas

// RelationKind classifies what a Relationship represents.
type RelationKind string

const (
	RelationKindCalls           RelationKind = "calls"
	RelationKindDataAccess      RelationKind = "data.access"
	RelationKindDataFlow        RelationKind = "data_flow"
	RelationKindPublishesTo     RelationKind = "publishes_to"
	RelationKindSubscribesTo    RelationKind = "subscribes_to"
	RelationKindAuthenticatesTo RelationKind = "authenticates_to"
	RelationKindAuthorizesVia   RelationKind = "authorizes_via"
	RelationKindObserves        RelationKind = "observes"
	RelationKindEnforces        RelationKind = "enforces"
	RelationKindDependsOn       RelationKind = "depends_on"
	RelationKindSpawns          RelationKind = "spawns"
	RelationKindDelegatesTo     RelationKind = "delegates_to"
)

// Operation is the generic, protocol-independent verb vocabulary for what
// a relationship does to its target. Concrete protocol operations (HTTP
// GET, SQL SELECT, MCP invoke_tool) specialize one of these verbs via
// catalog-supplied mappings, so a security view can show "read" while an
// implementation view shows "HTTPS GET /v1/orders" from the same
// underlying fact.
type Operation string

const (
	OperationDiscover   Operation = "discover"
	OperationRead       Operation = "read"
	OperationCreate     Operation = "create"
	OperationUpdate     Operation = "update"
	OperationDelete     Operation = "delete"
	OperationExecute    Operation = "execute"
	OperationAdminister Operation = "administer"
)

// SyncMode names whether a relationship is a synchronous call or an
// asynchronous interaction.
type SyncMode string

const (
	SyncModeSync  SyncMode = "sync"
	SyncModeAsync SyncMode = "async"
)

// Transport carries the wire-level facts of a relationship: what protocol
// and port it uses and whether it is encrypted. This is what most diagram
// formats collapse into an unlabeled arrow.
type Transport struct {
	// Protocol is the transport-layer protocol, e.g. "tcp", "udp".
	Protocol string `json:"protocol,omitempty"`

	// Port is the network port, when applicable.
	Port int `json:"port,omitempty"`

	// ApplicationProtocol is the application-layer protocol, e.g.
	// "https", "postgresql", "grpc".
	ApplicationProtocol string `json:"applicationProtocol,omitempty"`

	// Encryption names the encryption in use, e.g. "tls", "mtls", or
	// "none". Required on boundary-crossing edges by the security
	// profile.
	Encryption string `json:"encryption,omitempty"`
}

// DataFlow describes what data moves across a relationship.
type DataFlow struct {
	// Classifications lists the kinds of data carried, e.g.
	// "customer_data", "source_code", "pii". Required on edges touching
	// external nodes by the threat-model profile.
	Classifications []string `json:"classifications,omitempty"`
}

// Relationship is a directed edge between two nodes. It is deliberately
// the richest type in the core model: what most diagrams draw as an
// unlabeled arrow, SAS requires to be able to carry protocol, port,
// encryption, the specific operations performed, which identity the
// caller acts as, the entitlements exercised, what data crosses the wire,
// and whether the edge crosses a trust boundary — the facts a security or
// threat-model review actually needs.
type Relationship struct {
	// ID is the stable identity of this relationship.
	ID string `json:"id"`

	// From is the source node ID. Required.
	From string `json:"from"`

	// To is the target node ID. Required.
	To string `json:"to"`

	// Kind classifies what this relationship represents. Required.
	Kind RelationKind `json:"kind"`

	// Transport carries protocol/port/encryption facts.
	Transport *Transport `json:"transport,omitempty"`

	// Operations lists the specific verbs this relationship performs
	// against its target, from the generic operation vocabulary.
	// Required on write operations by the security profile
	// (authorization.entitlements must then also be present).
	Operations []Operation `json:"operations,omitempty"`

	// Identity names which identity the caller acts as when exercising
	// this relationship.
	Identity *IdentityRef `json:"identity,omitempty"`

	// Authorization carries the entitlements exercised by this
	// relationship.
	Authorization *Authorization `json:"authorization,omitempty"`

	// Data describes what data moves across this relationship.
	Data *DataFlow `json:"data,omitempty"`

	// CrossesBoundaries lists the IDs of boundaries this relationship
	// crosses. An edge crosses a boundary when exactly one endpoint is a
	// member of it; this field lets an author assert crossings
	// explicitly, and a validator can additionally derive and cross-check
	// it from Node.Boundaries membership.
	CrossesBoundaries []string `json:"crossesBoundaries,omitempty"`

	// CriticalPath marks that this relationship sits on a critical
	// execution path.
	CriticalPath bool `json:"criticalPath,omitempty"`

	// Sync names whether this relationship is synchronous or
	// asynchronous.
	Sync SyncMode `json:"sync,omitempty"`

	// Assurance references evidence that this relationship exists and
	// works as declared.
	Assurance *Assurance `json:"assurance,omitempty"`

	// ProtocolRef points at a PIDL protocol definition this relationship
	// participates in, e.g. "pidl://oauth/dpop-authorization-code". SAS
	// never duplicates protocol choreography — see ProtocolBinding.
	ProtocolRef string `json:"protocolRef,omitempty"`

	// Extensions carries namespaced, typed data for concerns that do not
	// belong in the core shape.
	Extensions *Extensions `json:"extensions,omitempty"`
}
