package main

import "fmt"

func main() {
	list := []int{38, 27, 43, 3, 9, 82, 10, 21}
	n := len(list)
	fmt.Println("before", list)
	sort(list, 0, n-1)
	fmt.Println("after", list)
}

func sort(list []int, left, right int) {
	if left < right {
		q := partition(list, left, right)

		sort(list, left, q-1)
		sort(list, q+1, right)
	}
}

func partition(list []int, left, right int) int {
	pivot := list[left] // 첫 원소를 기준으로
	low := left + 1     // 첫 원소 다음 원소를 최솟값으로 잡음
	high := right

	for low <= high { // 교차할 때까지 반복
		for low <= right && list[low] <= pivot {
			low++
		}

		for left <= high && list[high] > pivot {
			high--
		}

		if low < high {
			list[low], list[high] = list[high], list[low]
		}

	}

	// 피벗 위치를 가운데로 옮김
	list[left], list[high] = list[high], list[left]

	// 피벗 위치 리턴
	return high
}
