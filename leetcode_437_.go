package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
	prefixMap := make(map[int]int)

	prefixMap[0] = 1

	return dfs(root, 0, targetSum, prefixMap)
}

func dfs(node *TreeNode, currentSum int, target int, prefixMap map[int]int) int {
	if node == nil {
		return 0
	}

	currentSum += node.Val

	count := prefixMap[currentSum-target]

	prefixMap[currentSum]++
	count += dfs(node.Left, currentSum, target, prefixMap)
	count += dfs(node.Right, currentSum, target, prefixMap)

	prefixMap[currentSum]--

	return count
}
