package main

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	// ASC
	sort.Ints(nums)

	var res [][]int

	// 루프를 돌면서 map에 넣어줌
	for i := 0; i < len(nums)-2; i++ {

		// 기준 값이 1보다 커지면 더했을때 0 불가
		if nums[i] > 0 {
			break
		}

		// 현재 i 값이 이전과 같다면 다음 수로
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 투 포인터 운영
		left := i + 1
		right := len(nums) - 1

		for left < right {
			// 세 수의 합이 0인경우
			if nums[left]+nums[right]+nums[i] == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})

				// 각 포인터의 중복을 피하기 위해 움직임
				for left < right && nums[left] == nums[left+1] {
					left++
				}

				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if nums[left]+nums[right]+nums[i] < 0 { // 세 수의 합이 0보다 작은경우
				left++
			} else { // 세 수의 합이 0보다 큰경우
				right--
			}
		}

	}

	return res
}
