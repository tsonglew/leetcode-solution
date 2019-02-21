package main

import (
	"reflect"
	"testing"
)

func TestSetZeroes(t *testing.T) {
	tables := []struct {
		input  [][]int
		output [][]int
	}{
		{
			[][]int{[]int{1, 1, 1}, []int{1, 0, 1}, []int{1, 1, 1}},
			[][]int{[]int{1, 0, 1}, []int{0, 0, 0}, []int{1, 0, 1}},
		},
		{
			[][]int{[]int{0, 1, 2, 0}, []int{3, 4, 5, 2}, []int{1, 3, 1, 5}},
			[][]int{[]int{0, 0, 0, 0}, []int{0, 4, 5, 0}, []int{0, 3, 1, 0}},
		},
		{
			[][]int{[]int{1, 1, 1}, []int{0, 1, 2}},
			[][]int{[]int{0, 1, 1}, []int{0, 0, 0}},
		},
	}
	for i := range tables {
		if setZeroes(tables[i].input); !reflect.DeepEqual(tables[i].input, tables[i].output) {
			t.Errorf("setZeroes should return %v, but got %v",
				tables[i].output, tables[i].input)
		}
	}
}
