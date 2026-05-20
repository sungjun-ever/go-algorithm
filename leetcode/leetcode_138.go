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
