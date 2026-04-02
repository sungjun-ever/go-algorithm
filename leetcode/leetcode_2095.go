package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
	size := 0
	curr := head

	for curr != nil {
		curr = curr.Next
		size++
	}

	midIndex := size / 2

	curr2 := head

	i := 0

	if size == 1 {
		return head.Next
	}

	for i < midIndex-1 {
		curr2 = curr2.Next
		i++
	}

	if i+2 <= size {
		curr2.Next = curr2.Next.Next
	}

	return head
}

func deleteMiddle2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head
	fast = fast.Next.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	slow.Next = slow.Next.Next

	return head
}
