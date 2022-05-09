package csidb

import (
	"sort"
)

type SortNode struct {
	Parent    Node
	Column    string
	Ascending bool
	Buffer    []Row
	Empty     bool
}

func (n *SortNode) Next() bool {
	return !n.Empty
}

func (n *SortNode) Execute() Row {
	if !n.Empty && len(n.Buffer) == 0 {
		n.Init()
	}

	r := n.Buffer[0]
	n.Buffer = n.Buffer[1:]
	if len(n.Buffer) == 0 {
		n.Empty = true
	}
	return r
}

func (n *SortNode) Init() {
	// slurp the whole table into buffer and sort it
	for n.Parent.Next() {
		n.Buffer = append(n.Buffer, n.Parent.Execute())
	}
	if len(n.Buffer) == 0 {
		n.Empty = true
		return
	}

	if n.Ascending {
		sort.Slice(n.Buffer, func(i, j int) bool { return n.Buffer[i].Get(n.Column).Value < n.Buffer[j].Get(n.Column).Value })
	} else {
		sort.Slice(n.Buffer, func(i, j int) bool { return n.Buffer[i].Get(n.Column).Value > n.Buffer[j].Get(n.Column).Value })
	}
}

func NewSortNode(parent Node, column string, ascending bool) *SortNode {
	var buf []Row
	return &SortNode{
		Parent:    parent,
		Column:    column,
		Ascending: ascending,
		Buffer:    buf,
		Empty:     false,
	}
}
