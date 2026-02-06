package main

import "strconv"

func compress(chars []byte) int {
	// 투포인터
	w := 0 // 쓰기용
	r := 0 // 읽기용
	length := len(chars)

	for r < length {
		prevChar := chars[r]
		cnt := 0

		// 이전 문자가 현재 문자와 다를때까지 반복문
		for r < length && prevChar == chars[r] {
			r++
			cnt++
		}

		chars[w] = prevChar
		w++
		if cnt > 1 {
			cntToStr := strconv.Itoa(cnt)

			for i := 0; i < len(cntToStr); i++ {
				chars[w] = cntToStr[i]
				w++
			}
		}
	}

	return w
}
