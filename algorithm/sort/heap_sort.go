package main

import "fmt"

// 힙 상태의 배열은 오름차순 또는 내림차순으로 값을 꺼내오는 우선순위 큐이다
// 주어진 배열을 빌드해서 힙을 만든 다음 값을 하나씩 꺼내오면 정렬된 순서로 나온다.
func main() {
	list := []int{38, 27, 43, 3, 9, 82, 10, 21}
	size := len(list)

	for i := (size / 2) - 1; i >= 0; i-- {
		heapify(list, size, i)
	}
	fmt.Println(list)
	for i := size - 1; i >= 0; i-- {
		// 가장 큰 값과 끝에 값의 위치를 바꿔주고
		list[0], list[i] = list[i], list[0]
		// 가장 끝을 제외하고 다시 힙 구조로 만들어준다
		heapify(list, i, 0)
	}
	fmt.Println(list)
}

func heapify(list []int, size, idx int) {
	largest := idx
	left := 2*idx + 1
	right := 2*idx + 2

	if left < size && list[largest] < list[left] {
		largest = left
	}

	if right < size && list[largest] < list[right] {
		largest = right
	}

	if largest != idx {
		list[largest], list[idx] = list[idx], list[largest]
		heapify(list, size, largest)
	}
}
