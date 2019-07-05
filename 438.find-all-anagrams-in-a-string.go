package main

import (
	"fmt"
)

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
}

func findAnagrams(s string, p string) []int {
    charCnt := make(map[byte]int)
    charTotal := make(map[byte]int)
    charPosition := make(map[byte][]int)
	result := []int{}
    for i := range p {
        if _, ok := charTotal[p[i]]; ok {
            charTotal[p[i]]++
        } else {
            charTotal[p[i]] = 1
        }
        charCnt[p[i]] = 0
        charPosition[p[i]] = []int{}
    }
      
    for head, tail := 0, 0; tail < len(s); tail++{
        if _, ok := charCnt[s[tail]]; !ok {
            for ;head < tail; head++ {
				charCnt[s[head]]--
                charPosition[s[head]] = charPosition[s[head]][1:]
            }
            head++
            continue
        }
        charCnt[s[tail]]++
        charPosition[s[tail]] = append(charPosition[s[tail]], tail)
        if charCnt[s[tail]] > charTotal[s[tail]] {
			formerPosition := charPosition[s[tail]][0]
            for ;head <= formerPosition; head++ {
                charCnt[s[head]]--
                charPosition[s[head]] = charPosition[s[head]][1:]
			}
			continue
        }
        if tail-head+1 == len(p) {
            result = append(result, head)
            charPosition[s[head]] = charPosition[s[head]][1:]
            charCnt[s[head]]--
            head++
        }
    }
    return result
}