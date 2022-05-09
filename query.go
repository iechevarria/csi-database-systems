package csidb

type Query struct {
	Parent Node
}

func (q *Query) Execute() []Row {
	var rows []Row
	for q.Parent.Next() {
		rows = append(rows, q.Parent.Execute())
	}
	return rows
}

func NewQuery(parent Node) *Query {
	return &Query{Parent: parent}
}
