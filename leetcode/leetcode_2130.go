package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
	maxSum := 0
	nodeSlice := make([]int, 0, 50000)
	var prev *ListNode
	curr := head
	length := 0

	for curr != nil {
		nodeSlice = append(nodeSlice, curr.Val)
		curr.Next, prev, curr = prev, curr, curr.Next
		length++
	}

	for i := 0; i < length/2; i++ {
		maxSum = max(nodeSlice[i]+prev.Val, maxSum)
		prev = prev.Next
	}

	return maxSum
}

// 공간 복잡도 줄인 방법
func pairSum2(head *ListNode) int {
	// 중간까지 이동 시키고 중간부터 뒤집어준다
	// 처음과 끝을 더하면서 가장 큰값 구함
	maxSum := 0
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 현재 slow 시작점부터 뒤집어줌
	var prev *ListNode
	curr := slow

	for curr != nil {
		curr.Next, prev, curr = prev, curr, curr.Next
	}

	p1, p2 := head, prev

	for p2 != nil {
		maxSum = max(maxSum, p1.Val+p2.Val)
		p1 = p1.Next
		p2 = p2.Next
	}

	return maxSum
}
