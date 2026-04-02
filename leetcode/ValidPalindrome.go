package main

import "unicode"

func isAlphanumeric(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsNumber(c)
}

func isPalindrome(s string) bool {
	// 투 포인터로 각 포인터가 위치한 문자가 같은지 체크
	// 문자, 숫자인 경우 해당 포인터 위치 이동

	left := 0
	right := len(s) - 1

	for left < right {
		for left < right && !isAlphanumeric(rune(s[left])) {
			left++
		}

		for left < right && !isAlphanumeric(rune(s[right])) {
			right--
		}

		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}

		left++
		right--
	}

	return true
}

// 아스키 코드 직접 비교를 통한 최적화 방법
func isPalindrome2(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		// 왼쪽 포인터가 알파벳/숫자가 아니면 건너뜀
		for l < r && !isAlnum(s[l]) {
			l++
		}
		// 오른쪽 포인터가 알파벳/숫자가 아니면 건너뜀
		for l < r && !isAlnum(s[r]) {
			r--
		}

		// 대소문자 구분 없이 비교
		if toLower(s[l]) != toLower(s[r]) {
			return false
		}
		l++
		r--
	}
	return true
}

// 직접 구현한 가벼운 영문자/숫자 판별 함수
func isAlnum(b byte) bool {
	return (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9')
}

// 직접 구현한 가벼운 소문자 변환 함수
func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + ('a' - 'A') // ASCII 차이를 이용한 변환
	}
	return b
}
