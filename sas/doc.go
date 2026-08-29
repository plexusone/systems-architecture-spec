// Package sas defines the canonical Systems Architecture Spec (SAS) graph
// model: the typed nodes, relationships, boundaries, identities, and
// entitlements that make up a machine-readable description of a software
// system.
//
// The Go types in this package are the source of truth. JSON Schema is
// generated from them (package schema) rather than hand-written, and the
// generated schema is validated to stay within the PlexusOne static-type
// profile: no oneOf/anyOf/allOf, no schema composition, no untyped maps.
// Every type here must have a direct, unambiguous JSON representation.
//
// Semantic understanding — what a node is, what an edge means, which
// boundaries it crosses, which identity it acts as — is this package's
// responsibility. Diagram rendering, threat-model export, and change
// analysis are separate packages that consume this model; they do not
// define it.
package sas
