package medium

// validUtf8 returns whether the provided integer array is a valid UTF-8 encoding
// Source: https://leetcode.com/problems/utf-8-validation
// A character in UTF8 can be from 1 to 4 bytes long, subjected to the following rules:
//
//	 For a 1-byte character, the first bit is a 0, followed by its Unicode code.
//	 For an n-bytes character, the first n bits are all one's, the n + 1 bit is 0, followed by n - 1 bytes with the most significant 2 bits being 10.
//	--------------------+-----------------------------------------
//	  Number of Bytes   |        UTF-8 Octet Sequence (binary)
//	--------------------+-----------------------------------------
//	         1          |   0xxxxxxx
//	         2          |   110xxxxx 10xxxxxx
//	         3          |   1110xxxx 10xxxxxx 10xxxxxx
//	         4          |   11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
func validUtf8(data []int) bool {
	current := 0

	for _, d := range data {
		if d < 0 || d > 255 { // ensure input is valid
			return false
		}

		if current > 0 {
			const continuation = 0x2 // most significant 2 bits being 10
			if getFirstTwoBits(d) != continuation {
				return false
			}
			current--
		} else {
			if getFirstTwoBits(d) > 1 {
				l := getCharacterLength(d)
				if l <= 1 || l > 4 {
					return false
				}
				current = l - 1
			}
		}

	}

	return current == 0
}

func getFirstTwoBits(d int) int {
	return d >> 6
}

func getCharacterLength(d int) int {
	length := 0

	const shift = 1 << 7
	for d&shift != 0 {
		d <<= 1
		length += 1
	}

	return length
}
