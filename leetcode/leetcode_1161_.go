package main

import "math"

// 내 풀이
func maxLevelSum(root *TreeNode) int {
	levelSum := make([]int, 0, 1000)

	var dfs func(node *TreeNode, depth int)

	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		if depth == len(levelSum) {
			levelSum = append(levelSum, node.Val)
		} else {
			levelSum[depth] += node.Val
		}

		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}

	dfs(root, 0)

	currMax := -1000000
	level := 1

	for k, v := range levelSum {
		if v > currMax {
			level = k + 1
			currMax = v
		}
	}

	return level
}

// bfs
func maxLevelSum2(root *TreeNode) int {
	maxSum := math.MinInt32
	queue := []*TreeNode{root}

	maxLevel := 1
	currLevel := 1

	for len(queue) > 0 {
		length := len(queue)
		currentSum := 0

		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			currentSum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		if currentSum > maxSum {
			maxSum = currentSum
			maxLevel = currLevel
		}

		currLevel++
	}

	return maxLevel
}
