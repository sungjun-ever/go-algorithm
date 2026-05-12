package main

func canJump(nums []int) bool {
	// 각 인덱스의 값 만큼 [최대]로 점프할 수 있다
	target := len(nums) - 1
	maxReach := 0

	// 현재 맥스 값보다 더 멀리 이동가능하면 갱신해준다

	for i, n := range nums {
		// 현재 갈 수 있는 거리보다 인덱스가 더 크다면 종료
		if i > maxReach {
			return false
		}

		maxReach = max(maxReach, i+n)

		if maxReach >= target {
			return true
		}
	}

	return maxReach >= target

}

func canJump2(nums []int) bool {
	// 각 인덱스의 값 만큼 [최대]로 점프할 수 있다
	goal := len(nums) - 1

	// 마지막 인덱스에서 출발해 첫 인덱스에 도착 가능한지
	for i := goal; i >= 0; i-- {
		if i+nums[i] >= goal {
			goal = i
		}
	}

	return goal == 0

}
