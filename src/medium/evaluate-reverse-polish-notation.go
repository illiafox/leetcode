package medium

import (
	"slices"
	"strconv"
)

// TODO: rewrite in rust
func evalRPN(tokens []string) int {
	var stack []any

	for _, token := range tokens {
		if slices.Contains([]string{"+", "-", "*", "/"}, token) {
			stack = append(stack, token)
		} else {
			num, err := strconv.Atoi(token)
			if err != nil {
				panic(err)
			}

			stack = append(stack, num)
		}

		for len(stack) > 2 {
			a := stack[len(stack)-3].(int)
			b := stack[len(stack)-2].(int)
			op, ok := stack[len(stack)-1].(string)
			if !ok {
				break
			}

			var res int
			switch op {
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			case "/":
				res = a / b
			}

			stack[len(stack)-3] = res
			stack = stack[:len(stack)-2]
		}
	}

	return stack[0].(int)
}
