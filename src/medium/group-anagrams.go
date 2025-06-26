package medium

import (
	"bytes"
	"maps"
	"slices"
)

// https://leetcode.com/problems/group-anagrams
func groupAnagrams(strs []string) [][]string {
	similar := make(map[string][]string)
	var buf bytes.Buffer

	for _, s := range strs {
		buf.Reset()
		buf.WriteString(s)
		slices.Sort(buf.Bytes())
		// user sorted string as hash
		hash := buf.String()

		similar[hash] = append(similar[hash], s)
	}

	return slices.Collect(maps.Values(similar))
}
