package csidb

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type ScanNode struct {
	Columns    []string
	Reader     *csv.Reader
	NextValues []string
	Empty      bool
}

func (n *ScanNode) Next() bool {
	return !n.Empty
}

func (n *ScanNode) Execute() Row {
	row := NewRow(n.Columns, n.NextValues)

	nextValues, err := n.Reader.Read()
	// should really be err == io.EOF but whatever lets terminate on any error lol
	if err != nil {
		n.Empty = true
	}
	n.NextValues = nextValues

	return row
}

func NewScanNode(name string) (*ScanNode, error) {
	f, err := os.Open(fmt.Sprintf("ml-20m/%s.csv", name))
	if err != nil {
		return nil, fmt.Errorf("Could not open file: %w", err)
	}
	r := csv.NewReader(bufio.NewReader(f))

	columns, err := r.Read()
	if err != nil {
		return nil, fmt.Errorf("Could not read column names: %w", err)
	}

	empty := false
	nextRow, err := r.Read()
	if err == io.EOF {
		empty = true
	} else if err != nil {
		return nil, fmt.Errorf("Could not read row: %w", err)
	}

	return &ScanNode{Columns: columns, Reader: r, NextValues: nextRow, Empty: empty}, nil
}
