package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

// 연결 리스트 중간에 있는 노드를 삭제한다.
func RemoveMiddleNode(head *Node) {
	if head == nil {
		return
	}

	// 중간을 어떻게 찾을 것인가?
	// 두 개의 포인터를 운영한다.
	curr := head
	slow, fast := curr, curr.Next.Next

	/*
		루프가 종료되고 slow의 다음이 중간 노드다
		1	3	4	7	2	6
		s		f
			s			f
				s				f
	*/

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	slow.Next = slow.Next.Next
}

func main() {
	head := &Node{
		Val: 1,
		Next: &Node{
			Val: 3,
			Next: &Node{
				Val: 4,
				Next: &Node{
					Val: 7,
					Next: &Node{
						Val: 2,
						Next: &Node{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}

	RemoveMiddleNode(head)
	curr := head
	for curr != nil {
		fmt.Print(curr.Val)
		fmt.Print("----->")
		curr = curr.Next
	}
}
