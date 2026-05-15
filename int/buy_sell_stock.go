package main

import (
	"fmt"
	"math"
)

// 최대의 이익을 리턴
func BuyAndSellStock(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	maxProfit := 0
	// int32 맥스 값으로 초기화 해준다
	minPrice := math.MaxInt32
	// 루프를 돌면서 최대의 최소 가격을 계속 갱신한다.
	// 갱신할 수 없으면 현재 최소 가격에서 빼고 maxProfit이 갱신 가능하면 갱신한다.

	for _, p := range prices {
		if p < minPrice {
			minPrice = p
		} else {
			maxProfit = max(maxProfit, p-minPrice)
		}
	}

	return maxProfit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(BuyAndSellStock(prices))
}
