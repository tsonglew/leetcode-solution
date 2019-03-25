package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
	fmt.Println(restoreIpAddresses("0000"))
	fmt.Println(restoreIpAddresses("010010"))
}

func restoreIpAddresses(s string) []string {
	current := []string{}
	result := &[]string{}
	if len(s) < 12 {
		f(s, current, result)
	}
	return *result
}

func f(s string, current []string, result *[]string) {
	if len(s) == 0 {
		if len(current) == 4 {
			*result = append(*result, strings.Join(current, "."))
		}
		return
	}
	if s[0] <= 255 {
		newCurrent := append([]string{}, current...)
		f(s[1:], append(newCurrent, string(s[0])), result)
	}
	if len(s) >= 2 {
		i, _ := strconv.Atoi(s[0:2])
		if i <= 255 && i >= 10 {
			newCurrent := append([]string{}, current...)
			f(s[2:], append(newCurrent, string(s[0:2])), result)
		}
	}
	if len(s) >= 3 {
		i, _ := strconv.Atoi(s[0:3])
		if i <= 255 && i >= 100 {
			newCurrent := append([]string{}, current...)
			f(s[3:], append(newCurrent, string(s[0:3])), result)
		}
	}
}
