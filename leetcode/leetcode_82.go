package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	// 정렬된 연결리스트가 있을 때, 중복된 값이 있는 모든 노드를 지운다
	// 헤드가 삭제될 경우를 대비해 더미 노드 생성
	dummy := &ListNode{0, head}
	prev := dummy

	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			prev.Next = head.Next
		} else {
			prev = prev.Next
		}
		head = head.Next
	}

	return dummy.Next
}
