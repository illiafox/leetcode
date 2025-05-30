package easy

// https://leetcode.com/problems/excel-sheet-column-number/
func titleToNumber(columnTitle string) (out int) {
	x := 1

	for i := len(columnTitle) - 1; i >= 0; i-- {
		out += x * int(columnTitle[i]-'A')
		x *= 26
	}

	return out
}
