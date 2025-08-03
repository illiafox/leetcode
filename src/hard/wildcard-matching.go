package hard

// // https://leetcode.com/problems/wildcard-matching/
func wildcardMatchMemoization(s string, p string) bool {
	memo := make(map[[2]int]bool) // [strPos, patternPos]

	var dfs func(strPos, patternPos int) bool
	dfs = func(strPos, patternPos int) bool {
		key := [2]int{strPos, patternPos}
		if val, ok := memo[key]; ok {
			return val
		}

		if patternPos == len(p) {
			return strPos == len(s)
		}

		if strPos == len(s) {
			// only match if p is all '*'
			for k := patternPos; k < len(p); k++ {
				if p[k] != '*' {
					memo[key] = false
					return false
				}
			}
			memo[key] = true
			return true
		}

		if p[patternPos] == '*' {
			// match zero characters or one character
			// same as isMatch(s, p[1:]) || isMatch(s[1:], p)
			memo[key] = dfs(strPos, patternPos+1) || dfs(strPos+1, patternPos)
		} else {
			firstMatch := s[strPos] == p[patternPos] || p[patternPos] == '?'
			memo[key] = firstMatch && dfs(strPos+1, patternPos+1)
		}

		return memo[key]
	}

	return dfs(0, 0)
}

// https://leetcode.com/problems/wildcard-matching/
// slow without memoization
func wildcardMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	if len(s) == 0 {
		// only match if p is all '*'
		for i := 0; i < len(p); i++ {
			if p[i] != '*' {
				return false
			}
		}
		return true
	}

	if p[0] == '*' {
		// match zero characters or one character
		return wildcardMatch(s, p[1:]) || wildcardMatch(s[1:], p)
	}

	firstMatch := s[0] == p[0] || p[0] == '?'

	return firstMatch && wildcardMatch(s[1:], p[1:])
}
