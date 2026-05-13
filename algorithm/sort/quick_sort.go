package main

import "fmt"

// 평균적으로 매운 빠른 정렬 속도를 가지며 추가적인 메모리 사용이 적은 고성능 정렬이다
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
	// 기준점 가운데
	mid := left + (right-left)/2 //
	pivot := list[mid]
	l, r := left, right

	for l <= r {
		for list[l] < pivot {
			l++
		}

		for list[r] > pivot {
			r--
		}

		if l <= r {
			list[l], list[r] = list[r], list[l]
			l++
			r--
		}
	}

	return l
	// 기준점 오른쪽
	// pivot := list[right]
	// i := left - 1

	// for j := left; j <= right; j++ {
	// 	if list[j] < pivot {
	// 		i++
	// 		list[i], list[j] = list[j], list[i]
	// 	}
	// }

	// list[right], list[i+1] = list[i+1], list[right]
	// return i + 1

	// 기준점 왼쪽
	// pivot := list[left]
	// i := left

	// for j := left + 1; j <= right; j++ {
	// 	if list[j] < pivot {
	// 		i++
	// 		list[i], list[j] = list[j], list[i]
	// 	}
	// }

	// list[left], list[i] = list[i], list[left]

	// return i
}
