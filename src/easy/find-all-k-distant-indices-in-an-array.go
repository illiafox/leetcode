package easy

func findKDistantIndices(nums []int, key, k int) []int {
	n := len(nums)
	mark := make([]bool, n)

	// mark every index that lies in a key-window
	for i, v := range nums {
		if v == key {
			left, right := max(0, i-k), min(n-1, i+k)
			for j := left; j <= right; j++ {
				mark[j] = true
			}
		}
	}

	// collect marked indices in ascending order
	out := make([]int, 0)
	for i, ok := range mark {
		if ok {
			out = append(out, i)
		}
	}
	return out
}
