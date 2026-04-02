package main

import (
	"container/heap"
	"sort"
)

type Score struct {
	n1 int64
	n2 int64
}

type ScoreHeap []Score

// Len, Less, Swap, Pop, Push

func (h ScoreHeap) Len() int { return len(h) }

// n1 기준 오름차순 정렬
func (h ScoreHeap) Less(i, j int) bool { return h[i].n1 < h[j].n1 }
func (h ScoreHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ScoreHeap) Push(x any) {
	*h = append(*h, x.(Score))
}

func (h *ScoreHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	h := &ScoreHeap{}
	heap.Init(h)
	scoreSlice := make([]Score, len(nums1))
	lenght := len(nums1)
	for i := 0; i < lenght; i++ {
		scoreSlice[i] = Score{int64(nums1[i]), int64(nums2[i])}
	}

	// nums2 기준 내림차순 정렬
	// 힙에 넣을 때 현재 n2 값이 항상 최솟값이 되도록
	sort.Slice(scoreSlice, func(i, j int) bool {
		return scoreSlice[i].n2 > scoreSlice[j].n2
	})

	var currentSum int64 = 0
	heapLen := h.Len()
	var ans int64 = 0
	for i := 0; i < lenght; i++ {
		curr := scoreSlice[i]
		currentSum += curr.n1
		heap.Push(h, Score{scoreSlice[i].n1, scoreSlice[i].n2})
		heapLen++

		if heapLen > k {
			popped := heap.Pop(h).(Score)
			currentSum -= popped.n1
			heapLen--
		}

		if heapLen == k {
			score := currentSum * curr.n2

			if ans < score {
				ans = score
			}
		}
	}

	return ans
}
