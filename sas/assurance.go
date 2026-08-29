package sas

// Assurance carries evidence references for a node or relationship: where
// to find the tests, metrics, detections, and deployment records that
// prove the declared element exists as described and behaves correctly.
// SAS stores pointers to evidence, never the evidence itself, and does not
// reconcile against it in v0.1 — that is a coverage-reporting and
// (eventually) runtime-reconciliation concern layered on top of this
// model, not part of the model's semantics.
type Assurance struct {
	// Tests references test suites proving the element works, e.g.
	// "test://integration/orders-db-authz".
	Tests []string `json:"tests,omitempty"`

	// Metrics references observability signals, e.g.
	// "otel://service/orders-api".
	Metrics []string `json:"metrics,omitempty"`

	// Detections references security monitoring rules, e.g.
	// "siem://rule/DB-UNEXPECTED-WRITE".
	Detections []string `json:"detections,omitempty"`

	// Deployment references records proving the element is actually
	// deployed, e.g. "aws://lambda/orders-api".
	Deployment []string `json:"deployment,omitempty"`
}
