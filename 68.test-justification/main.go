package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {

}

func sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func makeSpaces(num int) string {
	spaces := make([]string, num+1)
	return strings.Join(spaces, " ")
}

func justifyLine(words []string, wordLens []int, last bool, maxWidth int) string {
	spaceCount := maxWidth - sum(wordLens)
	wordsCount := len(words)
	if last || wordsCount == 1 {
		joined := strings.Join(words, " ")
		return joined + makeSpaces(maxWidth-len(joined))
	}
	spaceNums := make([]int, wordsCount-1)
	for {
		for i := range spaceNums {
			spaceNums[i]++
			spaceCount--
			if spaceCount <= 0 {
				goto buildStr
			}
		}
	}
buildStr:
	strBuffer := bytes.Buffer{}
	for i := 0; i < wordsCount-1; i++ {
		strBuffer.WriteString(words[i])
		strBuffer.WriteString(makeSpaces(spaceNums[i]))
	}
	strBuffer.WriteString(words[wordsCount-1])
	return strBuffer.String()
}

func fullJustify(words []string, maxWidth int) []string {
	result := make([]string, 0)
	lens := make([]int, len(words))

	for i := range words {
		lens[i] = len(words[i])
	}

	indexBufs := make([]int, 0)
	bufLen := 0
	for i := 0; i < len(words); i++ {
		indexBufs = append(indexBufs, i)
		fmt.Println(indexBufs)
		bufLen += lens[i]
		if i < len(words)-1 && bufLen+lens[i+1]+1 > maxWidth {
			newLine := justifyLine(
				words[indexBufs[0]:indexBufs[len(indexBufs)-1]+1],
				lens[indexBufs[0]:indexBufs[len(indexBufs)-1]+1],
				false,
				maxWidth,
			)
			result = append(result, newLine)
			indexBufs = make([]int, 0)
			bufLen = 0
		} else if i == len(words)-1 {
			newLine := justifyLine(
				words[indexBufs[0]:indexBufs[len(indexBufs)-1]+1],
				lens[indexBufs[0]:indexBufs[len(indexBufs)-1]+1],
				true,
				maxWidth,
			)
			result = append(result, newLine)
			indexBufs = make([]int, 0)
			bufLen = 0
		} else {
			bufLen++
		}
	}

	return result
}
