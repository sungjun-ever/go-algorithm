package main

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))
	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			n := len(stack)
			a, b := stack[n-1], stack[n-2]
			stack = stack[:n-2]

			var res int
			switch token {
			case "+":
				res = b + a
			case "-":
				res = b - a
			case "*":
				res = b * a
			case "/":
				res = b / a
			}

			stack = append(stack, res)

		default:
			val, _ := strconv.Atoi(token)
			stack = append(stack, val)
		}
	}

	return stack[0]
}
