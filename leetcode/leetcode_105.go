package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	// preorder, inorder 배열의 주어질 때 원형 바이너리 트리를 리턴
	// preorder의 첫 요소는 트리의 헤드다
	// inorder에서 헤드를 기준으로 왼쪽이 왼쪽 서브 트리, 오른쪽이 오른쪽 서브 트리
	if len(preorder) == 0 {
		return nil
	}

	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	// preorder의 요소의 첫 째는 부모 노드가 될 수 있다
	// preorder의 첫 요소를 기준으로 inorder에서 찾아 좌우를 나누면 왼쪽은 왼쪽 서브트리로 오른쪽은 반대로
	idx := 0
	for i, n := range inorder {
		if preorder[0] == n {
			idx = i
			break
		}
	}

	left, right := []int{}, []int{}
	left = inorder[0:idx]

	if idx < len(inorder)-1 {
		right = inorder[idx+1:]
	}

	root := &TreeNode{Val: preorder[0]}

	// 부모 기준으로 왼쪽 오른쪽을 나눈다
	// preorder는 inorder left, right 수를 기준으로 나눠준다
	root.Left = buildTree(preorder[1:len(left)+1], left)
	root.Right = buildTree(preorder[len(left)+1:], right)

	return root
}
