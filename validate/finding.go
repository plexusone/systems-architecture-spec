package validate

// Severity classifies how serious a Finding is.
type Severity string

const (
	SeverityError   Severity = "error"
	SeverityWarning Severity = "warning"
)

// Finding is a single validation result: something Validate checked,
// where it applies, and whether it passed. Only failures are returned by
// Validate — a clean run produces an empty slice.
type Finding struct {
	// RuleID identifies the rule that produced this finding, e.g.
	// "core.relationship-endpoints-resolve".
	RuleID string `json:"ruleId"`

	// Severity classifies how serious this finding is.
	Severity Severity `json:"severity"`

	// Profile names which profile required this check. Empty means the
	// rule is always-on (core referential integrity), independent of any
	// requested profile.
	Profile Profile `json:"profile,omitempty"`

	// Message is a human-readable explanation.
	Message string `json:"message"`

	// Path locates the offending element, e.g. "nodes[2]" or
	// "relationships[0]".
	Path string `json:"path"`
}
