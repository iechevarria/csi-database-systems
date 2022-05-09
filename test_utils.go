package csidb

import (
	"reflect"
	"testing"
)

func check(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func compareSlice(t *testing.T, s []string, expected []string) {
	if !reflect.DeepEqual(s, expected) {
		t.Fatalf("Wrong values: %v. Expected: %v", s, expected)
	}
}

func compareRow(t *testing.T, r Row, expected Row) {
	if !reflect.DeepEqual(r, expected) {
		t.Fatalf("Wrong values: %v. Expected: %v", r, expected)
	}
}

func assertNext(t *testing.T, next bool) {
	if !next {
		t.Fatalf("Expected true but got %v", next)
	}
}
