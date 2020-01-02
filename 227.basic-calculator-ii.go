import (
	"unicode"
)

func calculate(s string) int {
	nums := []int{}
	ops := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		if unicode.IsDigit(rune(s[i])) {
			num, ni := findNextNumber(s, i)
			nums = append(nums, num)
			i=ni
			continue
		}
		switch s[i] {
		case '+', '-':
			ops = append(ops, s[i])
		case '*':
			n1 := nums[len(nums)-1]
			n2, ni := findNextNumber(s, i+1)
			nums = nums[:len(nums)-1]
			nums = append(nums, n1*n2)
			i = ni
		case '/':
			n1 := nums[len(nums)-1]
			n2, ni := findNextNumber(s, i+1)
			nums = nums[:len(nums)-1]
			nums = append(nums, n1/n2)
			i = ni
		}
	}
	for i := 0; i < len(ops); i++ {
		switch ops[i] {
		case '-':
			nums[1] = nums[0] - nums[1]
		case '+':
			nums[1] = nums[0] + nums[1]
		}
		nums = nums[1:]
	}
	return nums[0]
}

func findNextNumber(s string, i int) (int, int) {
	buf := 0 
	for ;i<len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		n := s[i]
		if unicode.IsDigit(rune(n)) {
			buf = buf * 10 + int(n-'0')
		} else {
			return buf, i-1
		}
	}
	return buf, len(s)
}
