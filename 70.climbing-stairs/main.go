package main

func main() {

}

func climbStairs(n int) int {
	if n < 3 {
		steps := []int{0, 1, 2, 3}
		return steps[n]
	}
	steps := make([]int, n+1)
	steps[0], steps[1], steps[2] = 0, 1, 2
	for i := 3; i <= n; i++ {
		steps[i] = steps[i-1] + steps[i-2]
	}
	return steps[n]
}
