package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// 전체 길이를 구한다
	tail := head
	length := 1
	for tail.Next != nil {
		tail = tail.Next
		length++
	}

	k = k % length

	if k == 0 {
		return head
	}

	// 순환 리스트로 만들어준다
	tail.Next = head
	// 1 2 3 4 5 1 2 3 4 5 ~
	newTail := head
	for i := 0; i < length-k-1; i++ {
		newTail = newTail.Next
	}

	// 새 헤드를 만들어서 새 시작 지점과 연결 시켜준다
	newHead := newTail
	newHead = newTail.Next
	newTail.Next = nil

	return newHead
}
