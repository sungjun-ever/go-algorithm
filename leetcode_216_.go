package main

func combinationSum3(k int, n int) [][]int {
	// 1 ~ 9까지의 숫자중 (숫자는 중복될 수 없음)
	// k 개를 뽑아 n을 만들 수 있는 조합
	var ans [][]int
	var path []int

	var backtrack func(start, currentSum int)

	backtrack = func(start, currentSum int) {
		// 합이 목표를 넘어가면 종료
		if currentSum > n {
			return
		}

		if len(path) == k {
			if currentSum == n {
				ans = append(ans, append([]int{}, path...)) // path를 바로 넣어주면 참조 때문에 문제발생
			}
			return
		}

		for i := start; i <= 9; i++ {
			// 넣어주고
			path = append(path, i)
			// 재귀 탐색
			backtrack(i+1, currentSum+i)
			// 계산을 마치고 넣어준 값 다시 빼줌
			path = path[:len(path)-1]
		}
	}

	backtrack(1, 0)

	return ans
}
