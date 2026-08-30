// Package assure computes assurance-reference coverage over a
// sas.Architecture: which nodes and relationships declare tests,
// metrics, detections, and deployment evidence, and which don't. It
// answers "which declared elements have no proof they exist, work,
// are observable, or are monitored" — a coverage question over
// references the architecture already carries, not a reconciliation
// against live evidence (that is out of scope for v0.1; see
// sas.Assurance).
package assure
