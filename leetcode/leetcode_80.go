package main

func removeDuplicates(nums []int) int {
	// 2번 나온 숫자는 배열에서 삭제한다
	if len(nums) <= 2 {
		return len(nums)
	}

	/*
		1	1	1	2	2	3
				i,p
		1	1	1	2	2	3
				p	i
		1	1	2	2	2	3
					p	i
		1	1	2	2	2	3
						p	i
	*/

	// 항상 덮어 씌우는 곳 기준으로 2칸전을 보게한다
	point := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[point-2] {
			nums[point] = nums[i]
			point++
		}
	}

	return point
}
