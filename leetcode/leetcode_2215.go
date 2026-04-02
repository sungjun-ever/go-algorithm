package main

func findDifference(nums1 []int, nums2 []int) [][]int {
	cache1 := make(map[int]int)
	cache2 := make(map[int]int)

	for _, n := range nums1 {
		cache1[n] = 1
	}

	for _, n := range nums2 {
		cache2[n] = 1
	}

	answer := make([][]int, 2)
	for k, _ := range cache1 {
		if _, ok := cache2[k]; !ok {
			answer[0] = append(answer[0], k)
		}
	}

	for k, _ := range cache2 {
		if _, ok := cache1[k]; !ok {
			answer[1] = append(answer[1], k)
		}
	}

	return answer
}
