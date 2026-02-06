package main

func moveZeroes(nums []int) {
	// 투포인터
	f := 0 // non-zero
	s := 0 // zero
	length := len(nums) - 1

	for f <= length && s <= length {
		// find non-zero index
		for f < length && nums[f] == 0 {
			f++
		}

		// find zero index
		for s < length && nums[s] != 0 {
			s++
		}

		if f > s {
			nums[f], nums[s] = nums[s], nums[f]
			s++
		} else {
			f++
		}
	}
}
