package easy

func isValidParentheses(s string) bool {
	var queue []byte

	c := map[byte]byte{
		')': '(', ']': '[', '}': '{',
	}

	for i := range s {
		switch s[i] {
		case '(', '[', '{':
			queue = append(queue, s[i])
		case ')', ']', '}':
			if len(queue) == 0 || queue[len(queue)-1] != c[s[i]] {
				return false
			}
			queue = queue[:len(queue)-1]
		}
	}

	return len(queue) == 0
}
