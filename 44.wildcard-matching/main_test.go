package main

import "testing"

func TestIsMatch(t *testing.T) {
	examples := []struct {
		a, b string
		s    bool
	}{
		{"aa", "a", false},
		{"a", "*", true},
		{"cb", "?a", false},
		{"adceb", "*a*b", true},
		{"acdcb", "a*c?b", false},
	}
	for i := range examples {
		a := examples[i].a
		b := examples[i].b
		s := examples[i].s
		r := isMatch(a, b)
		if s != r {
			t.Errorf("%v matches %v should be %v, but get %v", a, b, s, r)
		}
	}
}
