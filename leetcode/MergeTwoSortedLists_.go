package main

// https://leetcode.com/problems/merge-two-sorted-lists/description/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 첫 노드가 어떤게 될지 모르기 때문에 새 연결리스트 만듦
	dummy := &ListNode{}
	curr := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}

		curr = curr.Next
	}

	if list1 != nil {
		curr.Next = list1
	} else if list2 != nil {
		curr.Next = list2
	}

	// 	dummy 노드의 첫 노드가 빈값이기 때문에 다음 노드로 이동해서
	return dummy.Next
}
