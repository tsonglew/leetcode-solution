package main

import "testing"

func TestIsMatch(t *testing.T) {
	examples := []struct {
		a []int
		b int
	}{
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{2, 3, 0, 1, 4}, 2},
		{[]int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}, 3},
	}
	for i := range examples {
		a := examples[i].a
		b := examples[i].b
		r := jump(a)
		if b != r {
			t.Errorf("%v needs %v jumps, but get %v", a, b, r)
		}
	}
}
