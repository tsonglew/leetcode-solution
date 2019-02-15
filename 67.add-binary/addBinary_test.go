package main

import "testing"

func TestAddBinary(t *testing.T) {
	tables := []struct {
		a, b, op string
	}{
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
		{"", "", ""},
		{"1", "", "1"},
		{"", "0", "0"},
	}

	for _, table := range tables {
		if op := addBinary(table.a, table.b); op != table.op {
			t.Errorf("addBinary(%s, %s) should return %s, but got %s",
				table.a, table.b, table.op, op)
		}
	}
}
