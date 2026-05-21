package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	// 지그재그로 레벨별 기록을 해서 리턴
	// 순서는 -> <-
	var ans [][]int

	if root == nil {
		return ans
	}

	queue := []*TreeNode{root}

	// 역방향 체크용
	isReverse := false

	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, size)

		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			idx := i
			if isReverse {
				idx = size - i - 1
			}

			level[idx] = curr.Val

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}

			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}

		ans = append(ans, level)
		isReverse = !isReverse
	}

	return ans
}
