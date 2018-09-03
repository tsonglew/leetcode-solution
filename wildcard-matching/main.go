package main

import (
	"fmt"
	"strings"
)

func isMatch(s string, p string) bool {
	if len(s) <= 0 {
		if len(strings.Trim(p, "*")) <= 0 {
			return true
		}
		return false
	}
	matches := make([][]bool, len(p)+1)
	for i := 0; i < len(p)+1; i++ {
		matches[i] = make([]bool, len(s)+1)
	}
	matches[0][0] = true
	for i := 1; i < len(s)+1; i++ {
		matches[0][i] = false
	}
	for i := 1; i < len(p)+1; i++ {
		if matches[i-1][0] == true && p[i-1] == '*' {
			matches[i][0] = true
		} else {
			matches[i][0] = false
		}
	}
	for i := 1; i < len(p)+1; i++ {
		for j := 1; j < len(s)+1; j++ {
			if matches[i-1][j-1] == true {
				if p[i-1] == s[j-1] || p[i-1] == '*' || p[i-1] == '?' {
					matches[i][j] = true
					continue
				}
			} else if matches[i-1][j] == true || matches[i][j-1] == true {
				if p[i-1] == '*' {
					matches[i][j] = true
					continue
				}
			}
			matches[i][j] = false
		}
	}
	return matches[len(p)][len(s)]
}

func main() {
	fmt.Println(isMatch("aa", "a"))
}
