package main

import "testing"

func TestMultiply(t *testing.T) {
	examples := []struct {
		a, b, s string
	}{
		{"2", "3", "6"},
		{"123", "456", "56088"},
	}
	for i := range examples {
		a := examples[i].a
		b := examples[i].b
		s := examples[i].s
		r := multiply(a, b)
		if s != r {
			t.Errorf("%v multiply %v should be %v, but get %v", a, b, s, r)
		}
	}
}
