package main

import (
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	tables := []struct {
		matrix [][]int
		target int
		output bool
	}{
		{
			[][]int{
				[]int{1, 3, 5, 7},
				[]int{10, 11, 16, 20},
				[]int{23, 30, 34, 50},
			},
			3,
			true,
		},
		{
			[][]int{
				[]int{1, 3, 5, 7},
				[]int{10, 11, 16, 20},
				[]int{23, 30, 34, 50},
			},
			13,
			false,
		},
	}
	for i := range tables {
		if op := searchMatrix(tables[i].matrix, tables[i].target); op != tables[i].output {
			t.Errorf("searchMatrix() should return %v, but got %v",
				tables[i].output, op)
		}
	}
}
