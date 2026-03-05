package main

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		minRight := root.Right
		for minRight.Left != nil {
			minRight = minRight.Left
		}

		root.Val = minRight.Val
		root.Right = deleteNode(root.Right, minRight.Val)
	}

	return root
}
