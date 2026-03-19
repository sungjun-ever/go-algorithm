package main

import "math"

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	sign := 1

	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		sign = -1
	}

	var quotient int64 = 0

	absDividend := abs(int64(dividend))
	absDivisor := abs(int64(divisor))

	for absDividend >= absDivisor {
		var tempDivisor int64 = absDivisor
		var multiple int64 = 1

		for absDividend >= (tempDivisor << 1) {
			tempDivisor <<= 1
			multiple <<= 1
		}

		absDividend -= tempDivisor
		quotient += multiple
	}

	if sign == -1 {
		quotient = -quotient
	}

	return int(quotient)
}

func abs(n int64) int64 {
	if n < 0 {
		return -n
	}

	return n
}
