package main

// [1, 1, 1, 1, 1 ,1 ,1, 1, 1, 1]
// 위와 같은 경우 시간 복잡도가 O(N)으로 되기 때문에
// O(logN) 복잡도에는 맞지 않음
func searchRange(nums []int, target int) []int {
	ans := []int{-1, -1}
	n := len(nums)

	if n == 0 {
		return ans
	}

	left := 0
	right := n - 1

	if nums[0] == target {
		ans[0] = 0
	}

	if nums[n-1] == target {
		ans[1] = n - 1
	}

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			ans = []int{mid, mid}
			left = mid
			right = mid
			for left >= 0 && nums[left] == target {
				ans[0] = min(left, ans[0])
				left--
			}

			for right <= n-1 && nums[right] == target {
				ans[1] = max(right, ans[1])
				right++
			}

			return ans
		}
	}

	return ans
}

// 최적화된 풀이
// 시작점을 찾는 이진탐색과 끝 점을 찾는 이진 탐색을 사용하여 O(logN) 복잡도를 유지
func searchRange2(nums []int, target int) []int {
	leftIndex := findBound(nums, target, true)

	if leftIndex == -1 {
		return []int{-1, -1}
	}

	rightIndex := findBound(nums, target, false)

	return []int{leftIndex, rightIndex}

}

func findBound(nums []int, target int, isStart bool) int {
	left := 0
	right := len(nums) - 1
	bound := -1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			bound = mid

			// 첫 번째 위치를 찾는 중이라면 더 앞쪽에 있는지 범위를 줄이고
			// 마지막 위치를 찾는 중이라면 더 뒤쪽에 있는지 범위를 줄임
			if isStart {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return bound
}
