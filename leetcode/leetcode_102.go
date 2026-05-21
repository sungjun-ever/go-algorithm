package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	// 각 레벨별 배열을 리턴, 순서는 왼쪽에서 오른쪽으로
	var ans [][]int

	if root == nil {
		return ans
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		// 현재 레벨의 값들을 넣어줌
		levelSize := len(queue)

		temp := make([]int, levelSize)

		for i := 0; i < levelSize; i++ {
			curr := queue[0]
			queue = queue[1:]
			temp[i] = curr.Val

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}

			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}

		ans = append(ans, temp)
	}

	return ans
}
