package main

import (
	"math"
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)

	n := len(potions)
	ans := make([]int, len(spells))

	for i := 0; i < len(spells); i++ {
		left, right := 0, n-1
		minLeft, maxRight := math.MaxInt32, right
		currSpell := spells[i]

		lS := int64(currSpell) * int64(potions[left])
		rS := int64(currSpell) * int64(potions[right])

		// 가장 끝이 제일 작으면 다음으로
		if rS < success {
			continue
		}

		// 둘 다 크면 다음으로
		if lS >= success && rS >= success {
			ans[i] = right - left + 1
			continue
		}

		for left <= right {
			mid := left + (right-left)/2
			mS := int64(currSpell) * int64(potions[mid])

			// 미드 >= success  앞쪽 탐색
			// 미드 < success 작으면 뒤쪽 탐색
			if mS >= success {
				minLeft = mid
				right = mid - 1
			} else {
				left = mid + 1
			}
		}

		if minLeft < maxRight {
			ans[i] = maxRight - minLeft + 1
		} else {
			ans[i] = 1
		}

	}

	return ans
}

// ai 리팩토링 코드
func successfulPairs2(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)

	m := len(potions)

	ans := make([]int, len(spells))

	for i, spell := range spells {
		left, right := 0, m-1

		// 🌟 핵심: 조건을 만족하는 첫 번째 포션의 인덱스.
		// 기본값을 m으로 두어, 아무것도 조건을 만족하지 못하면 0개가 나오도록 유도
		firstValidIdx := m

		for left <= right {
			mid := left + (right-left)/2

			if int64(spell)*int64(potions[mid]) >= success {
				firstValidIdx = mid // 일단 기록! (더 작은 인덱스도 되는지 보러 감)
				right = mid - 1     // 왼쪽 절반 탐색
			} else {
				left = mid + 1 // 오른쪽 절반 탐색
			}
		}

		// 전체 포션 개수 - 처음 성공한 포션의 인덱스 = 성공한 포션의 총 개수
		ans[i] = m - firstValidIdx
	}

	return ans
}
