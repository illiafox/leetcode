package easy

func findPoisonedDuration(timeSeries []int, duration int) int {
	total := 0
	poisonedUntil := 0

	for _, t := range timeSeries {
		if t < poisonedUntil {
			total -= poisonedUntil - t
		}
		total += duration
		poisonedUntil = t + duration
	}

	return total
}
