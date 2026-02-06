package main

func maxAreas(height []int) int {
	// area = (yIndex - xIndex) * secondHeight
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		currArea := (right - left) * min(height[right], height[left])
		maxArea = max(maxArea, currArea)

		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}
