package main

// https://leetcode.com/problems/reorder-list/description/

func reorderList(head *ListNode) {
	// 전반부, 후반부 리스트를 나눔
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	secondHalfList := slow.Next // slow.Next ~ end
	slow.Next = nil             // 연결을 끊어줌 head ~ slow

	// 후반부 리스트를 뒤집음
	var prev *ListNode // 뒤집은 후반부 리스트
	curr := secondHalfList
	for curr != nil {
		curr.Next, prev, curr = prev, curr, curr.Next
	}

	// 합침
	first := head
	second := prev

	for second != nil {
		tmp1, tmp2 := first.Next, second.Next
		first.Next, second.Next = second, tmp1
		first, second = tmp1, tmp2
	}
}
