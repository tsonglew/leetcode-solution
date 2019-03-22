package main

func numDecodingsTime(s string) int {
	cnt := 0
	cache := make(map[int]int)
	f1(&cnt, &cache, s, 0)
	return cnt
}

func f1(cnt *int, cache *map[int]int, s string, i int) {
	if len(s) == 0 {
		(*cnt)++
		return
	}
	if c, ok := (*cache)[i]; ok {
		(*cnt) += c
		return
	}
	if int(s[0]) > int('0') {
		f1(cnt, cache, s[1:], i+1)
		(*cache)[i] = *cnt
	}
	if s[0] == '1' && len(s) > 1 ||
		s[0] == '2' && len(s) > 1 && int(s[1]) <= int('6') {
		f1(cnt, cache, s[2:], i+2)
		(*cache)[i] = *cnt
	}
}
