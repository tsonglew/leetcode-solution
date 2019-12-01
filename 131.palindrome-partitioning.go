package main

import "fmt"

func partition(s string) [][]string {
	current := []string{}
	result := [][]string{}
	dfs(s, 0, current, &result)
	fmt.Println(result)
	return result
}

func dfs(s string, pos int, current []string, result *[][]string) {
	if pos >= len(s) {
		*result = append(*result, append([]string{}, current...))
		return
	}
	for i := pos; i < len(s); i++ {
		if isPalindrome(s[pos : i+1]) {
			current = append(current, s[pos:i+1])
			dfs(s, i+1, current, result)
			current = current[:len(current)-1]
		}
	}
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	partition("aab")
	partition("aabbcbb")
}
