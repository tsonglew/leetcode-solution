package main

func main() {}

func uniquePaths(m int, n int) int {
	downSteps := n - 1
	rightSteps := m - 1
	return combination(downSteps+rightSteps, mn(downSteps, rightSteps))
}

func mn(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func combination(m, n int) int {
	a, b := 1, 1
	for n > 0 {
		a *= m
		b *= n
		m--
		n--
	}
	return a / b
}
