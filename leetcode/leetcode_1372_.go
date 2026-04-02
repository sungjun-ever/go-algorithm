package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestZigZag(root *TreeNode) int {
	maxLen := 0
	var dfs func(node *TreeNode, direction, length int)

	dfs = func(node *TreeNode, direction, length int) {
		if node == nil {
			return
		}

		maxLen = max(length, maxLen)

		if direction == 0 {
			dfs(node.Right, 1, length+1)
			dfs(node.Left, 0, 1)
		} else {
			dfs(node.Left, 0, length+1)
			dfs(node.Right, 1, 1)
		}
	}

	if root != nil {
		dfs(root.Right, 1, 1)
		dfs(root.Left, 0, 1)
	}

	return maxLen
}
