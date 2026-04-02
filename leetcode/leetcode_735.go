package main

import "math"

// 내가 푼 코드
func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0, len(asteroids))

	for i := 0; i < len(asteroids); i++ {
		curr := asteroids[i]
		if len(stack) == 0 {
			stack = append(stack, curr)
		} else {
			currCrashed := false
			top := stack[len(stack)-1]

			for top > 0 && curr < 0 && len(stack) > 0 {
				if math.Abs(float64(top)) == math.Abs(float64(curr)) {
					stack = stack[:len(stack)-1]
					currCrashed = true
					break
				} else if math.Abs(float64(top)) < math.Abs(float64(curr)) {
					stack = stack[:len(stack)-1]
				} else {
					currCrashed = true
					break
				}

				if len(stack) > 0 {
					top = stack[len(stack)-1]
				}

			}

			if !currCrashed {
				stack = append(stack, curr)
			}
		}
	}

	return stack
}

// ai가 수정한 코드
func asteroidCollision2(asteroids []int) []int {
	// 결과 스택 (미리 용량 할당)
	stack := make([]int, 0, len(asteroids))

	for _, ast := range asteroids {
		alive := true // 현재 행성의 생존 여부

		// 충돌 조건:
		// 1. 현재 행성이 살아있고(alive)
		// 2. 왼쪽으로 가고 있으며(ast < 0)
		// 3. 스택의 Top이 오른쪽으로 가고 있을 때(stack[top] > 0)
		for alive && len(stack) > 0 && stack[len(stack)-1] > 0 && ast < 0 {
			top := stack[len(stack)-1]

			if top < -ast {
				// 내(ast)가 더 커서, Top을 파괴함 -> 스택에서 제거하고 계속 진행
				stack = stack[:len(stack)-1]
				continue
			} else if top == -ast {
				// 크기가 같음 -> 둘 다 파괴됨
				stack = stack[:len(stack)-1]
				alive = false
			} else {
				// Top이 더 큼 -> 내가 파괴됨
				alive = false
			}
		}

		// 충돌 루프가 끝난 후에도 살아있다면 스택에 추가
		if alive {
			stack = append(stack, ast)
		}
	}

	return stack
}
