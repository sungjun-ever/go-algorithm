package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode) {
	// [1,5,3,4,2]
	// first 처음 어긋남 5 < 3 -> 5가 범인
	// second 다음 어긋남 4 < 2 -> 2가 범인
	var first, second, prev *TreeNode
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)

		// 앞의 값이 더 크면 틀어짐 발생
		if prev != nil && prev.Val > node.Val {
			// 첫 번째 틀어짐은 전 노드가 문제
			if first == nil {
				first = prev
			}

			// 두 번째 틀어짐은 현재 노드가 문제
			second = node
		}

		prev = node

		inorder(node.Right)
	}

	inorder(root)

	if first != nil && second != nil {
		first.Val, second.Val = second.Val, first.Val
	}
}
