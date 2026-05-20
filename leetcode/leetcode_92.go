package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// left, right가 주어질 때 해당 포지션에 사이에 있는 노드들을 역순으로(left, right 포함)
	dummy := &ListNode{Val: 0, Next: head}
	before := dummy

	// left 앞까지 이동시킨다
	for i := 0; i < left-1; i++ {
		before = before.Next
	}

	// left의 시작점
	start := before.Next

	// start를 고정된 위치로 생각하고
	// start 앞에 있는 노드를 항상 start 자리로 보내준다
	for i := 0; i < right-left; i++ {
		// 바뀌는 대상은 start의 Next
		next := start.Next

		start.Next = next.Next  // 2 4
		next.Next = before.Next // 3 2
		before.Next = next      // 1 3
	}

	return dummy.Next
}
