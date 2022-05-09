package csidb

import (
	"strings"
	"testing"
)

func TestFileNotExist(t *testing.T) {
	_, err := NewScanNode("potatoes")
	if !strings.HasPrefix(err.Error(), "Could not open file") {
		t.Fatalf("Wrong error: \"%v\"", err)
	}
}

func TestNewScanNode(t *testing.T) {
	n, err := NewScanNode("movies")
	check(t, err)
	compareSlice(t, n.Columns, []string{"movieId", "title", "genres"})

	assertNext(t, n.Next())

	row := n.Execute()
	compareRow(t, row, NewRow(n.Columns, []string{"1", "Toy Story (1995)", "Adventure|Animation|Children|Comedy|Fantasy"}))
	row = n.Execute()
	compareRow(t, row, NewRow(n.Columns, []string{"2", "Jumanji (1995)", "Adventure|Children|Fantasy"}))
}

func TestLimit(t *testing.T) {
	sn, err := NewScanNode("movies")
	check(t, err)
	ln := NewLimitNode(5, sn)
	query := NewQuery(ln)
	rows := query.Execute()

	if len(rows) != 5 {
		t.Fatalf("Wrong number of rows: %v", len(rows))
	}
}

func TestProjection(t *testing.T) {
	sn, err := NewScanNode("movies")
	check(t, err)
	pn := NewProjectionNode(sn, []string{"movieId"})
	ln := NewLimitNode(1, pn)
	query := NewQuery(ln)
	rows := query.Execute()

	if len(rows[0].Entries) != 1 {
		t.Fatalf("Wrong number of entries: %v", len(rows[0].Entries))
	}
}
