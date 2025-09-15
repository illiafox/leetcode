package easy

// https://leetcode.com/problems/maximum-number-of-words-you-can-type
func canBeTypedWords(text string, brokenLetters string) (count int) {
	var mask int32
	for i := 0; i < len(brokenLetters); i++ {
		mask |= 1 << (brokenLetters[i] - 'a')
	}

	for i := 0; i < len(text); {
		valid := true
		for i < len(text) && text[i] != ' ' {
			if mask&(1<<(text[i]-'a')) != 0 {
				valid = false

				for i < len(text) && text[i] != ' ' {
					i++
				}

				break
			}
			i++
		}

		if valid {
			count++
		}

		for i < len(text) && text[i] == ' ' {
			i++
		}
	}

	return
}
