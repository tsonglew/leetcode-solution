package main

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tables := []struct {
		a, op []string
	}{
		{[]string{"1", "2", "3"}, []string{"3", "2", "1"}},
		{[]string{"1", "2"}, []string{"2", "1"}},
		{[]string{"1"}, []string{"1"}},
	}

	for _, table := range tables {
		if op := reverse(table.a); !reflect.DeepEqual(op, table.op) {
			t.Errorf("reverse(%s) should return %s, but got %s",
				table.a, table.op, op)
		}
	}
}

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
