package main

// https://leetcode.com/problems/container-with-most-water/description/

func maxArea(height []int) int {
	// 투 포인터 운영
	l, r := 0, len(height)-1

	// 최대 넓이
	var maxWidth int

	for l < r {
		// 현재 넓이
		currentWidth := (r - l) * min(height[l], height[r])

		// 비교
		maxWidth = max(maxWidth, currentWidth)

		// 더 작은쪽의 포인터를 움직임
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}

	return maxWidth
}
