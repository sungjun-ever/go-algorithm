package programmers

func solution(n int) int64 {
	// 한 칸 또는 두 칸을 뛸 수 있다
	// 칸이 n개 주어질 때 미자막 칸에 도착할 수 있는 경우의 수를 1234567로 나눈 수를 리턴한다.
	if n == 1 {
		return int64(1)
	}

	// n == 0 => 0
	// n == 1 => 1
	// n == 2 => 2 (1, 1 // 2)
	// n == 3 => 3(1,1,1 / 1, 2 / 2, 1)

	prev := int64(1)
	curr := int64(2)

	for i := 3; i <= n; i++ {
		next := (prev + curr) % 1234567
		prev = curr
		curr = next
	}

	return curr

}
