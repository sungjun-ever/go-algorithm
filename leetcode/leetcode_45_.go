package main

func jump(nums []int) int {
	// 목표 도달을 해야한다

	target := len(nums) - 1
	jumps := 0   // 점프 횟수
	currEnd := 0 // 현재 점프의 한계치
	far := 0     // 현재 갈 수 있는 가장 먼 곳

	for i := 0; i < target; i++ {
		// 현재 포지션에서 갈 수 있는 가장 먼 곳을 갱신 가능하면 갱신
		if i+nums[i] > far {
			far = i + nums[i]
		}

		// 인덱스가 현재 점프 한계치에 도달하면
		// 다시 점프를하고 현재 점프 한계치를 가장 먼 곳으로 갱신한다
		if i == currEnd {
			currEnd = far
			jumps++

			// 이미 도달 가능하면 종료
			if currEnd >= target {
				break
			}
		}
	}

	return jumps
}
