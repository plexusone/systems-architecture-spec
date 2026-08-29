package sas

// ProtocolBinding instantiates a PIDL protocol's abstract participants as
// concrete nodes in this architecture. PIDL models protocols using
// abstract entities (e.g. "client", "authorization_server",
// "resource_server" for OAuth); SAS never duplicates that choreography —
// it binds those entity IDs to real system identity, so the same PIDL
// protocol definition can be instantiated into many architectures.
type ProtocolBinding struct {
	// ID is the stable identity of this binding.
	ID string `json:"id"`

	// ProtocolRef points at the PIDL protocol this binding instantiates,
	// as "pidl://<protocol-id>" where protocol-id matches the
	// referenced PIDL document's protocol.id field, e.g.
	// "pidl://oauth2-authorization-code".
	ProtocolRef string `json:"protocolRef"`

	// Participants maps a PIDL entity ID (protocol.entities[].id in the
	// referenced document) to the concrete SAS node ID that plays that
	// role in this architecture.
	Participants map[string]string `json:"participants"`
}
