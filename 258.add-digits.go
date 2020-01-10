func addDigits(num int) int {
    if num < 10 {
		return num
	}
	return (num-1) % 9 + 1
}
