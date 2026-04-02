package main

import "math"

// 가장 작은수, 두 번째로 작은 수를 계속 구하고
// 현재 수가 그 두 개의 수보다 큰 경우 true
// 아무것도 없으면 false
func increasingTriplet(nums []int) bool {
	first, second := math.MaxInt, math.MaxInt

	for _, n := range nums {
		if n <= first {
			first = n
		} else if n <= second {
			second = n
		} else {
			return true
		}
	}

	return false
}
