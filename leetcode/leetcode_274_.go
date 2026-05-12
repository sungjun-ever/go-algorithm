package main

import "sort"

func hIndex(citations []int) int {
	// h번 이상 인용된 논문이 h편 이상인 최대 h를 찾는 문제

	sort.Slice(citations, func(i, j int) bool {
		return citations[i] >= citations[j]
	})

	var h int

	for i := 0; i < len(citations); i++ {
		if citations[i] >= i+1 {
			h = i + 1
		} else {
			break
		}
	}

	return h
}
