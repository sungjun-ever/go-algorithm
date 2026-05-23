package main

func rotate(nums []int, k int) {
	// 배열 회전
	// 1 2 3 4 5 6 7
	// 7 6 5 4 3 2 1
	// 5 6 7 1 2 3 4
	// 앞에서 k개를 뒤집는다
	// 나머지를 뒤집는다
	if len(nums) == 0 {
		return
	}
	// nums=[1, 2, 3], k = 5
	// 3 1 2 => 1
	// 2 3 1 => 2
	// 1 2 3 => 3
	// 3 1 2 => 4
	// 2 3 1 => 5
	if k > len(nums) {
		k = k % len(nums)
	}

	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)

}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
