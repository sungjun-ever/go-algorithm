package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaf1 := make([]int, 0, 200)
	leaf2 := make([]int, 0, 200)

	findLeaf(root1, &leaf1)
	findLeaf(root2, &leaf2)

	if len(leaf1) != len(leaf2) {
		return false
	}

	for i := 0; i < len(leaf1); i++ {
		if leaf1[i] != leaf2[i] {
			return false
		}
	}

	return true

}

func findLeaf(node *TreeNode, leafSlice *[]int) {
	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil {
		*leafSlice = append(*leafSlice, node.Val)
		return
	}

	findLeaf(node.Left, leafSlice)
	findLeaf(node.Right, leafSlice)

}
