package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	// 바이너리 트리를 연결 리스트로 변경
	// preorder다
	// 스택을 사용해 오른쪽, 왼쪽 순서로 넣어줘 항상 왼쪽을 먼저 다루게한다
	if root == nil {
		return
	}

	stack := []*TreeNode{root}
	// in: 1
	// out:1, in: 5, 2
	// out: 2, in: 5, 4, 3
	// out: 3, in: 5, 4
	// out: 4, in: 5
	// out: 5, in: 6
	// out: 6, in: none
	var prev *TreeNode = nil

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}

		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}

		if prev != nil {
			prev.Right = curr
			prev.Left = nil
		}
		prev = curr
	}
}
