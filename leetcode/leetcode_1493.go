package main

func longestSubarray(nums []int) int {
	maxLen := 0
	currLen := 0
	zeroCnt := 0
	left := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroCnt++
		}

		for zeroCnt > 1 {
			if nums[left] == 0 {
				zeroCnt--
			}
			left++
		}

		currLen = right - left

		maxLen = max(maxLen, currLen)
	}

	return maxLen

}
