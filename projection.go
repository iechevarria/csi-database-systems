package csidb

type ProjectionNode struct {
	Parent  Node
	Columns []string
}

func (n *ProjectionNode) Execute() Row {
	row := n.Parent.Execute()

	var entries []Entry
	for _, c := range n.Columns {
		for _, e := range row.Entries {
			if e.Column == c {
				entries = append(entries, Entry{c, e.Value})
			}
		}
	}

	return Row{Entries: entries}
}

func (n *ProjectionNode) Next() bool {
	return n.Parent.Next()
}

func NewProjectionNode(parent Node, columns []string) *ProjectionNode {
	return &ProjectionNode{
		Parent:  parent,
		Columns: columns,
	}
}
