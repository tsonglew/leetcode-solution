package main

import (
	"testing"
)

func TestMinWindow(t *testing.T) {
	tables := []struct {
		s, t   string
		output string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"ABBBC", "ABC", "ABBBC"},
		{"ABC", "ABC", "ABC"},
		{"aa", "aa", "aa"},
		{"acbdbaab", "aabd", "dbaa"},
	}
	for i := range tables {
		if op := minWindow(tables[i].s, tables[i].t); op != tables[i].output {
			t.Errorf("minWindow() should return %v, but got %v", tables[i].output, op)
		}
	}
}
