package main

import "fmt"

func main() {
	var input string
	fmt.Scan(&input)
	isNegitive := false

	if input[0] == '-' {
		input = input[1:]
		isNegitive = true
	}

	nums := 0
	for _, s := range input {
		nums = nums*10 + int(s-'0')
	}

	if isNegitive {
		nums = -nums
	}
	fmt.Printf("%T", nums)
	fmt.Println(nums)
}
