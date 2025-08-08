package hard

func countDigitOne(n int) int {
	count := 0
	factor := 1

	for factor <= n {
		higher := n / (factor * 10)
		current := (n / factor) % 10
		lower := n % factor

		switch current {
		case 0:
			count += higher * factor
		case 1:
			count += higher*factor + lower + 1
		default:
			count += (higher + 1) * factor
		}

		factor *= 10
	}

	return count
}
