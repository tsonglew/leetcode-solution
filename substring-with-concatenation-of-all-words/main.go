package main

import (
	"fmt"
)

func main() {
	s := "barfoofoobarthefoobarman"
	words := []string{"foo", "bar", "the"}
	fmt.Println(findSubstring(s, words))
}

func findSubstring(s string, words []string) []int {
	sLen := len(s)
	wSize := len(words)
	wLen := len(words[0])
	res := []int{}
	if sLen < 1 || wSize < 1 || sLen < wLen {
		return res
	}

	idxs := make(map[string]int)
	for _, v := range words {
		idxs[v]++
	}
	for i := 0; i < sLen-wSize*wLen+1; i++ {
		seen := make(map[string]int)
		wCnt := 0
		for j := i; j < sLen-wLen+1; j += wLen {
			subStr := s[j : j+wLen]
			if _, ok := idxs[subStr]; ok {
				seen[subStr]++
			} else {
				break
			}
			if seen[subStr] > idxs[subStr] {
				break
			}
			wCnt++
		}
		if wCnt == wSize {
			res = append(res, i)
		}
	}
	return res
}
