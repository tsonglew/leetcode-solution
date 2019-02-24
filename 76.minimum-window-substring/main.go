package main

import (
	"strings"
)

func main() {

}

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	}
	minWindow, minWindowLen := []int{0, 0}, len(s)+1
	leftPointer, rightPointer := 0, 0
	chars := strings.Split(t, "")
	sChars := strings.Split(s, "")
	charsMap := make(map[string]int)
	charsMapThreshold := make(map[string]int)
	charsInWindowCount := 0
	charsCount := 0
	for i := range chars {
		if _, ok := charsMap[chars[i]]; !ok {
			charsMap[chars[i]] = 0
		}
		if _, ok := charsMapThreshold[chars[i]]; !ok {
			charsMapThreshold[chars[i]] = 1
		} else {
			charsMapThreshold[chars[i]]++
		}
		charsCount++
	}

	for rightPointer < len(s) && leftPointer < len(s) {
		for ; rightPointer < len(s); rightPointer++ {
			if v, ok := charsMap[sChars[rightPointer]]; ok {
				if v < charsMapThreshold[sChars[rightPointer]] {
					charsInWindowCount++
				}
				charsMap[sChars[rightPointer]]++

				if charsInWindowCount == charsCount {
					rightPointer++
					if rightPointer-leftPointer < minWindowLen {
						minWindow = []int{leftPointer, rightPointer}
						minWindowLen = rightPointer - leftPointer
					}
					break
				}
			}
		}

		if rightPointer == len(s) && charsInWindowCount < charsCount {
			break
		}

		for ; leftPointer < rightPointer; leftPointer++ {
			if v, ok := charsMap[sChars[leftPointer]]; ok {
				charsMap[sChars[leftPointer]]--
				if v == charsMapThreshold[sChars[leftPointer]] {
					charsInWindowCount--
					if rightPointer-leftPointer < minWindowLen {
						minWindow = []int{leftPointer, rightPointer}
						minWindowLen = rightPointer - leftPointer
					}
					leftPointer++
					break
				}
			}
		}
	}
	return strings.Join(sChars[minWindow[0]:minWindow[1]], "")
}
