package csidb

type LimitNode struct {
	Parent Node
	Limit  int // count of rows to return
	Count  int // current count of rows returned
}

func (n *LimitNode) Execute() Row {
	n.Count += 1
	return n.Parent.Execute()
}

func (n *LimitNode) Next() bool {
	return n.Parent.Next() && (n.Count < n.Limit)
}

func NewLimitNode(parent Node, limit int) *LimitNode {
	return &LimitNode{
		Parent: parent,
		Limit:  limit,
		Count:  0,
	}
}
