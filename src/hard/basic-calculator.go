package hard

// https://leetcode.com/problems/basic-calculator/
// final version
func calculate(s string) int {
	var parse func() int
	i := 0

	parse = func() int {
		sum := 0
		sign := 1

		for i < len(s) {
			switch s[i] {
			case ' ':
				i++
			case '+':
				sign = 1
				i++
			case '-':
				sign = -1
				i++
			case '(':
				i++
				val := parse()
				sum += sign * val
			case ')':
				i++
				return sum
			default:
				num := 0
				for i < len(s) && s[i] >= '0' && s[i] <= '9' {
					num = num*10 + int(s[i]-'0')
					i++
				}
				sum += sign * num
			}
		}
		return sum
	}

	return parse()
}

// first version
func calculateOld(s string) int {
	sum := 0
	sign := 1

	for i := 0; i < len(s); {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = 1
			i++
		case '-':
			sign = -1
			i++
		case '(':
			openCount := 1
			closedCount := 0

			j := i + 1
			for ; closedCount != openCount; j++ {
				switch s[j] {
				case '(':
					openCount++
				case ')':
					closedCount++
				}
			}
			expr := calculateOld(s[i+1 : j-1])
			sum += expr * sign
			i = j
		default:
			expr := 0
			for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
				expr = expr*10 + int(s[i]-'0')
			}
			sum += expr * sign
		}
	}

	return sum
}
