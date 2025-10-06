package easy

// https://leetcode.com/problems/number-of-segments-in-a-string/
func countSegments(s string) (parts int) {
	present := false

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if present {
				parts++
				present = false
			}
		} else {
			present = true
		}
	}

	if present {
		parts++
	}

	return parts
}
