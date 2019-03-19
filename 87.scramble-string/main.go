package main

import "fmt"

func main() {
	fmt.Println(isScramble("great", "rgeat"))
	fmt.Println(isScramble("abab", "aabb"))
	fmt.Println(isScramble("abcde", "caebd"))
	fmt.Println(isScramble("cbbcacababcbc", "bacbabbbcccca"))
}

func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if len(s1) == 0 || s1 == s2 {
		return true
	}
	m := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		m[s1[i]]++
		m[s2[i]]--
	}
	for k := range m {
		if m[k] != 0 {
			return false
		}
	}
	for i := 0; i < len(s1)-1; i++ {
		if isScramble(s1[:i+1], s2[:i+1]) && isScramble(s1[i+1:], s2[i+1:]) ||
			isScramble(s1[:i+1], s2[len(s1)-1-i:]) && isScramble(s1[i+1:], s2[:len(s1)-1-i]) {
			return true
		}
	}
	return false
}
