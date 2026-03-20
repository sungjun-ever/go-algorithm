package main

// strconv를 사용하지 않고
// 자릿수를 늘려가며 해야하는 경우
// prev * 10 + curr
func sumNumbers(root *TreeNode) int {
	// root부터 끝까지의 숫자들의 경우를 다 더하기
	// 4 -> 9 -> 5 = 495
	// 4 -> 0 = 40
	// 495 + 40

	var sum func(node *TreeNode, currentSum int) int

	sum = func(node *TreeNode, currentSum int) int {
		if node == nil {
			return 0
		}

		currentSum = currentSum*10 + node.Val

		if node.Left == nil && node.Right == nil {
			return currentSum
		}

		return sum(node.Left, currentSum) + sum(node.Right, currentSum)
	}

	return sum(root, 0)
}
