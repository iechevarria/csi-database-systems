package csidb

type Node interface {
	Next() bool
	Execute() Row
}

type Row struct {
	Entries []Entry
}

type Entry struct {
	Column string
	Value  string
}
