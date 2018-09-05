package main

import "testing"

func TestIsMatch(t *testing.T) {
	examples := []struct {
		a, b []int
		s    float64
	}{
		{[]int{1, 3}, []int{2}, 2},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{1, 2}, []int{-1, 3}, 1.5},
		{[]int{1, 2}, []int{1, 2, 3}, 2.0},
		{[]int{2}, []int{1, 3, 4}, 2.5},
		{[]int{}, []int{1, 2, 3, 4}, 2.5},
		{[]int{4}, []int{1, 2, 3}, 2.5},
		{[]int{3, 4}, []int{1, 2}, 2.5},
		{[]int{3}, []int{-2, -1}, -1},
		{[]int{3}, []int{1, 2, 4, 5}, 3},
	}
	for i := range examples {
		a := examples[i].a
		b := examples[i].b
		s := examples[i].s
		if r := findMedianSortedArrays(a, b); s != r {
			t.Errorf("%v and %v should be %v, but get %v", a, b, s, r)
		} else {
			t.Logf("passed: %v and %v got %v", a, b, r)
		}
	}
}
