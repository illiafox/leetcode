package medium

func lengthOfLongestSubstring(s string) int {
	maxLength := -1

	dupl := make(map[byte]struct{})
	start := 0

	for end := 0; end < len(s); end++ {
		for _, ok := dupl[s[end]]; ok; _, ok = dupl[s[end]] {
			delete(dupl, s[start])
			start++
		}
		dupl[s[end]] = struct{}{}
		maxLength = max(maxLength, end-start+1)
	}

	return max(maxLength, len(dupl))
}
