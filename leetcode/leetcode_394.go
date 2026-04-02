package main

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	// 반복 수 스택
	// 문자열 저장 스택
	var countStack []int
	var strStack []string

	currStr := "" // 현재 문자열
	r := 0        // 반복수
	for _, c := range s {
		// 숫자인경우 반복수를 구함
		if c >= '0' && c <= '9' {
			n, _ := strconv.Atoi(string(c))
			r = r*10 + n
		} else if c == '[' { // 여는 괄호
			countStack = append(countStack, r)
			r = 0

			strStack = append(strStack, currStr)
			currStr = ""

		} else if c == ']' { // 닫는 괄호
			count := countStack[len(countStack)-1]
			countStack = countStack[:len(countStack)-1]

			prevStr := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			currStr = prevStr + strings.Repeat(currStr, count)
		} else {
			currStr += string(c)
		}
	}

	return currStr
}
