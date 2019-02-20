package main

func main() {

}

func min(ints ...int) int {
	r := ints[0]
	for i := range ints {
		if ints[i] < r {
			r = ints[i]
		}
	}
	return r
}

func minDistance(word1 string, word2 string) int {
	word1Len := len(word1)
	word2Len := len(word2)

	if word1Len == 0 || word2Len == 0 {
		return word1Len + word2Len
	}

	table := make([][]int, word1Len)
	for i := range table {
		table[i] = make([]int, word2Len)
	}
	for i, in := 0, false; i < word1Len; i++ {
		if word1[i] == word2[0] {
			in = true
		}
		if in {
			table[i][0] = i
		} else {
			table[i][0] = i + 1
		}

	}
	for j, in := 0, false; j < word2Len; j++ {
		if word2[j] == word1[0] {
			in = true
		}
		if in {
			table[0][j] = j
		} else {
			table[0][j] = j + 1
		}
	}
	for i := 1; i < word1Len; i++ {
		for j := 1; j < word2Len; j++ {
			if word1[i] == word2[j] {
				table[i][j] = table[i-1][j-1]
			} else {
				table[i][j] = min(table[i-1][j-1], table[i-1][j], table[i][j-1]) + 1
			}
		}
	}
	return table[word1Len-1][word2Len-1]
}
