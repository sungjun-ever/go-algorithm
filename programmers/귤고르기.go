package programmers

import "sort"

func solution(k int, tangerine []int) int {
	// 크기 별로 분류된 귤이 있다.
	// k 개의 귤을 담을 때
	// 크기가 서로 다른 종류의 수를 최소화하려 한다.

	// 먼저 귤의 크기 빈도수를 체크한다.
	// 서로 다른 종류를 최소화 하려면 빈도수가 높은 귤부터 차례대로 k개를 채우면 된다
	sizeMap := make(map[int]int, len(tangerine))

	// 크기별로 개수를 증가 시킨다.
	for _, t := range tangerine {
		sizeMap[t]++
	}

	counts := make([]int, 0, len(sizeMap))

	// 개수를 배열에 넣어준다
	for _, v := range sizeMap {
		counts = append(counts, v)
	}

	// 개수가 많은 순서대로 정렬을 시킨다.
	sort.Slice(counts, func(i, j int) bool {
		return counts[i] > counts[j]
	})

	count := 0
	for _, size := range counts {
		k -= size
		count++

		if k <= 0 {
			break
		}

	}

	return count
}
