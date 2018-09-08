package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 103}
	m := map[int]([]string){}
	for _, str := range strs {
		mul := 1
		for _, c := range str {
			mul *= primes[c-97]
		}
		if _, ok := m[mul]; !ok {
			m[mul] = []string{}
		}
		m[mul] = append(m[mul], str)
	}
	r := make([]([]string), len(m))
	r_cnt := 0
	for _, v := range m {
		if len(v) <= 0 {
			continue
		}
		r[r_cnt] = v
		r_cnt++
	}
	return r
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
