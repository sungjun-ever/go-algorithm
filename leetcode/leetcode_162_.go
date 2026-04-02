package main

func findPeakElement(nums []int) int {
	// 산 정상 숫자
	// 오른쪽, 왼쪽보다 크면 정상
	left, right := 0, len(nums)-1

	// left와 right가 만나면 정상
	for left < right {
		mid := left + (right-left)/2

		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}
