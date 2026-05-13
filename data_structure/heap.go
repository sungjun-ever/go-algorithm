package main

import "fmt"

type Heaps struct {
	heap   []int
	length int
}

func (h *Heaps) insert(val int) {
	h.heap = append(h.heap, val)
	h.length++
	h.heapifyUp(h.length - 1)
}

func (h *Heaps) heapifyUp(index int) {
	parent := (index - 1) / 2
	if index > 0 && h.heap[index] < h.heap[parent] {
		h.heap[index], h.heap[parent] = h.heap[parent], h.heap[index]
		h.heapifyUp(parent)
	}
}

func (h *Heaps) heapifyDown(idx int) {
	left := 2*idx + 1
	right := 2*idx + 2
	smallest := idx
	size := len(h.heap)

	if left < size && h.heap[left] < h.heap[smallest] {
		smallest = left
	}

	if right < size && h.heap[right] < h.heap[smallest] {
		smallest = right
	}

	if smallest != idx {
		h.heap[smallest], h.heap[idx] = h.heap[idx], h.heap[smallest]
		h.heapifyDown(smallest)
	}

}

func (h *Heaps) removeMin() *int {
	if h.length == 0 {
		return nil
	}

	min := h.heap[0]
	last := h.heap[h.length-1]
	h.heap = h.heap[0 : h.length-1]
	h.length--

	if h.length > 0 {
		h.heap[0] = last
		h.heapifyDown(0)
	}

	return &min

}

func main() {
	h := &Heaps{
		heap:   []int{4, 1, 32, 13, 9, 52, 6},
		length: 7,
	}

	for i := (h.length / 2) - 1; i >= 0; i-- {
		h.heapifyDown(i)
	}

	fmt.Println(h.heap)
	fmt.Println(*h.removeMin())
	fmt.Println(*h.removeMin())
	fmt.Println(h.heap)
}
