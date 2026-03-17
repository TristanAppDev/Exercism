package tree

import (
	"fmt"
	"sort"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

// Nodes slice of records.
type Nodes []*Node

// Len returns the length of the slice.
func (r Nodes) Len() int {
	return len(r)
}

// Swap swaps two records.
func (r Nodes) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

// Less compares two records by their IDs and their parents IDs.
// If the records have the same ID and Parent, the first record is less than the second.
func (r Nodes) Less(i, j int) bool {
	return r[i].ID < r[j].ID
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	// every node will have id equal to its index in the slice nodes
	nodes := make([]Node, len(records))
	parents := make([]*Node, len(records))
	// to check if a node has a duplicate
	passed := make([]bool, len(records))

	for _, r := range records {
		if r.ID >= len(records) {
			return nil, fmt.Errorf("invalid id: %d", r.ID)
		}
		// return error if current record is not root but has lower/equal id than its parent
		if r.ID != 0 && r.ID <= r.Parent {
			return nil, fmt.Errorf("invalid parent id: %d", r.Parent)
		}
		// return error if id is already used
		if passed[r.ID] {
			return nil, fmt.Errorf("duplicate record id: %d", r.ID)
		}

		passed[r.ID] = true
		if r.ID != 0 {
			parents[r.ID] = &nodes[r.Parent]
		} else if r.Parent != 0 {
			return nil, fmt.Errorf("invalid parent id: %d", r.Parent)
		}
	}

	for i := 1; i < len(nodes); i++ {
		parents[i].Children = append(parents[i].Children, &nodes[i])
	}

	for i, n := range nodes {
		nodes[i].ID = i
		sort.Sort(Nodes(n.Children))
	}

	return &nodes[0], nil
}
