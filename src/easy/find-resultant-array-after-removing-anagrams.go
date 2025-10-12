package easy

// https://leetcode.com/problems/find-resultant-array-after-removing-anagrams
func removeAnagrams(words []string) (out []string) {
	for i := 0; i < len(words); {
		word := words[i]

		var original [26]byte
		for j := 0; j < len(word); j++ {
			original[word[j]-'a']++
		}

		for i++; i < len(words); i++ {
			next := words[i]
			if next == word {
				continue
			}

			var cmp [26]byte
			for k := 0; k < len(next); k++ {
				cmp[next[k]-'a']++
			}

			if original != cmp {
				break
			}
		}

		out = append(out, word)
	}

	return out
}
