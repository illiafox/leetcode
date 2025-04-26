package easy

func countBits(n int) []int {
	out := make([]int, n+1)
	for i := range out {
		out[i] = out[i>>1] + (i & 1)
	}

	return out
}
