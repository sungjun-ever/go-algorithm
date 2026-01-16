package main

// https://leetcode.com/problems/valid-parentheses/description/
func isValid(s string) bool {
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []byte{}

	for i := 0; i < len(s); i++ {
		char := s[i]

		// 닫는 괄호인 경우
		if open, ok := pairs[char]; ok {
			// 스택이 비어있거나 짝이 아닌경우
			if len(stack) == 0 || open != stack[len(stack)-1] {
				return false
			}

			// 짝이 맞으면 제거
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}
