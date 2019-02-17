package main

import (
	"testing"
)

func TestSqrtx(t *testing.T) {
	tables := []struct {
		in, out int
	}{
		{1, 1},
		{4, 2},
		{8, 2},
	}
	for _, table := range tables {
		if op := mySqrt(table.in); op != table.out {
			t.Errorf("mySqrt(%v) should return '%v', but got '%v'",
				table.in, table.out, op)
		}
	}
}
