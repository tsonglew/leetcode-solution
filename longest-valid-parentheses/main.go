package main

import (
	"fmt"
)

func main() {
	s := ")()())()()("
	fmt.Println(longestValidParentheses(s))
}

func longestValidParentheses(s string) int {
	sLen := len(s)
	if sLen == 0 {
		return sLen
	}
	tmp := make([]int, sLen)

	for i := range s {
		if i == 0 || s[i] == '(' {
			continue
		}
		// s[i] == ')'
		if s[i-1] == '(' {
			if i >= 2 {
				tmp[i] = tmp[i-2] + 2
			} else {
				tmp[i] = 2
			}
		} else { // s[i-1] == ')'
			idxBeforePart := i - tmp[i-1] - 1
			if idxBeforePart < 0 {
				continue
			}
			if s[idxBeforePart] == '(' {
				tmp[i] = tmp[i-1] + 2
				if i-tmp[i] >= 0 {
					tmp[i] += tmp[i-tmp[i]]
				}
			}
		}
	}
	fmt.Println(tmp)

	maxVal := 0
	for _, v := range tmp {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}
