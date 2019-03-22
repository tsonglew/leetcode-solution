package main

func numDecodingsMem(s string) int {
	cnt := 0
	f2(&cnt, s)
	return cnt
}

func f2(cnt *int, s string) {
	if len(s) == 0 {
		(*cnt)++
		return
	}
	if int(s[0]) > int('0') {
		f2(cnt, s[1:])
	}
	if s[0] == '1' && len(s) > 1 ||
		s[0] == '2' && len(s) > 1 && int(s[1]) <= int('6') {
		f2(cnt, s[2:])
	}
}
