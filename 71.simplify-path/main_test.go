package main

import (
	"testing"
)

func TestSimplyfyPath(t *testing.T) {
	tables := []struct {
		in, out string
	}{
		{"/home/", "/home"},
		{"/../", "/"},
		{"/home//foo/", "/home/foo"},
		{"/a/./b/../../c/", "/c"},
		{"/a/../../b/../c//.//", "/c"},
		{"/a//b////c/d//././/..", "/a/b/c"},
		{"/...", "/..."},
	}
	for _, table := range tables {
		if op := simplifyPath(table.in); op != table.out {
			t.Errorf("simplifyPath(%v) should return %v, but got %v",
				table.in, table.out, op)
		}
	}
}
