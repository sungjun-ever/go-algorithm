package main

// https://leetcode.com/problems/top-k-frequent-elements/description/
func topKFrequent(nums []int, k int) []int {

	countMap := make(map[int]int)
	// 각 숫자의 빈도 수 체크
	for _, n := range nums {
		countMap[n]++
	}

	// 빈도수 별로 숫자를 묶음
	buckets := make([][]int, len(nums)+1)
	for num, cnt := range countMap {
		buckets[cnt] = append(buckets[cnt], num)
	}

	res := []int{}
	// 빈도수 내림차순으로 k개 만큼 숫자를 가져옴
	for i := len(buckets) - 1; i >= 0; i-- {
		if len(buckets[i]) > 0 {
			for _, n := range buckets[i] {
				res = append(res, n)

				if len(res) == k {
					return res
				}
			}
		}
	}

	return res
}
