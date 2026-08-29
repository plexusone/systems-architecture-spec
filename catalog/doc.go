// Package catalog provides technology-specific display metadata and
// protocol-operation-to-generic-verb mappings, kept outside the core sas
// package on purpose: "AWS Lambda" is not a core NodeKind, it's
// Kind "compute.function" plus Technology{Provider: "aws", Service:
// "lambda"}. Catalogs translate that pair into something renderers and
// authors find useful (a display name, an icon hint) without teaching
// the core graph model about any specific vendor.
package catalog
