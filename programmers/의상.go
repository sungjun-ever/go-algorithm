func solution(clothes [][]string) int {
	// 조합 문제다
	// 각 의상 별로 나올 수 있는 조합은
	// 각 의상 수 + 안입는 경우의 수다
	// c1[n] * c2[n] .... -1
	// 전체를 안입는 경우의 수는 없기 때문에
	// 마지막에 -1을 한다

	class := make(map[string]int, len(clothes))
	for _, c := range clothes {
		class[c[1]]++
	}

	ans := 1
	for _, v := range class {
		ans *= v + 1
	}

	return ans - 1
}