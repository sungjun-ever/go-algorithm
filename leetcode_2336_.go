package main

import "container/heap"

type IntHeap []int

// Len, Less, Swap, Push, Pop

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type SmallestInfiniteSet struct {
	nextNum   int          // 대기열 시작 번호
	addedHeap *IntHeap     // set에 다시 들어갈 없는 숫자를 정렬할 힙
	addedSet  map[int]bool // 숫자 중복을 막는 셋
}

func Constructor() SmallestInfiniteSet {
	h := &IntHeap{}
	heap.Init(h)

	return SmallestInfiniteSet{
		nextNum:   1,
		addedHeap: h,
		addedSet:  make(map[int]bool),
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	// 힙에 숫자가 있는 경우 먼저 내보냄
	// 힙에 숫자가 있으면 현재 대기열 번호보다 작은 숫자가 있음
	if this.addedHeap.Len() > 0 {
		smallest := heap.Pop(this.addedHeap).(int)

		delete(this.addedSet, smallest) // 셋에서도 없는 숫자 처리

		return smallest
	}

	// 힙이 없으면 원래 대기열에서 내보냄
	ans := this.nextNum
	this.nextNum++ // 다음 대기열 번호로 바꿔줌
	return ans
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	// 셋에도 없고 현재 대기열보다 작은 경우에만 힙에 넣을 수 있음
	if !this.addedSet[num] && num < this.nextNum {
		heap.Push(this.addedHeap, num)
		this.addedSet[num] = true // 셋에 있음 처리
	}
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
