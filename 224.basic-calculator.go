import "strconv"

func calculate(s string) int {
	buf := ""
	result := 0
	var op byte = '+'
	for i := 0; i < len(s); i++ {
		if s[i] == '+' || s[i] == '-' {
			if len(buf) != 0 {
				bufValue, _ := strconv.Atoi(buf)
				if op == '+' {
					result += bufValue
				} else if op == '-' {
					result -= bufValue
				}
				buf = ""
			}
			op = s[i]
		} else if s[i] == '(' {
			l := 1
			r := 0
			j := i + 1
			for ; l != r; j++ {
				if s[j] == '(' {
					l++
				} else if s[j] == ')' {
					r++
				}
			}
			subResult := calculate(s[i+1 : j-1])
			if op == '+' {
				result += subResult
			} else if op == '-' {
				result -= subResult
			}
			op = '+'
			i = j -1
		} else if s[i] != ' '{
			buf += string(s[i])
		}
	}
    if len(buf) > 0 {
        bufValue, _ := strconv.Atoi(buf)
		if op == '+' {
			result += bufValue
		} else if op == '-' {
			result -= bufValue
		}
    }
	return result
}
