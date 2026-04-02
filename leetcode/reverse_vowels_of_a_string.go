package main

func reverseVowels(s string) string {
	b := []byte(s)
	left, right := 0, len(s)-1

	for left < right {
		for left < right && !isVowel(b[left]) {
			left++
		}

		for left < right && !isVowel(b[right]) {
			right--
		}

		if left < right {
			b[left], b[right] = b[right], b[left]
			left++
			right--
		}
	}

	return string(b)
}

func isVowel(c byte) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}
