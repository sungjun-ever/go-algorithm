package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	// 각 레벨의 평균 값
	level := []*TreeNode{root}
	var ans []float64

	for len(level) > 0 {
		levelSize := len(level)
		total := 0
		for i := 0; i < levelSize; i++ {
			curr := level[0]
			level = level[1:]
			total += curr.Val

			if curr.Left != nil {
				level = append(level, curr.Left)
			}

			if curr.Right != nil {
				level = append(level, curr.Right)
			}
		}

		avg := float64(total) / float64(levelSize)
		ans = append(ans, avg)
	}

	return ans
}
