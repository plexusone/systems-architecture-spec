// Package validate applies profile-conditional rules to a sas.Architecture.
// A profile declares which optional semantics in the core graph become
// mandatory for a given use case (development, deployment, security,
// threat-model, sre); the core schema itself never branches on profile.
//
// This package is a consumer of the sas package's semantic model, not
// part of it: which rules exist and what they require is a use-case
// requirement layered on top of the graph, not a fact about the graph
// itself.
package validate
