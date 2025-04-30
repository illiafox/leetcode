package medium

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	top, bottom := 0, n-1
	left, right := 0, n-1
	nextValue := 1

	for nextValue <= n*n {
		for i := left; i <= right; i++ {
			matrix[top][i] = nextValue
			nextValue++
		}
		top++

		for i := top; i <= bottom; i++ {
			matrix[i][right] = nextValue
			nextValue++
		}
		right--

		for i := right; i >= left; i-- {
			matrix[bottom][i] = nextValue
			nextValue++
		}
		bottom--

		for i := bottom; i >= top; i-- {
			matrix[i][left] = nextValue
			nextValue++
		}
		left++
	}

	return matrix
}
