package main

import (
	"strings"
)

type Interval struct {
	Start int
	End   int
}

func lengthOfLastWord(s string) int {
	wordSlice := strings.Split(strings.Trim(s, " "), " ")
	if len(wordSlice) < 1 {
		return 0
	}
	return len(wordSlice[len(wordSlice)-1])
}

func main() {
}
