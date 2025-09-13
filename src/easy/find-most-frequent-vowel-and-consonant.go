package easy

// https://leetcode.com/problems/find-most-frequent-vowel-and-consonant
func maxFreqSum(s string) int {
	// 1 <= s.length <= 100
	// s consists of lowercase English letters only.
	var letters [26]byte

	for i := 0; i < len(s); i++ {
		letters[s[i]-'a']++
	}

	var maxVowel byte
	var maxLetter byte

	// a=0, e=4, i=8, o=14, u=20
	const vowelMask = (1 << 0) | (1 << 4) | (1 << 8) | (1 << 14) | (1 << 20)

	for i := 0; i < len(letters); i++ {
		if (vowelMask & (1 << i)) != 0 { // mask with only the i-th bit set
			maxVowel = max(maxVowel, letters[i])
		} else {
			maxLetter = max(maxLetter, letters[i])
		}
	}

	return int(maxVowel + maxLetter)
}
