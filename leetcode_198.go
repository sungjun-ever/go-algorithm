package main

func rob(nums []int) int {
	prev2 := 0 // 전전
	prev1 := 0 // 전

	for _, n := range nums {
		curr := max(prev2+n, prev1)

		prev2 = prev1
		prev1 = curr
	}

	return prev1

}
