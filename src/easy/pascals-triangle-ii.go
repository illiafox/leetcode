package easy

// https://leetcode.com/problems/pascals-triangle-ii/
func getRow(numRows int) []int {
	res := make([]int, numRows+1)

	for i := 0; i <= numRows; i++ {
		res[i] = 1
		for j := i - 1; j > 0; j-- {
			res[j] = res[j] + res[j-1]
		}
	}

	return res
}
