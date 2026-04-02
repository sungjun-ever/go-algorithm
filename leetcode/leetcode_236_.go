package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	type state struct {
		node   *TreeNode
		parent *TreeNode
		step   int
	}

	nodeMap := make(map[int]state)

	findP := false
	findQ := false

	var dfs func(node, parent *TreeNode, step int)
	dfs = func(node, parent *TreeNode, step int) {
		if node == nil {
			return
		}

		step++
		nodeMap[node.Val] = state{node, parent, step}

		if node == p {
			findP = true
		} else if node == q {
			findQ = true
		}

		dfs(node.Left, node, step)
		dfs(node.Right, node, step)

		if findP && findQ {
			return
		}

	}

	dfs(root, root, 0)

	upStep := 0
	currentP := nodeMap[p.Val].node
	currentQ := nodeMap[q.Val].node
	currentStep := min(nodeMap[p.Val].step, nodeMap[q.Val].step)

	if nodeMap[p.Val].step <= nodeMap[q.Val].step {
		upStep = nodeMap[q.Val].step - nodeMap[p.Val].step

		for i := 0; i < upStep; i++ {
			currentQ = nodeMap[currentQ.Val].parent
		}
	} else {
		upStep = nodeMap[p.Val].step - nodeMap[q.Val].step
		for i := 0; i < upStep; i++ {
			currentP = nodeMap[currentP.Val].parent
		}
	}

	for currentStep >= 0 {

		if currentP.Val == currentQ.Val {
			return currentP
		}

		parentP := nodeMap[currentP.Val].parent
		parentQ := nodeMap[currentQ.Val].parent

		if currentP.Val == parentQ.Val {
			return currentP
		}

		if currentQ.Val == parentP.Val {
			return currentQ
		}

		currentP = nodeMap[currentP.Val].parent
		currentQ = nodeMap[currentQ.Val].parent
		currentStep--
	}

	return root
}

// ai 버전
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	// 1. 종료 조건: 바닥에 닿았거나, 내가 찾던 p나 q를 발견한 경우
	if root == nil || root == p || root == q {
		return root
	}

	// 2. 왼쪽과 오른쪽 서브트리에 각각 탐색 지시 (분할)
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 3. 양쪽에서 모두 뭔가(p나 q)를 찾아왔다면, 내가 바로 공통 조상!
	if left != nil && right != nil {
		return root
	}

	// 4. 한쪽에서만 찾아왔다면, 찾은 쪽의 결과를 위로 그대로 패스
	if left != nil {
		return left
	}

	return right
}
