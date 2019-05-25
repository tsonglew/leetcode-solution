func isPalindrome(s string) bool {
	bs := []byte(s)
	a, b := 0, 0
	for i, j := 0, len(bs)-1; i < j; {
		if bs[i] <= '9' && bs[i] >= '0' || bs[i] >= 'a' && bs[i] <= 'z' {
			a = int(bs[i])
		} else if bs[i] >= 'A' && bs[i] <= 'Z' {
			a = int(bs[i]) + 32
		} else {
			i++
			continue
		}
		if bs[j] <= '9' && bs[j] >= '0' || bs[j] >= 'a' && bs[j] <= 'z' {
			b = int(bs[j])
		} else if bs[j] >= 'A' && bs[j] <= 'Z' {
			b = int(bs[j]) + 32
		} else {
			j--
			continue
		}
		if a != b {
			return false
		}
		i++
		j--
	}
	return true
}