package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func combinationSum2(candidates []int, target int) [][]int {
	result := [][]int{}
	m := make(map[string]bool)
	sort.Ints(candidates)
	if len(candidates) == 0 {
		return result
	}
	current := []int{}
	sub(candidates, &current, target, 0, &result, &m)
	return result
}

func sub(candidates []int, current *[]int, target, i int, result *[][]int, m *map[string]bool) {
	if target == 0 {
		tmpSlice := make([]int, len(*current))
		strSlice := make([]string, len(*current))
		for _, i := range *current {
			strSlice = append(strSlice, strconv.Itoa(i))
		}
		copy(tmpSlice, *current)
		if _, ok := (*m)[strings.Join(strSlice, ",")]; !ok {
			(*m)[strings.Join(strSlice, ",")] = true
			*result = append(*result, tmpSlice)
		}
		return
	}
	if target < 0 {
		return
	}
	for j := i; j < len(candidates); j++ {
		newTarget := target - candidates[j]
		*current = append(*current, candidates[j])
		sub(candidates, current, newTarget, j+1, result, m)
		*current = (*current)[0 : len(*current)-1]
	}
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
