package validate

// Profile names a use case that requires additional semantics beyond the
// core graph shape. Requesting a profile does not change what an
// Architecture document is allowed to contain — the schema is the same
// for every profile — it changes which optional fields become mandatory
// for Validate to consider the document conformant.
type Profile string

const (
	// ProfileDevelopment is the base profile: nodes and relationships
	// only. It adds no rules beyond the always-on referential-integrity
	// checks — a minimal architecture file is already conformant to it.
	ProfileDevelopment Profile = "development"

	// ProfileDeployment requires deployment-relevant facts: node
	// technology, membership in an account/region/network boundary, and
	// relationship transport.
	ProfileDeployment Profile = "deployment"

	// ProfileSecurity requires identity and encryption on boundary
	// crossings, and entitlements on relationships that write.
	ProfileSecurity Profile = "security"

	// ProfileThreatModel requires data classifications on relationships
	// touching a node outside every boundary the source node belongs to
	// (an external dependency), so a threat model always knows what data
	// leaves a trust zone.
	ProfileThreatModel Profile = "threat-model"

	// ProfileSRE requires owner, an SRE extension with an availability
	// target, and assurance metrics on tier0/tier1 nodes.
	ProfileSRE Profile = "sre"
)

// KnownProfiles lists every profile Validate recognizes, in a stable
// order suitable for CLI help text and error messages.
func KnownProfiles() []Profile {
	return []Profile{
		ProfileDevelopment,
		ProfileDeployment,
		ProfileSecurity,
		ProfileThreatModel,
		ProfileSRE,
	}
}

// IsKnown reports whether p is one of KnownProfiles.
func (p Profile) IsKnown() bool {
	for _, known := range KnownProfiles() {
		if p == known {
			return true
		}
	}
	return false
}
