package main

import (
	"fmt"
)

func main() {
	var input int
	fmt.Scan(&input)

	if input == 0 {
		fmt.Println("0")
		return
	}

	n := int64(input)

	isNegative := false

	if n < 0 {
		isNegative = true
		n = -n
	}

	result := make([]byte, 0, 11)

	for n > 0 {
		result = append(result, byte(n%10)+'0')
		n /= 10
	}

	if isNegative {
		result = append(result, '-')
	}

	left, right := 0, len(result)-1

	for left < right {
		result[left], result[right] = result[right], result[left]
		left++
		right--
	}

	fmt.Println(string(result))

}
