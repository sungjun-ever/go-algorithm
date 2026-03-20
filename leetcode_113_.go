package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	// root에서 출발해 마지막 노드까지 갔을 때, 목표에 도달하는 모든 경우 구하기

	ans := [][]int{}
	path := []int{}

	if root == nil {
		return ans
	}

	var findNodes func(node *TreeNode, total int)
	findNodes = func(node *TreeNode, total int) {
		if node == nil {
			return
		}

		total += node.Val
		path = append(path, node.Val)

		if total == targetSum && node.Left == nil && node.Right == nil {
			// 슬라이스는 참조
			// 새 슬라이스를 만들어 깊은 복사로 현재 경로 저장 후 넣어줌
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
		}

		findNodes(node.Left, total)
		findNodes(node.Right, total)
		path = path[:len(path)-1]
	}

	findNodes(root, 0)

	return ans
}
