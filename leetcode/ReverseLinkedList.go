package main

// prev // curr
// nil // 1 -> 2 -> 3 -> nil
// nil <- 1 // 2 -> 3 -> nil
// nil <- 1 <- 2 // 3 -> nil
// nil <- 1 <- 2 <- 3 // nil
func reverseList(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode

	for curr != nil {
		curr.Next, prev, curr = prev, curr, curr.Next
	}

	return prev
}
