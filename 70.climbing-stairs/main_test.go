package main

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	tables := []struct {
		in, out int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{6, 13},
		{7, 21},
		{8, 34},
	}
	for _, table := range tables {
		if op := climbStairs(table.in); op != table.out {
			t.Errorf("climbStairs(%v) should return %v, but got %v",
				table.in, table.out, op)
		}
	}
}
