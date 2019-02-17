package main

func main() {

}

func mySqrt(x int) int {
	var min, max, mid int
	min = 1
	max = x/2 + 1
	for min <= max {
		mid = (min + max) / 2
		if (mid+1)*(mid+1) <= x {
			min = mid + 1
		} else if mid*mid > x {
			max = mid - 1
		} else if mid*mid <= x && (mid+1)*(mid+1) > x {
			break
		}
	}
	return mid
}
