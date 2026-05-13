package main

import "fmt"

// 선택 정렬은 시간 복잡도 O(n2) 을 가진다
// 구현은 간단하지만 효율은 낮은 정렬 방식이다
// 버블 정렬과 달리 배열의 상태와 상관없이 같은 수의 연산을 한다.
func main() {
	arr := []int{5, 52, 21, 1, 100, 88, 76}
	size := len(arr)

	for i := 0; i < size; i++ {
		minIndex := i
		for j := i + 1; j < size; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		arr[i], arr[minIndex] = arr[minIndex], arr[i]

	}

	fmt.Println(arr)
}
