package easy

func isUgly(n int) bool {
	if n == 0 {
		return false
	}

	for {
		switch {
		case n%5 == 0:
			n /= 5
		case n%3 == 0:
			n /= 3
		case n%2 == 0:
			n /= 2
		default:
			return n == 1
		}
	}
}
