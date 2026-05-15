package main

import "fmt"

func ValidParentheses(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := make([]rune, 0, len(s))

	for _, c := range s {
		// 	문자가 닫는 괄호라면
		if open, ok := pairs[c]; ok && len(stack) > 0 {
			top := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			if open != top {
				return false
			}
		} else {
			stack = append(stack, c)
		}
	}

	return len(stack) == 0
}

func main() {
	s := "()]]]]]]]]]]]"
	fmt.Println(ValidParentheses(s))
}
