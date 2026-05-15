package main

import "fmt"

// 0들을 배열 끝으로 이동 시켜야 한다.
func MoveZeros(nums []int) {
	// write 포인터 하나를 두고
	// 0이 아닌 경우 쓰고 포인터 증강
	// 순회를 마치고 wp++ when wp < size
	writeIdx := 0
	size := len(nums)

	for _, n := range nums {
		if n != 0 {
			nums[writeIdx] = n
			writeIdx++
		}
	}

	for i := writeIdx; i < size; i++ {
		nums[i] = 0
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12, 9, 0, 4, 0, 0}
	MoveZeros(nums)
	fmt.Println(nums)
}
