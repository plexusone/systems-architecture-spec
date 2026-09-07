// Package diff computes a typed ChangeSet between two versions of a
// sas.Architecture: which nodes, relationships, and boundaries were
// added, removed, or modified, and which specific fields changed on each.
//
// Diff is a pure structural comparison. It classifies no impact and makes
// no judgment about significance — turning a ChangeSet into a semantic
// classification (new external dependency, authentication mechanism
// change, entitlement expansion, ...) is a separate, later concern layered
// on top of this package.
package diff
