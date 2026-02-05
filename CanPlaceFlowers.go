package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	length := len(flowerbed)
	// 현재 기준으로 앞뒤가 0이어야한다
	for i := 0; i < length; i++ {
		if flowerbed[i] == 0 {
			prevEmpty := i == 0 || flowerbed[i-1] == 0
			nextEmpty := i == length-1 || flowerbed[i+1] == 0

			if prevEmpty && nextEmpty {
				flowerbed[i] = 1
				n--
			}

			if n == 0 {
				return true
			}
		}
	}

	return false
}
