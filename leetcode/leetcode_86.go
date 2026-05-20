package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	// x를 보다 작은 값들은 x보다 크거나 같은 값들 앞에 와야한다
	// x보다 작은 값들을 연결시킬 리스트 하나랑, x보다 크거나 같은 값들을 저장하는 리스트 하나
	// 마지막에 둘이 연결?

	dummy1 := &ListNode{Next: nil}
	dummy2 := &ListNode{Next: nil}
	lt, gte := dummy1, dummy2

	for head != nil {
		if head.Val < x {
			lt.Next = head
			lt = lt.Next
		} else {
			gte.Next = head
			gte = gte.Next
		}
		head = head.Next
	}

	gte.Next = nil

	lt.Next = dummy2.Next

	return dummy1.Next
}
