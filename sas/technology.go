package sas

// Technology names the concrete vendor product implementing a node,
// separately from its core Kind. "AWS Lambda" is never a core NodeKind —
// it is Kind "compute.function" plus Technology{Provider: "aws", Service:
// "lambda"}. This keeps generic graph analysis (find all databases, find
// externally exposed compute) working across clouds while still letting
// renderers show familiar vendor icons via the catalog packages.
type Technology struct {
	// Provider names the vendor or platform, e.g. "aws", "gcp", "k8s".
	Provider string `json:"provider"`

	// Service names the specific product within that provider, e.g.
	// "lambda", "cloud-run", "deployment".
	Service string `json:"service,omitempty"`
}
