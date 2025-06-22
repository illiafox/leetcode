package medium

import "strings"

// https://leetcode.com/problems/sort-characters-by-frequency
func frequencySort(s string) string {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	buckets := make([][]byte, len(s)+1)
	for ch, freq := range m {
		buckets[freq] = append(buckets[freq], ch)
	}

	var builder strings.Builder
	for freq := len(buckets) - 1; freq > 0; freq-- {
		for _, ch := range buckets[freq] {
			for i := 0; i < freq; i++ {
				builder.WriteByte(ch)
			}
		}
	}

	return builder.String()
}

func frequencySortInefficient(s string) string {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	builder := strings.Builder{}

	for n := len(s); n >= 1; n-- {
		for k, v := range m {
			if v == n {
				delete(m, k)
				for i := 0; i < v; i++ {
					builder.WriteByte(k)
				}
			}
		}
	}

	return builder.String()
}
