package easy

import "strings"

func repeatedSubstringPattern(s string) bool {
	// if s is a repeat of a substring, it will appear inside (s + s)[1:len(s)*2-1].
	return strings.Contains((s + s)[1:len(s)*2-1], s)
}

func repeatedSubstringPatternIneffective(s string) bool {
	for l := len(s) / 2; l >= 1; l -= 1 {
		if len(s)%l != 0 {
			continue
		}
		start := s[0:l]

		bad := false
		for i := l; i <= len(s)-l; i += l {
			if s[i:i+l] != start {
				bad = true
				break
			}
		}
		if !bad {
			return true
		}
	}

	return false
}
