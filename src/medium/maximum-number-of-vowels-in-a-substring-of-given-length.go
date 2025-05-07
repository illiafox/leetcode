package medium

func maxVowels(s string, k int) int {
	isVowel := func(b byte) bool {
		switch b {
		case 'a', 'e', 'i', 'o', 'u':
			return true
		}
		return false
	}

	var vowelsCount int

	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			vowelsCount++
		}
	}

	var maxCount = vowelsCount
	for i := 0; i < len(s)-k; i++ {
		if isVowel(s[i]) {
			vowelsCount--
		}
		if isVowel(s[i+k]) {
			vowelsCount++
		}
		maxCount = max(maxCount, vowelsCount)
	}

	return maxCount
}
