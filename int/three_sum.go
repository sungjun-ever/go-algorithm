package main

import (
	"fmt"
	"sort"
)

// 3개 요소의 합이 0이 되는 set의 배열을 리턴한다.
func ThreeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	var ans [][]int
	// 정렬을 진행해 시작 요소가 0보다 커지면 종료한다
	sort.Ints(nums)

	// 2개의 포인터를 운영한다.
	// 현재 요소의 다음부터 시작하는 포인터 1
	// 맨 끝부터 출발하는 포인터 2
	size := len(nums)
	for start := 0; start < size-2; start++ {
		left := start + 1
		right := size - 1

		// 시작 요소가 0보다 커지면 불가능
		if nums[start] > 0 {
			break
		}

		// 첫 루프를 제외하고 start와 이전 요소가 같으면 이미 확인한 경우이기 때문에 중복이 발생함
		if start > 0 && nums[start] == nums[start-1] {
			continue
		}

		for left < right {
			currentSum := nums[start] + nums[left] + nums[right]
			if currentSum == 0 {
				ans = append(ans, []int{nums[start], nums[left], nums[right]})
				left++
				right--
			} else if currentSum > 0 {
				right--
			} else {
				left++
			}
		}
	}

	return ans
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(ThreeSum(nums))
}
