package main

// https://leetcode.com/problems/valid-anagram/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	letters := [26]int{}

	for _, c := range s {
		letters[c-97]++
	}

	for _, c := range t {
		letters[c-97]--
	}

	for _, v := range letters {
		if v != 0 {
			return false
		}
	}

	return true
}
