package medium

// https://leetcode.com/problems/shifting-letters-ii
func shiftingLetters(s string, shifts [][]int) string {
	b := []byte(s)

	diffArray := make([]int, len(s))

	for _, shift := range shifts {
		change := -1 // shift backward
		if shift[2] == 1 {
			change = 1 // shift forward
		}

		diffArray[shift[0]] += change
		if shift[1]+1 < len(s) {
			diffArray[shift[1]+1] -= change
		}
	}

	currentShifts := 0

	for i, change := range diffArray {
		currentShifts = (currentShifts + change) % 26
		if currentShifts < 0 {
			currentShifts += 26
		}

		// does not work when wrapping from 'a' to 'z'?
		b[i] = 'a' + byte(int(b[i])-'a'+currentShifts)%26
	}

	return string(b)
}
