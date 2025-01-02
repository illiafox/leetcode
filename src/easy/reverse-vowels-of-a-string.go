package easy

// https://leetcode.com/problems/reverse-vowels-of-a-string/
func reverseVowels(s string) string {
	str := []byte(s)

	isVowel := func(c byte) bool {
		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
		}

		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			return true
		}

		return false
	}

	left, right := 0, len(str)-1
	for left < right {
		for left < len(str) && !isVowel(str[left]) {
			left++
		}
		for right >= 0 && !isVowel(str[right]) {
			right--
		}

		if left >= right {
			break
		}

		str[left], str[right] = str[right], str[left]
		left++
		right--
	}

	return string(str)
}
