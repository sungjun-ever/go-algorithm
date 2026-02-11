package main

func longestOnes(nums []int, k int) int {
	// 두 개의 포인터가 있음
	// 하나는 계속가고
	// 하나는 0 카운트가 k가 될때까지 당김
	left, zeroCnt, maxLen := 0, 0, 0
	for right, v := range nums {
		if v == 0 {
			zeroCnt++
		}

		for zeroCnt > k {
			if nums[left] == 0 {
				zeroCnt--
			}
			left++
		}

		currLen := right - left + 1
		maxLen = max(maxLen, currLen)
	}

	return maxLen
}
