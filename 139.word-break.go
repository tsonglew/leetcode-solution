func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	legals := make([]bool, len(s))
	for i := range s {
		for j := range wordDict {
			starti := i - len(wordDict[j]) + 1
			if starti > 0 && s[starti:i+1] == wordDict[j] && legals[i-len(wordDict[j])] == true || starti == 0 && s[starti:i+1] == wordDict[j] {
				legals[i] = true
				break
			}
		}
	}
	return legals[len(legals)-1]
}