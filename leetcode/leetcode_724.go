package main

func pivotIndex(nums []int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}

	leftSum := 0

	for i, n := range nums {
		if leftSum == total-leftSum-n {
			return i
		}

		leftSum += n
	}

	return -1
}
