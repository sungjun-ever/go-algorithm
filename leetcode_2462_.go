package main

import (
	"container/heap"
	"math"
)

type IntHeap []int64

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int64))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func totalCost(costs []int, k int, candidates int) int64 {
	lHeap := &IntHeap{}
	rHeap := &IntHeap{}

	heap.Init(lHeap)
	heap.Init(rHeap)

	left, right := 0, len(costs)-1

	// 매 루프마다 왼쪽 오른쪽을 따로 해서 포인터가 넘어가는지 지속 확인
	for i := 0; i < candidates; i++ {
		if left <= right {
			heap.Push(lHeap, int64(costs[left]))
			left++
		}
		if left <= right {
			heap.Push(rHeap, int64(costs[right]))
			right--
		}
	}

	var totalCost int64 = 0

	for i := 0; i < k; i++ {
		// 힙이 비어있는 경우에 대비해 int64 최댓값으로 초기화
		var lVal, rVal int64 = math.MaxInt64, math.MaxInt64

		if lHeap.Len() > 0 {
			lVal = (*lHeap)[0]
		}
		if rHeap.Len() > 0 {
			rVal = (*rHeap)[0]
		}

		if lVal <= rVal {
			totalCost += heap.Pop(lHeap).(int64)
			if left <= right {
				heap.Push(lHeap, int64(costs[left]))
				left++
			}
		} else {
			totalCost += heap.Pop(rHeap).(int64)
			if left <= right {
				heap.Push(rHeap, int64(costs[right]))
				right--
			}
		}
	}

	return totalCost
}
