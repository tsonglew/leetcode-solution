package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) (result [][]int) {
	result = [][]int{}
	if len(candidates) == 0 {
		return result
	}
	sort.Ints(candidates)

	current := []int{}
	sub(candidates, current, target, 0, &result)
	return
}

func sub(candidates []int, current []int, target, i int, result *[][]int) {
	if target == 0 {
		tmp := make([]int, len(current))
		copy(tmp, current)
		*result = append(*result, tmp)
		return
	}
	if target < 0 {
		return
	}
	for j := i; j < len(candidates); j++ {
		newTarget := target - candidates[j]
		if newTarget < 0 {
			return
		}
		current = append(current, candidates[j])
		sub(candidates, current, newTarget, j, result)
		current = (current)[0 : len(current)-1]
	}
}

func main() {
	candidates := []int{2, 3, 5}
	target := 8
	fmt.Println(combinationSum(candidates, target))
}
