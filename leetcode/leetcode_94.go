package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	// 가장 왼쪽에 있는 노드부터 담아준다
	ans := []int{}

	var loop func(node *TreeNode)

	loop = func(node *TreeNode) {
		if node == nil {
			return
		}

		loop(node.Left)
		ans = append(ans, node.Val)
		loop(node.Right)
	}

	loop(root)

	return ans
}
