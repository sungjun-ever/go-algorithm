package main

func rob(nums []int) int {
	prev2 := 0 // 전전
	prev1 := 0 // 전

	// 두, 세 번째부터 턴다고 가정했을때
	// 세 번째는 첫 번째 + 현재, 두 번째는 두 번째
	// 둘 중 더 큰걸 더해준다
	for _, n := range nums {
		curr := max(prev2+n, prev1)

		prev2 = prev1
		prev1 = curr
	}

	return prev1

}
