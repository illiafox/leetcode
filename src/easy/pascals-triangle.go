package easy

// https://leetcode.com/problems/pascals-triangle/
func generatePascalTriangle(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}

	out := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		res := make([]int, i+1)
		res[0], res[i] = 1, 1
		for j := 1; j < i; j++ {
			res[j] = out[i-1][j] + out[i-1][j-1]
		}
		out[i] = res
	}

	return out
}
