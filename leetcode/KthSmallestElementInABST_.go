package main

func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	curr := root

	for root != nil || len(stack) == 0 {

		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k -= 1

		if k == 0 {
			return curr.Val
		}

		curr = curr.Right
	}

	return -1
}
