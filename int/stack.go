package main

import (
	"fmt"
)

type node struct {
	Val     int
	CurrMin int
}

type Stack struct {
	stack []node
}

func Constructor() *Stack {
	return &Stack{
		stack: []node{},
	}
}

func (s *Stack) Push(val int) {
	min := val

	if len(s.stack) > 0 {
		topMin := s.stack[len(s.stack)-1].CurrMin
		if topMin < min {
			min = topMin
		}
	}

	s.stack = append(s.stack, node{
		Val:     val,
		CurrMin: min,
	})

}

func (s *Stack) Pop() int {
	curr := s.stack[len(s.stack)-1]
	s.stack = s.stack[0 : len(s.stack)-1]
	return curr.Val
}

func (s *Stack) Top() int {
	return s.stack[len(s.stack)-1].Val
}

func (s *Stack) GetMin() int {
	return s.stack[len(s.stack)-1].CurrMin
}

func main() {
	s := Constructor()
	s.Push(7)
	s.Push(2)
	fmt.Println(s.GetMin())
	s.Push(3)
	s.Push(4)
	s.Push(5)
	fmt.Println(s.Top())
	fmt.Println(s)
	fmt.Println(s.Pop())
	fmt.Println(s.Top())
	s.Push(1)
	s.Push(5)
	fmt.Println(s.GetMin())
}
