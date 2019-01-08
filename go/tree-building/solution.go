package tree

import (
	"errors"
	"sort"
)

type sortableRecords []Record

func (r sortableRecords) Len() int {
	return len(r)
}

func (r sortableRecords) Less(i, j int) bool {
	return r[i].ID < r[j].ID
}

func (r sortableRecords) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

// Record is an entity stored in a DB
type Record struct {
	ID     int
	Parent int
}

// Node is an entity of a tree
type Node struct {
	ID       int
	Children []*Node
}

// Build returns the reconstructed tree
func Build(records sortableRecords) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.Sort(records)

	if records[0].ID != 0 {
		return nil, errors.New("no root node")
	}
	if records[0].Parent != 0 {
		return nil, errors.New("root not has not parent")
	}

	nodes := make(map[int]*Node)
	nodes[0] = &Node{ID: 0}

	for idx, record := range records[1:] {
		if idx+1 != record.ID {
			return nil, errors.New("non-continuous")
		}
		if record.ID == record.Parent {
			return nil, errors.New("cycle directly")
		}
		nodes[record.ID] = &Node{ID: record.ID}
		parent, ok := nodes[record.Parent]
		if !ok {
			return nil, errors.New("no such parent")
		}
		parent.Children = append(parent.Children, nodes[record.ID])
	}

	return nodes[0], nil
}
