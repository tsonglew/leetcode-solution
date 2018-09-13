package main

import (
	"log"
	"testing"
)

type Case struct {
	a string
	b int
}

func TestMain(t *testing.T) {
	cases := []Case{
		Case{"Hello World", 5},
		Case{"a ", 1},
	}
	for i := range cases {
		c := cases[i]
		if r := lengthOfLastWord(c.a); r != c.b {
			log.Fatalf("%v got %v, should get %v", c.a, r, c.b)
		}
	}
}
