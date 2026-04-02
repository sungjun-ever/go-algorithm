package main

import "math"

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
func maxProfit(prices []int) int {
	var minPrice = math.MaxInt
	var maximum int

	for _, p := range prices {
		minPrice = min(minPrice, p)

		if p > minPrice {
			currMaximum := p - minPrice
			maximum = max(maximum, currMaximum)
		}
	}

	return maximum
}
