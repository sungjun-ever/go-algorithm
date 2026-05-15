package main

import "fmt"

// 중복된 문자가 없는 가장 긴 서브스트링의 길이
func LongestSubstring(s string) int {
	// 문자의 위치를 기록해야한다
	// 어떻게 기록할 것이냐? 맵으로 키를 문자 값을 인덱스로?
	maxLen := 0
	feq := make(map[rune]int, len(s))
	left := 0 // 서브 문자열의 길이를 따질때 시작하는 포인터

	for i, c := range s {
		// 맵에 이미 문자가 있다면?
		if idx, ok := feq[c]; ok {
			// 시작 포인터를 이전 문자의 인덱스 다음으로 바꿔준다
			left = idx + 1
		}

		// 맵에 문자의 인덱스를 기록한다
		feq[c] = i
		maxLen = max(maxLen, i-left+1)
	}

	return maxLen
}

func main() {
	s := "aaaaababcdefgabcdefghijks"
	fmt.Println(LongestSubstring(s))
}
