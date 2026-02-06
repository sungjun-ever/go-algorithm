package main

import "sort"

// map 사용
func maxOperations(nums []int, k int) int {
	// 순회를 돌면서 map에 넣어준다?
	// 짝이 있으면 map에서 제거, cnt++?
	cntMap := make(map[int]int)
	cnt := 0

	for _, n := range nums {
		if n >= k {
			continue
		}

		target := k - n

		if count, ok := cntMap[target]; ok && count > 0 {
			cnt++
			cntMap[target]--
		} else {
			cntMap[n]++
		}
	}
	return cnt
}

// 투포인터 사용버전
func maxOperations2(nums []int, k int) int {
	sort.Ints(nums) // O(n log n)
	left, right := 0, len(nums)-1
	ops := 0

	for left < right {
		sum := nums[left] + nums[right]
		if sum == k {
			ops++
			left++
			right--
		} else if sum < k {
			left++ // 합이 작으면 더 큰 숫자가 필요함
		} else {
			right-- // 합이 크면 더 작은 숫자가 필요함
		}
	}
	return ops
}
