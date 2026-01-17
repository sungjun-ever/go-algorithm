package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode

	for curr != nil {
		curr.Next, prev, curr = prev, curr, curr.Next
	}

	return prev
}
