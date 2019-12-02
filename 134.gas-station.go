package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	if len(gas) == 0 {
		return -1
	}
	origin := make([]int, len(gas))
	sum := 0
	for i := range gas {
		origin[i] = gas[i] - cost[i]
		sum += origin[i]
	}
	if sum < 0 {
		return -1
	}
	for i := 0; i < len(gas); i++ {
		if origin[i] < 0 {
			continue
		}
		tmp := origin[i]
		p := (i + 1) % len(gas)
		for p != i {
			tmp += origin[p]
			if tmp < 0 {
				break
			}
			p = (p + 1) % len(gas)
		}
		if p == i {
			return i
		}
		i = p
	}
	return -1
}

func main() {
	fmt.Println(canCompleteCircuit([]int{3, 1, 1}, []int{1, 2, 2}))
}
