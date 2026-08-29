// Package render holds logic shared by every format-specific renderer
// (render/mermaid, render/d2, render/dot): grouping selected nodes by
// boundary membership and choosing a node shape and edge label from the
// semantic model. Each renderer is a pure function (Architecture, View)
// -> string; sharing this logic guarantees the different output formats
// agree on what a view means, and differ only in syntax.
package render
