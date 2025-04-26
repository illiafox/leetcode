package medium

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))

	type s struct {
		val   int
		count int
	}

	stack := []s{{
		val:   temperatures[0],
		count: 1,
	}}

	for i := 1; i < len(temperatures); i++ {
		if stack[len(stack)-1].val == temperatures[i] {
			stack[len(stack)-1].count++
			continue
		}
		stack = append(stack, s{temperatures[i], 1})
	}

	nextIdx := 0

	for len(stack) > 0 {
		first := stack[0]

		for first.count > 0 {
			steps := first.count

			found := false
			for j := 1; j < len(stack); j++ {
				if stack[j].val > first.val {
					result[nextIdx] = steps
					nextIdx++
					first.count--
					found = true
					break
				}
				steps += stack[j].count
			}

			if !found {
				break
			}
		}

		nextIdx += first.count
		stack = stack[1:]
	}

	return result
}

func dailyTemperaturesNaive(temperatures []int) []int {
	result := make([]int, len(temperatures))

	for i := len(temperatures) - 1; i >= 0; i-- {
		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[j] > temperatures[i] {
				result[i] = j - i
				break
			}
		}
	}

	return result
}
