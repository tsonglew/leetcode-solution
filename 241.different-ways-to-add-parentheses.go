import "strconv"

func isOp(op byte) bool {
	switch op {
	case '*', '+', '-':
		return true
	}
	return false
}

func calculate(ln, rn int, op byte) int {
	switch op {
	case '*':
		return ln*rn
	case '+':
		return ln +rn
	case '-':
		return ln-rn
	}
	return ln
}

func diffWaysToCompute(input string) []int {
	result := []int{}
	withOp := false
    for i, n := range input {
		if isOp(byte(n)) {
			withOp = true
			left := diffWaysToCompute(input[:i])
			right := diffWaysToCompute(input[i+1:])
			for _, ln := range left {
				for _, rn := range right {
					result = append(result, calculate(ln, rn, byte(n)))
				}
			}
		}
	}
	if !withOp {
		num, _ := strconv.Atoi(input)
		result = append(result, num)
	}
	return result
}
