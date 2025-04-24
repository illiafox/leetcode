package easy

func isPowerOfTwo(n int) bool {
	for n > 0 {
		if n&1 == 1 {
			return n == 1
		}
		n >>= 1
	}

	return false
}
