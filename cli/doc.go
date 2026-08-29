// Package cli implements the business logic behind the sas command-line
// tool: loading architecture documents and running library packages
// (validate, and later render/diff/bind) against them, plus formatting
// results for console or JSON output. cmd/sas is a thin Cobra adapter
// over this package — flag parsing and process exit codes live there;
// everything else is here and independently testable.
package cli
