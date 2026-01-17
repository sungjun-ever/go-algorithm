package main

func reverseList(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode

	for curr != nil {
		curr.Next, prev, curr = prev, curr, curr.Next
	}

	return prev
}
