package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	// 먼저 오른쪽 노드들을 뒤집어준다
	// 그리고 왼쪽과 오른쪽이 같은지 확인한다
	invert(root.Right)
	return isSame(root.Left, root.Right)
}

func invert(node *TreeNode) {
	if node == nil || node.Left == nil && node.Right == nil {
		return
	}

	invert(node.Left)
	invert(node.Right)
	node.Left, node.Right = node.Right, node.Left
}

func isSame(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	return isSame(p.Left, q.Left) && isSame(p.Right, q.Right)
}
