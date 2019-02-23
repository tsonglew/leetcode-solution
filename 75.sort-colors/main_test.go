package main

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	tables := []struct {
		nums   []int
		output []int
	}{
		{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{[]int{1, 2}, []int{1, 2}},
	}
	for i := range tables {
		if sortColors(tables[i].nums); !reflect.DeepEqual(tables[i].nums, tables[i].output) {
			t.Errorf("sortColors() should return %v, but got %v",
				tables[i].output, tables[i].nums)
		}
	}
}
