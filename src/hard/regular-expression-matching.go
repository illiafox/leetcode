package hard

// https://leetcode.com/problems/regular-expression-matching/
func isMatchMemo(s string, p string) bool {
	memo := make(map[[2]int]bool)

	var dp func(strPos, patternPos int) bool
	dp = func(strPos, patternPos int) bool {
		key := [2]int{strPos, patternPos}
		if val, ok := memo[key]; ok {
			return val
		}

		if patternPos == len(p) {
			return strPos == len(s)
		}

		firstMatch := strPos < len(s) && (s[strPos] == p[patternPos] || p[patternPos] == '.')

		var result bool
		if patternPos+1 < len(p) && p[patternPos+1] == '*' {
			// Match zero or more of the preceding element
			result = dp(strPos, patternPos+2) || (firstMatch && dp(strPos+1, patternPos))
		} else {
			result = firstMatch && dp(strPos+1, patternPos+1)
		}

		memo[key] = result
		return result
	}

	return dp(0, 0)
}

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
