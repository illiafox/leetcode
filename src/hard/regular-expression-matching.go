package hard

// https://leetcode.com/problems/regular-expression-matching/
func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	firstMatch := len(s) > 0 && (s[0] == p[0] || p[0] == '.')

	if len(p) >= 2 && p[1] == '*' {
		// try zero occurrence OR one/more if first character matches
		return isMatch(s, p[2:]) || (firstMatch && isMatch(s[1:], p))
	} else {
		// just skip one letter
		return firstMatch && isMatch(s[1:], p[1:])
	}
}
