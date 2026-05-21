package main

func buildTree(inorder []int, postorder []int) *TreeNode {
	// inorder, postorder가 주어질 때
	// 바이너리 트리 만들기
	// inorder에서 root 노드 기준 좌우로 서브 트리를 만들 수 있다
	// postorder의 마지막 요소가 첫 노드다
	// postorder기준으로 끝부터 ~ 까지 오른쪽 노드가 만들어진다
	// inorder에서 첫 노드를 기준으로 좌, 우 시작 노드를 찾는다
	// postorder에서 왼쪽 시작 노드를 만나기 전까지가 오른쪽 서브트리다

	inorderMap := make(map[int]int, len(inorder))
	for i, val := range inorder {
		inorderMap[val] = i
	}

	postIdx := len(postorder) - 1

	var recursive func(start, end int) *TreeNode

	recursive = func(start, end int) *TreeNode {
		if start > end {
			return nil
		}

		rootVal := postorder[postIdx]
		root := &TreeNode{Val: rootVal}
		idx := inorderMap[rootVal]
		postIdx--

		root.Right = recursive(idx+1, end)
		root.Left = recursive(start, idx-1)

		return root
	}

	return recursive(0, len(inorder)-1)
}

func buildTree2(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootVal}
	idx := 0

	// 루트 찾기
	for i, n := range inorder {
		if n == rootVal {
			idx = i
			break
		}
	}

	left, right := inorder[0:idx], []int{}
	if idx < len(inorder)-1 {
		right = inorder[idx+1:]
	}

	root.Left = buildTree(left, postorder[0:len(left)])
	root.Right = buildTree(right, postorder[len(left):len(postorder)-1])

	return root
}
