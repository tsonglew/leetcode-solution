package main

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	tables := []struct {
		word1, word2 string
		output       int
	}{
		{"horse", "ros", 3},
		{"intention", "execution", 5},
		{"", "", 0},
	}
	for _, table := range tables {
		if op := minDistance(table.word1, table.word2); op != table.output {
			t.Errorf("minDistance(%v, %v) should return %v, but got %v",
				table.word1, table.word2, table.output, op)
		}
	}
}
