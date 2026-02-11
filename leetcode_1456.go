package main

func maxVowels(s string, k int) int {
	isVowel := func(b byte) bool {
		return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
	}

	currCnt := 0

	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			currCnt++
		}
	}

	maxCnt := currCnt
	if maxCnt == k {
		return k
	}

	for i := k; i < len(s); i++ {
		if isVowel(s[i]) {
			currCnt++
		}

		if isVowel(s[i-k]) {
			currCnt--
		}

		maxCnt = max(maxCnt, currCnt)

		if maxCnt == k {
			return k
		}
	}

	return maxCnt
}
