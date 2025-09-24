package easy

// https://leetcode.com/problems/find-n-unique-integers-sum-up-to-zero
func sumZero(n int) []int {
	if n <= 1 {
		return []int{0}
	}

	arr := make([]int, n)

	i := 0
	for div := 1; div <= n/2; div++ {
		arr[i] = div
		arr[i+1] = -div
		i += 2
	}

	return arr
}
