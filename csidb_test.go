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
	ln := NewLimitNode(sn, 5)
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
	ln := NewLimitNode(pn, 1)
	query := NewQuery(ln)
	rows := query.Execute()

	if len(rows[0].Entries) != 1 {
		t.Fatalf("Wrong number of entries: %v", len(rows[0].Entries))
	}
}

func TestSort(t *testing.T) {
	scan, err := NewScanNode("movies")
	check(t, err)
	limit := NewLimitNode(scan, 5)
	sort := NewSortNode(limit, "title", true)
	query := NewQuery(sort)
	rows := query.Execute()
	if rows[0].Get("title").Value != "Father of the Bride Part II (1995)" {
		t.Fatalf("Wrong sorting")
	}

	scan, err = NewScanNode("movies")
	check(t, err)
	sort = NewSortNode(scan, "title", true)
	limit = NewLimitNode(sort, 10)
	query = NewQuery(limit)
	rows = query.Execute()
	if rows[0].Get("title").Value != "\"Great Performances\" Cats (1998)" {
		t.Fatalf("Wrong sorting")
	}
}
