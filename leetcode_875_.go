package main

func minEatingSpeed(piles []int, h int) int {
	// 최소치 1시간, 최대치 더미 중 가장 큰 수
	left := 1
	right := 1

	for _, p := range piles {
		if p > right {
			right = p
		}
	}

	ans := right

	for left <= right {
		mid := left + (right-left)/2

		var totalHours int = 0
		// 전체 덩어리를 먹는데 걸리는 시간 계산
		for _, p := range piles {
			totalHours += calHours(p, mid, 0)
		}

		// 전체 시간이 h보다 작거나 같으면 더 한 번에 먹는 양을 줄여서 시간 줄일 수 있는지 확인
		// 아니면 한 번에 먹는 양을 늘림
		if totalHours <= h {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}

func calHours(pile, mid, total int) int {
	a := pile / mid
	b := pile % mid
	total += a
	if b == 0 {
		return total
	} else if b <= mid {
		return total + 1
	}

	return calHours(b, mid, total)
}
