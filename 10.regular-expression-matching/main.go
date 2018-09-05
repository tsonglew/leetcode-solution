package regular_expression_matching

type pattern struct {
	b byte
	star bool
}

func isMatch(s string, p string) bool {
	patterns := []*pattern{}
	for i := 0; i < len(p); i++ {
		if p[i] == '*' {
			patterns[len(patterns)-1].star = true
		} else {
			patterns = append(patterns, &pattern{
				b: p[i],
				star: false,
			})
		}
	}
	matchMap := make([][]bool, len(patterns)+1)
	for i := range matchMap {
		matchMap[i] = make([]bool, len(s)+1)
	}
	matchMap[0][0] = true
	for i := 1; i < len(s)+1; i++ {
		matchMap[0][i] = false
	}
	for i := 1; i<len(patterns)+1;i++{
		if patterns[i-1].star && matchMap[i-1][0] == true {
			matchMap[i][0] = true
		} else {
			matchMap[i][0] = false
		}
	}
	for i := 1; i < len(patterns)+1; i++ {
		for j := 1; j < len(s) +1; j++ {
			if matchMap[i-1][j-1] == true {
				if patterns[i-1].b == s[j-1] || patterns[i-1].b=='.' {
					matchMap[i][j] = true
					continue
				}
			}
			if matchMap[i][j-1] == true { // byte left
				if (patterns[i-1].b == s[j-1] || patterns[i-1].b=='.') && patterns[i-1].star {
					matchMap[i][j] = true
					continue
				}
			}
			if matchMap[i-1][j] == true {
				if patterns[i-1].star {
					matchMap[i][j] = true
					continue
				}
			}
			matchMap[i][j] = false
		}
	}
	return matchMap[len(patterns)][len(s)]
}
