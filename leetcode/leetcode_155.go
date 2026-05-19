package main

type Node struct {
	Val int
	Min int
}

type MinStack struct {
	stack []Node
}

func Constructor() MinStack {
	return MinStack{stack: []Node{}}
}

func (this *MinStack) Push(val int) {
	currMin := val
	if len(this.stack) > 0 {
		top := this.stack[len(this.stack)-1]
		currMin = min(currMin, top.Min)
	}

	this.stack = append(this.stack, Node{
		Val: val,
		Min: min(val, currMin),
	})
}

func (this *MinStack) Pop() {
	this.stack = this.stack[0 : len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1].Val
}

func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1].Min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
