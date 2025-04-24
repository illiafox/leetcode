package easy

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}

	return (n-1)&n == 0
}

func isPowerOfTwoAlternative(n int) bool {
	for n > 0 {
		if n&1 == 1 {
			return n == 1
		}
		n >>= 1
	}

	return false
}
