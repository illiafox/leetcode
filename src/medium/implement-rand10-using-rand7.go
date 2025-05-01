package medium

import "math/rand"

// given by leetcode
func rand7() int {
	return rand.Int() % 7
}

func rand10() int {
	for {
		a := rand7()
		b := rand7()
		num := (a-1)*7 + b

		if num <= 40 {
			return 1 + (num-1)%10
		}
	}
}
