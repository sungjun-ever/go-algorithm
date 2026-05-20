package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 끝에서 n 번째에 있는 노드를 삭제한다
	// 두 개의 포인터를 두고, 하나의 포인터를 먼저 n번 만큼 전진시킨다

	// head 삭제를 생각해서 더미 노드를 만든다
	dummy := &ListNode{0, head}
	slow, fast := dummy, dummy

	// fast를 n번 만큼 먼저 움직인다
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// fast가 nil되기 전까지 움직인다
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// slow의 다음 노드가 삭제될 노드다
	slow.Next = slow.Next.Next

	return dummy.Next
}
