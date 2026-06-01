package programmers

func solution(brown int, yellow int) []int {
	// 노란색 부분 크기를 만족시키는 공식은 사각형의 넓이를 구하는 공식을 생각하면 된다
	// row 기준으로 전체 row 수에서 테두리 두 줄을 빼고, col 기준으로 전체 col 수에서 테두리 두 줄을 뺀다
	// yellow = w-2 * h-2
	// 전체 넓이를 만들 수 있는 수의 조합해서 위 공식을 만족시키는 수의 조합을 찾으면 된다.

	// 노란색 부분이 갈색 부분 안쪽에 오려면 전체 최소 높이는 3이 되어야한다.
	total := brown + yellow
	for h := 3; h*h <= total; h++ {
		if total%h != 0 {
			continue
		}

		w := total / h

		if (w-2)*(h-2) == yellow {
			return []int{w, h}
		}
	}

	return []int{}
}
