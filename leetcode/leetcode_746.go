package main

func minCostClimbingStairs(cost []int) int {
	// 0, 1 위치에서 시작 가능
	// 코스트를 지불하는 경우 1, 2칸 이동 가능
	n := len(cost)
	prev2 := cost[0]
	prev1 := cost[1]

	for i := 2; i < n; i++ {
		curr := cost[i] + min(prev1, prev2)

		prev2 = prev1
		prev1 = curr
	}

	return min(prev2, prev1)
}
