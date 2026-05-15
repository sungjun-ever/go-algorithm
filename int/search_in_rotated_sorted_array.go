package main

import "fmt"

func Search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	// 범위를 어떻게 줄일 것이냐
	size := len(nums)
	left := 0
	right := size - 1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		if target == nums[left] {
			return left
		}

		if target == nums[right] {
			return right
		}

		// 가운데를 기준으로 왼쪽 탐색한다.
		if nums[left] <= nums[mid] {
			// 목표가 왼쪽 ~ 중간 사이에 있는 경우
			if target > nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // 가운데를 기준으로 오른쪽 탐색
			// 목표가 중간 ~ 오른쪽 사이에 았는 경우
			if target > nums[mid] && target < nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

	}

	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(Search(nums, 2))
}
