package main

// 내 풀이
func rightSideView(root *TreeNode) []int {
	levelSlice := make([]int, 0, 1000)

	if root == nil {
		return []int{}
	}

	var dfs func(node *TreeNode, depth int)

	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		levelSlice = append(levelSlice, -101)
		levelSlice[depth] = node.Val

		depth++
		dfs(node.Left, depth)
		dfs(node.Right, depth)

	}

	dfs(root, 0)

	answer := make([]int, 0, 1000)

	for _, v := range levelSlice {
		if v == -101 {
			break
		}
		answer = append(answer, v)
	}

	return answer
}

// ai dfs
func rightSideView2(root *TreeNode) []int {
	var answer []int

	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		// 🌟 핵심: 현재 깊이가 정답 배열의 길이와 같다면?
		// = 이 깊이에 '처음' 도달했다는 뜻! (오른쪽부터 탐색했으므로 무조건 가장 우측 노드)
		if depth == len(answer) {
			answer = append(answer, node.Val)
		}

		// 기존의 Left -> Right 순서를 Right -> Left로 변경
		dfs(node.Right, depth+1) // 오른쪽을 먼저 깊게 파고듦
		dfs(node.Left, depth+1)  // 그다음 왼쪽을 탐색
	}

	dfs(root, 0)
	return answer
}

// ai bfs
func rightSideView3(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var answer []int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue) // 현재 층에 있는 노드의 개수

		for i := 0; i < levelSize; i++ {
			curr := queue[0]
			queue = queue[1:] // PopLeft

			// 현재 층의 가장 마지막 노드라면 정답에 추가
			if i == levelSize-1 {
				answer = append(answer, curr.Val)
			}

			// 다음 층의 자식들을 큐에 추가
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
	}

	return answer
}
