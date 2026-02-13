package main

func removeStars(s string) string {
	// 별을 만나면 별자신과 왼쪽을 지워줌
	// 문자는 스택에 넣고 별인경우
	// 스택 top을 제거
	stack := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return string(stack)
}

// ai가 사용한 append 오버헤드 줄이는 in-place 방식
// 포인터로 하는 방식
func removeStars2(s string) string {
	// 문자열을 수정 가능한 byte 슬라이스로 변환
	b := []byte(s)
	j := 0 // Write Pointer (스택의 Top 역할)

	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			// 별이 오면 j를 뒤로 당김 (즉, 방금 쓴 문자를 무효화)
			// 문제 보장상 j > 0 임이 확실하므로 검사 생략 가능
			j--
		} else {
			// 별이 아니면 현재 j 위치에 문자를 덮어쓰고 j 전진
			b[j] = s[i]
			j++
		}
	}

	// 0부터 j까지가 유효한 문자열
	return string(b[:j])
}
