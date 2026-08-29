package sas

// ViewLevel names the abstraction level a view renders operations at:
// the generic verb vocabulary, or the protocol-specific detail an author
// supplied (e.g. "HTTPS GET /v1/orders" instead of "read").
type ViewLevel string

const (
	ViewLevelGeneric  ViewLevel = "generic"
	ViewLevelSpecific ViewLevel = "specific"
)

// View is a query over an Architecture, not a diagram file: it selects
// which nodes, relationships, and boundaries participate in a projection
// and how to group them, so one canonical model can produce development,
// deployment, security, and threat-model projections without semantic
// drift between separately maintained diagrams. Rendering (Mermaid, D2,
// DOT) is a later concern that consumes a View's selection; this package
// only defines what a view includes.
type View struct {
	// ID is the stable identity of this view.
	ID string `json:"id"`

	// Name is the human-readable name, e.g. "Security View".
	Name string `json:"name"`

	// IncludeKinds restricts the view to nodes of these NodeKind values.
	// Empty means include all kinds.
	IncludeKinds []string `json:"includeKinds,omitempty"`

	// IncludeRelations restricts the view to relationships of these
	// RelationKind values. Empty means include all kinds. A relationship
	// is only included if both its endpoints survive the node filter.
	IncludeRelations []string `json:"includeRelations,omitempty"`

	// IncludeBoundaries restricts the view to these boundary IDs. Empty
	// means include all boundaries referenced by the selected nodes.
	IncludeBoundaries []string `json:"includeBoundaries,omitempty"`

	// GroupBy names a BoundaryKind to nest selected nodes by when
	// rendering, e.g. "network". Rendering-only; Select does not use it.
	GroupBy string `json:"groupBy,omitempty"`

	// Level chooses the operations abstraction: generic verbs or the
	// protocol-specific detail an author supplied. Rendering-only;
	// Select does not use it.
	Level ViewLevel `json:"level,omitempty"`
}

// Select applies a View's inclusion filters to an Architecture and
// returns the resulting sub-architecture: the nodes matching
// IncludeKinds, the relationships matching IncludeRelations whose
// endpoints both survived the node filter, and the boundaries matching
// IncludeBoundaries (or every boundary referenced by a selected node,
// when IncludeBoundaries is empty). GroupBy and Level are carried through
// on the result's Metadata-free selection; they are rendering hints
// consumed downstream, not selection criteria.
func (a *Architecture) Select(v View) Architecture {
	nodeKinds := toStringSet(v.IncludeKinds)
	nodes := make([]Node, 0, len(a.Nodes))
	nodeIDs := make(map[string]bool, len(a.Nodes))
	for _, n := range a.Nodes {
		if len(nodeKinds) > 0 && !nodeKinds[string(n.Kind)] {
			continue
		}
		nodes = append(nodes, n)
		nodeIDs[n.ID] = true
	}

	relKinds := toStringSet(v.IncludeRelations)
	relationships := make([]Relationship, 0, len(a.Relationships))
	for _, r := range a.Relationships {
		if !nodeIDs[r.From] || !nodeIDs[r.To] {
			continue
		}
		if len(relKinds) > 0 && !relKinds[string(r.Kind)] {
			continue
		}
		relationships = append(relationships, r)
	}

	var boundaries []Boundary
	if len(v.IncludeBoundaries) > 0 {
		want := toStringSet(v.IncludeBoundaries)
		for _, b := range a.Boundaries {
			if want[b.ID] {
				boundaries = append(boundaries, b)
			}
		}
	} else {
		referenced := make(map[string]bool)
		for _, n := range nodes {
			for _, id := range n.Boundaries {
				referenced[id] = true
			}
		}
		for _, b := range a.Boundaries {
			if referenced[b.ID] {
				boundaries = append(boundaries, b)
			}
		}
	}

	return Architecture{
		Version:       a.Version,
		Metadata:      a.Metadata,
		Nodes:         nodes,
		Relationships: relationships,
		Boundaries:    boundaries,
	}
}

func toStringSet(items []string) map[string]bool {
	if len(items) == 0 {
		return nil
	}
	set := make(map[string]bool, len(items))
	for _, item := range items {
		set[item] = true
	}
	return set
}
