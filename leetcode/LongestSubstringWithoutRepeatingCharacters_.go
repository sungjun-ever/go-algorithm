package main

// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
func lengthOfLongestSubstring(s string) int {
	lastPos := [128]int{}

	for i := range lastPos {
		lastPos[i] = -1
	}

	maxLen := 0
	l := 0

	for r := 0; r < len(s); r++ {
		char := s[r]
		// lastPos에 있는 경우
		if lastPos[char] >= l {
			l = lastPos[char] + 1
		}

		// lastPos에 없는 경우
		lastPos[char] = r
		maxLen = max(maxLen, r-l+1)
	}

	return maxLen
}
