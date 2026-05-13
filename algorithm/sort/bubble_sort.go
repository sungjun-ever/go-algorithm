package main

import "fmt"

// 버블 정렬은 시간 복잡도 O(n2) 을 가진다
// 버블 정렬은 인접 요소와 대소를 비교 후 스왑한다.
// 배열이 어느 정도 정렬되어있으면 버블 정렬은 일찍 종료 가능하다.
func main() {
	arr := []int{5, 52, 21, 1, 100, 88, 76}
	size := len(arr)

	for i := 0; i < size; i++ {
		swapped := false
		for j := 0; j < size-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	fmt.Println(arr)
}
