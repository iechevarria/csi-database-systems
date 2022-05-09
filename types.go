package csidb

import (
	"fmt"
)

type Node interface {
	Next() bool
	Execute() Row
}

type Row struct {
	Entries []Entry
}

func (r Row) Get(column string) Entry {
	for _, e := range r.Entries {
		if e.Column == column {
			return e
		}
	}
	panic(fmt.Sprintf("No column \"%s\"", column))
}

type Entry struct {
	Column string
	Value  string
}
