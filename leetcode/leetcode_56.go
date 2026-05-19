package main

import "sort"

func merge(intervals [][]int) [][]int {
	// 서로 겹치는 간격이 있다면 합쳐 준다
	// 겹치는 구간을 어떻게 체크할 것인가
	// 원본 배열을 첫 요소 기준으로 정렬 시켜준다

	if len(intervals) == 1 {
		return intervals
	}

	var ans [][]int
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 루프를 돌며 첫 요소가 이전 배열에 포함된다면 같은 간격으로 취급한다
	size := len(intervals)
	start, end := intervals[0][0], intervals[0][1]
	var curr []int
	for i := 1; i < size; i++ {
		curr = intervals[i]
		// 범위 안이라면
		if curr[0] >= start && curr[0] <= end {
			start = min(start, curr[0])
			end = max(end, curr[1])
		} else { // 범위 밖이라면
			ans = append(ans, []int{start, end})
			start = curr[0]
			end = curr[1]
		}
	}

	if len(curr) > 0 {
		ans = append(ans, []int{start, end})
	}

	return ans
}
