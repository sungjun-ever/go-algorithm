package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	nodeMap := make(map[*Node]*Node)

	// 각 노드 -  노드의 값만 가지는 노드를 맵핑한다
	curr := head
	for curr != nil {
		nodeMap[curr] = &Node{Val: curr.Val}
		curr = curr.Next
	}

	// 노드 맵에 있는 노드의 Next, Random을 노드맵에서 가져와 맵핑한다
	curr = head
	for curr != nil {
		nodeMap[curr].Next = nodeMap[curr.Next]
		nodeMap[curr].Random = nodeMap[curr.Random]
		curr = curr.Next
	}

	return nodeMap[head]
}

func copyRandomList2(head *Node) *Node {
	// 1. 원본 노드 다음에 복사 노드를 연결하고 복사 노드 다음에 원본 노드의 다음을 연결한다
	curr := head
	for curr != nil {
		next := curr.Next
		copyNode := &Node{Val: curr.Val}
		curr.Next = copyNode
		copyNode.Next = next
		curr = next
	}
	// 현재 상태
	// A -> A` -> B -> B` -> ...... -> nil
	// A의 랜덤 노드가 B였다면
	// A`의 랜덤 노드는 B`가 되어야 한다
	// curr.Next.Random (A.B.B`) = curr.Random.Next (A.B.B`)

	// 2. 이제 랜덤 포인터를 연결한다
	curr = head
	for curr != nil {
		if curr.Random != nil {
			curr.Next.Random = curr.Random.Next
		}
		curr = curr.Next.Next
	}

	// 하나의 리스트로 있는 원본 리스트와 복사 리스트를 나눈다
	curr = head
	dummy := &Node{}
	copyNode := dummy
	for curr != nil {
		// 복사 노드는 원본 노드의 다음 노드이다
		copyNext := curr.Next
		copyNode.Next = copyNext
		copyNode = copyNext

		// 원본 노드는 두 칸씩 전진해야 원본 노드끼리 이어진다
		next := curr.Next.Next
		curr.Next = next
		curr = next
	}

	return dummy.Next

}
