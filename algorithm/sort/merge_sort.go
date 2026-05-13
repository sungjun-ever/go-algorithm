package main

import "fmt"

// 시간 복잡도 O(n logn)
// 분할 정복 기법을 활용한 정렬이다
// 배열을 쪼갤 수 없을 만큼 나누고 나눈 배열을 합치면서 정렬을 한다.
// 데이터가 정렬되어있지 않아도 성능이 일관적이다.
// 원소의 순서를 유지하는 안정 정렬이다.
// 큰 데이터를 안정적으로 정렬해야하거나, 외부 저장장치와 함께 작업할 때 많이 사용한다.
func main() {
	arr := []int{5, 52, 21, 1, 100, 88, 76}
	size := len(arr)
	mergeSort(arr, 0, size-1)
	fmt.Println(arr)
}

func mergeSort(list []int, left, right int) {
	if left >= right {
		return
	}

	mid := left + (right-left)/2

	mergeSort(list, left, mid)
	mergeSort(list, mid+1, right)
	merge(list, left, mid, right)
}

func merge(list []int, left, mid, right int) {
	temp := make([]int, right-left+1)
	i, j, k := left, mid+1, 0

	for i <= mid && j <= right {
		if list[i] <= list[j] {
			temp[k] = list[i]
			k++
			i++
		} else {
			temp[k] = list[j]
			k++
			j++
		}
	}

	for i <= mid {
		temp[k] = list[i]
		k++
		i++
	}

	for j <= right {
		temp[k] = list[j]
		k++
		j++
	}

	for idx, t := range temp {
		list[left+idx] = t
	}
}
