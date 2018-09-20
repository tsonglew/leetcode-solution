package main

func main() {

}

func plusOne(digits []int) []int {
	var sum int
	plusone := 1
	i := len(digits) - 1
	for plusone == 1 && i >=0 {
		sum = digits[i] + 1
		plusone = sum / 10
		digits[i] = sum % 10
		i--
	}
	if plusone == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}