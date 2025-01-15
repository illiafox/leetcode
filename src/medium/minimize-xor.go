package medium

import "math/bits"

func minimizeXor(num1 int, num2 int) int {
	num2BitsCount := bits.OnesCount64(uint64(num2))

	res := num1
	resBitsCount := bits.OnesCount64(uint64(res))

	for bit := 0; resBitsCount != num2BitsCount; bit++ {
		if resBitsCount < num2BitsCount {
			if res&(1<<bit) == 0 {
				res |= 1 << bit
				resBitsCount++
			}
		} else {
			if res&(1<<bit) != 0 {
				res &= ^(1 << bit)
				resBitsCount--
			}
		}
	}

	return res
}
