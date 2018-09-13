package main

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	wordSlice := strings.Split(strings.Trim(s, " "), " ")
	if len(wordSlice) < 1 {
		return 0
	}
	return len(wordSlice[len(wordSlice)-1])
}

func main() {
}
