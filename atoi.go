package main

import (
	"fmt"
	"math"
)

func main() {
	var input string
	fmt.Scan(&input)
	isNegitive := false

	if len(input) == 0 {
		panic("empty string")
	}

	if input[0] == '-' {
		input = input[1:]
		isNegitive = true
	} else if input[0] == '+' {
		input = input[1:]
	}

	nums := 0
	for _, s := range input {
		if s < '0' || s > '9' {
			panic("invalid character")
		}

		digit := int(s - '0')

		if nums > (math.MaxInt32-digit)/10 {
			panic("overflow")
		}

		nums = nums*10 + digit
	}

	if isNegitive {
		nums = -nums
	}

	fmt.Println(nums)
}
