package hard

import (
	"maps"
	"strings"
)

// https://leetcode.com/problems/substring-with-concatenation-of-all-words
func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return nil
	}

	wordLength := len(words[0])
	wordsCount := len(words)
	totalLen := wordLength * wordsCount
	if len(s) < totalLen {
		return nil
	}

	need := make(map[string]int, wordsCount)
	for _, w := range words {
		need[w]++
	}

	var res []int

	for offset := 0; offset < wordLength; offset++ {
		left := offset
		matchedCount := 0
		seen := make(map[string]int)

		for right := offset; right+wordLength <= len(s); right += wordLength {
			w := s[right : right+wordLength]

			if _, ok := need[w]; !ok {
				for k := range seen {
					delete(seen, k)
				}
				matchedCount = 0
				left = right + wordLength
				continue
			}

			seen[w]++
			matchedCount++

			for seen[w] > need[w] {
				lw := s[left : left+wordLength]
				seen[lw]--
				matchedCount--
				left += wordLength
			}

			if matchedCount == wordsCount {
				res = append(res, left)

				sub := s[left : left+wordLength]
				seen[sub]--
				matchedCount--
				left += wordLength
			}
		}
	}

	return res
}

// covers case when strings of words variable are not of the same length
func findSubstringOld(s string, words []string) []int {
	totalLen := 0
	for _, w := range words {
		totalLen += len(w)
	}

	wordsM := make(map[string]int)
	for _, w := range words {
		wordsM[w]++
	}

	var res []int

	for i := 0; i <= len(s)-totalLen; i++ {
		requiredWords := maps.Clone(wordsM)

		sub := s[i : i+totalLen]

		if len(requiredWords) == 1 {
			if strings.Count(sub, words[0])*len(words[0]) == totalLen {
				res = append(res, i)
			}
			continue
		}

		for len(requiredWords) > 0 {

			found := false
			for word := range requiredWords {
				if strings.HasPrefix(sub, word) {
					requiredWords[word]--
					if requiredWords[word] == 0 {
						delete(requiredWords, word)
					}
					sub = sub[len(word):]
					found = true
				}
			}

			if !found {
				break
			}

			if len(requiredWords) == 0 {
				res = append(res, i)
			}
		}
	}

	return res
}
