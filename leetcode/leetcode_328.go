package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	odd, even := head, head.Next
	evenHead := even

	for even != nil && even.Next != nil {
		// 홀수의 다음은 짝수의 다음
		odd.Next = even.Next // odd.Next.Next 와 같음
		odd = odd.Next

		// 짝수의 다음은 홀수의 다음
		even.Next = odd.Next // even.Next.Next 와 같음
		even = even.Next
	}

	odd.Next = evenHead // odd와 even 연결

	return head

}
