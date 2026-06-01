package programmers

func solution(s string) int {
	n := len(s)
	answer := 0

	for start := 0; start < n; start++ {
		if isValid(s, start) {
			answer++
		}
	}

	return answer
}

func isValid(s string, start int) bool {
	pairs := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	n := len(s)
	stack := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		idx := (i + start) % n
		ch := s[idx]

		if ch == '}' || ch == ']' || ch == ')' {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if pairs[ch] != top {
				return false
			}

		} else {
			stack = append(stack, ch)
		}
	}

	return len(stack) == 0
}
