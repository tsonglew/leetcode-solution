package main

import (
	"reflect"
	"testing"
)

func TestMakeSpaces(t *testing.T) {
	tables := []struct {
		num    int
		output string
	}{
		{1, " "},
		{0, ""},
		{2, "  "},
	}
	for _, table := range tables {
		if op := makeSpaces(table.num); op != table.output {
			t.Errorf("makeSpaces(%d) should return '%v', but got '%v'", table.num, table.output, op)
		}
	}
}

func TestJustifyLine(t *testing.T) {
	tables := []struct {
		words    []string
		wordLens []int
		last     bool
		maxWidth int
		output   string
	}{
		{
			[]string{"This", "is", "an"},
			[]int{4, 2, 2},
			false,
			16,
			"This    is    an",
		},
		{
			[]string{"example", "of", "text"},
			[]int{7, 2, 4},
			false,
			16,
			"example  of text",
		},
		{
			[]string{"justification."},
			[]int{14},
			true,
			16,
			"justification.  ",
		},
		{
			[]string{"acknowledgment"},
			[]int{14},
			false,
			16,
			"acknowledgment  ",
		},
		{
			[]string{"shall", "be"},
			[]int{5, 2},
			true,
			16,
			"shall be        ",
		},
	}
	for _, table := range tables {
		if op := justifyLine(table.words, table.wordLens, table.last, table.maxWidth); op != table.output {
			t.Errorf("justifyLine(%v, %v, %v, %v) should return '%v', but got '%v'",
				table.words, table.wordLens, table.last, table.maxWidth, table.output, op)
		}
	}
}

func TestFullJustify(t *testing.T) {
	tables := []struct {
		words    []string
		maxWidth int
		output   []string
	}{
		{
			[]string{"This", "is", "an", "example", "of", "text", "justification."},
			16,
			[]string{"This    is    an", "example  of text", "justification.  "},
		},
		{
			[]string{"What", "must", "be", "acknowledgment", "shall", "be"},
			16,
			[]string{"What   must   be", "acknowledgment  ", "shall be        "},
		},
		{
			[]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain",
				"to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			20,
			[]string{"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  "},
		},
	}
	for _, table := range tables {
		if out := fullJustify(table.words, table.maxWidth); !reflect.DeepEqual(out, table.output) {
			t.Errorf("fullJustify(%v, %d) should return %v, but got %v", table.words, table.maxWidth, table.output, out)
		}
	}
}
