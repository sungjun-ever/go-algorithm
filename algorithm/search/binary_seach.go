package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 10, 14, 42, 52, 99}
	target := 1
	size := len(arr)
	left := 0
	right := size - 1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			fmt.Println(mid, arr[mid])
			return
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	fmt.Println("404")
}
