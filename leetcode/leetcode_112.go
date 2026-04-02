package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	var sum func(node *TreeNode, total int) bool
	sum = func(node *TreeNode, total int) bool {
		if node == nil {
			return false
		}

		total += node.Val

		if total == targetSum && node.Left == nil && node.Right == nil {
			return true
		}

		return sum(node.Left, total) || sum(node.Right, total)
	}

	return sum(root, 0)
}
