package medium

// https://leetcode.com/problems/number-of-substrings-containing-all-three-characters
// Given a string s consisting only of characters a, b and c.
// Return the number of substrings containing at least one occurrence of all these characters a, b and c.
//
// 3 <= s.length <= 5 x 10^4
// s only consists of a, b or c characters.
func numberOfSubstrings(s string) (total int) {
	var charCount [3]int
	left := 0

	for right := 0; right < len(s); right++ {
		charCount[s[right]-'a']++

		for charCount[0] >= 1 && charCount[1] >= 1 && charCount[2] >= 1 {
			total += len(s) - right
			charCount[s[left]-'a']--
			left++
		}
	}

	return total
}
