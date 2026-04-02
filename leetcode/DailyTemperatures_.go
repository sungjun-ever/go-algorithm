package main

// https://leetcode.com/problems/daily-temperatures/description/

func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0, len(temperatures))
	res := make([]int, len(temperatures))

	for i, curr := range temperatures {

		// 스택이 비어있지 않고, 현재 온도가 top보다 높은 경우
		for len(stack) > 0 && curr > temperatures[stack[len(stack)-1]] {
			// 하나씩 빼면서 온도차를 구해줌
			prevIndex := stack[len(stack)-1]
			res[prevIndex] = i - prevIndex
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	return res
}
