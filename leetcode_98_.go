package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	// root는 왼쪽보다 커야하고, 오른쪽보다는 작아야한다
	var loop func(node *TreeNode, min, max *int) bool

	loop = func(node *TreeNode, min, max *int) bool {
		if node == nil {
			return true
		}

		if (min != nil && node.Val <= *min) || (max != nil && node.Val >= *max) {
			return false
		}

		return loop(node.Left, min, &node.Val) && loop(node.Right, &node.Val, max)
	}

	return loop(root, nil, nil)
}
