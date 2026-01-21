package main

func isValidBST(root *TreeNode) bool {
	return validate(root, nil, nil)
}

func validate(node *TreeNode, min *int, max *int) bool {
	if node == nil {
		return true
	}

	if (min != nil && node.Val <= *min) || (max != nil && node.Val >= *max) {
		return false
	}

	return validate(node.Left, min, &node.Val) && validate(node.Right, &node.Val, max)
}
