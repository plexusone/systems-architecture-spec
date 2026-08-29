package sas

import "sort"

// NodeByID looks up a node by ID. This is the graph's own lookup, not a
// consumer helper, because resolving IDs is a prerequisite for every
// other piece of semantic reasoning over the graph (crossing detection,
// view selection, validation).
func (a *Architecture) NodeByID(id string) (*Node, bool) {
	for i := range a.Nodes {
		if a.Nodes[i].ID == id {
			return &a.Nodes[i], true
		}
	}
	return nil, false
}

// BoundaryByID looks up a boundary by ID.
func (a *Architecture) BoundaryByID(id string) (*Boundary, bool) {
	for i := range a.Boundaries {
		if a.Boundaries[i].ID == id {
			return &a.Boundaries[i], true
		}
	}
	return nil, false
}

// ViewByID looks up a declared view by ID.
func (a *Architecture) ViewByID(id string) (*View, bool) {
	for i := range a.Views {
		if a.Views[i].ID == id {
			return &a.Views[i], true
		}
	}
	return nil, false
}

// CrossedBoundaries derives which boundaries a relationship crosses from
// its endpoints' declared membership: a relationship crosses boundary B
// when exactly one of its From/To nodes is a member of B. This is
// computed independently of Relationship.CrossesBoundaries, which lets an
// author assert crossings explicitly and lets a validator compare the two
// — derived crossings the author didn't assert, or asserted crossings the
// membership doesn't support.
//
// Returns nil if either endpoint does not resolve to a known node.
func (a *Architecture) CrossedBoundaries(rel Relationship) []string {
	from, ok := a.NodeByID(rel.From)
	if !ok {
		return nil
	}
	to, ok := a.NodeByID(rel.To)
	if !ok {
		return nil
	}

	fromSet := make(map[string]bool, len(from.Boundaries))
	for _, id := range from.Boundaries {
		fromSet[id] = true
	}
	toSet := make(map[string]bool, len(to.Boundaries))
	for _, id := range to.Boundaries {
		toSet[id] = true
	}

	var crossed []string
	seen := make(map[string]bool)
	for id := range fromSet {
		if !toSet[id] && !seen[id] {
			crossed = append(crossed, id)
			seen[id] = true
		}
	}
	for id := range toSet {
		if !fromSet[id] && !seen[id] {
			crossed = append(crossed, id)
			seen[id] = true
		}
	}
	sort.Strings(crossed)
	return crossed
}
