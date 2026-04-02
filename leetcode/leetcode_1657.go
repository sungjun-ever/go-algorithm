package main

import "sort"

func closeStrings(word1 string, word2 string) bool {
	// 1. 문자열의 길이가 같아야함
	// 2. 서로 같은 문자를 가지고 있어야함
	// 3. 빈도수가 같아야함

	// 길이가 다르면 false
	if len(word1) != len(word2) {
		return false
	}

	var c1 [26]int
	var c2 [26]int

	// 빈도수 구하기
	for _, c := range word1 {
		c1[c-'a']++
	}

	for _, c := range word2 {
		c2[c-'a']++
	}

	// 문자를 가지고 있지 않은지 체크
	for i := 0; i < 26; i++ {
		if (c1[i] == 0 && c2[i] > 0) || (c2[i] == 0 && c1[i] > 0) {
			return false
		}
	}

	// 배열을 슬라이스로 변
	sort.Ints(c1[:])
	sort.Ints(c2[:])

	// 빈도수가 같은지 체크
	for i := 0; i < len(c1); i++ {
		if c1[i] != c2[i] {
			return false
		}
	}

	return true
}
