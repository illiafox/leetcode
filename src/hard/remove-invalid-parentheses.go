package hard

func tryParentheses(memo map[string]int, out *[]string, s string, removeCount int, minRemoval *int) {
	if removeCount > *minRemoval {
		return
	}

	if prev, ok := memo[s]; ok && removeCount >= prev {
		return
	}
	memo[s] = removeCount

	open := 0
	for i := range s {
		switch s[i] {
		case '(', ')':
			tryParentheses(memo, out, s[:i]+s[i+1:], removeCount+1, minRemoval)
		}
		switch s[i] {
		case '(':
			open++
		case ')':
			open--
			if open < 0 {
				return
			}
		}
	}

	if open != 0 {
		return
	}

	if *minRemoval > removeCount {
		*minRemoval = removeCount
		*out = (*out)[:0]
	}
	*out = append(*out, s)
}

// https://leetcode.com/problems/remove-invalid-parentheses
func removeInvalidParentheses(s string) []string {
	memo := make(map[string]int)
	curMinRemovals := len(s)

	var out []string
	tryParentheses(memo, &out, s, 0, &curMinRemovals)
	return out
}
