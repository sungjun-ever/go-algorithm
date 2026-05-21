package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	// 각 레벨별 왼쪽 -> 오른쪽으로 노드를 연결하고 오른쪽에 더이상 노드가 없다면 '#'으로 끝낸다
	// 어떻게 할까? 먼저 끝까지 내려간다?
	// queue 만들어서 레벨별로 이어준다?
	if root == nil {
		return nil
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		var prev *Node = nil

		for i := 0; i < levelSize; i++ {
			curr := queue[0]
			queue = queue[1:]

			if prev != nil {
				prev.Next = curr
			}

			prev = curr

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}

			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
	}

	return root
}
