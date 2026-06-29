package easy

import "strings"

// https://leetcode.com/problems/number-of-strings-that-appear-as-substrings-in-word
// Given an array of strings patterns and a string word, return the number of strings in patterns that exist as a substring in word.
// A substring is a contiguous sequence of characters within a string.
// 1 <= patterns.length <= 100
// 1 <= patterns[i].length <= 100
// 1 <= word.length <= 100
// patterns[i] and word consist of lowercase English letters.

// // 1. Naive strings.Contains
func numOfStrings(patterns []string, word string) (out int) {
	for i := range patterns {
		if strings.Contains(word, patterns[i]) {
			out++
		}
	}
	return out
}

// // 2. Custom all in one loop
func numOfStringsCustom(patterns []string, word string) (out int) {
	seen := make([]int8, len(patterns))

	for i := range word {
		letter := word[i]

		for j := range seen {
			s := seen[j]
			if s < 0 {
				continue
			}

			pattern := patterns[j]
			if letter == pattern[s] {
				s += 1
				if s == int8(len(pattern)) {
					seen[j] = -1
					out++
				} else {
					seen[j] = s
				}

				continue
			}

			start := int(s) - 1
			s = 0

			for ; start >= 0; start-- {
				if word[i-start] == pattern[s] {
					s++
				} else {
					s = 0
				}
			}

			seen[j] = s
		}
	}

	return out
}

// // 3. Aho-Corasick state machine
// https://en.wikipedia.org/wiki/Aho%E2%80%93Corasick_algorithm

func numOfStringsAhoCorasick(patterns []string, word string) (out int) {
	trie := BuildAC(patterns)
	return trie.Search(word, len(patterns))
}

type ACNode struct {
	// [256] array is much faster than map[byte]*ACNode for byte lookups.
	children [256]*ACNode
	fail     *ACNode
	// We store the indices of the patterns that end at this node.
	matches []int
}

func BuildAC(patterns []string) *ACNode {
	root := &ACNode{}

	// build the standard Trie
	for i, p := range patterns {
		curr := root
		for j := 0; j < len(p); j++ {
			c := p[j]
			if curr.children[c] == nil {
				curr.children[c] = &ACNode{}
			}
			curr = curr.children[c]
		}
		// mark the end of a pattern by storing its index
		curr.matches = append(curr.matches, i)
	}

	// build the failure links using BFS
	queue := []*ACNode{}

	// initialize the first layer to fail back to the root
	for _, child := range root.children {
		if child != nil {
			child.fail = root
			queue = append(queue, child)
		}
	}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for i, child := range curr.children {
			if child == nil {
				continue
			}

			// traverse failure links to find the longest valid suffix
			failNode := curr.fail
			for failNode != nil && failNode.children[i] == nil {
				failNode = failNode.fail
			}

			if failNode == nil {
				child.fail = root
			} else {
				child.fail = failNode.children[i]
				// merge matches from the failure link.
				child.matches = append(child.matches, child.fail.matches...)
			}

			queue = append(queue, child)
		}
	}

	return root
}

func (root *ACNode) Search(word string, numPatterns int) int {
	curr := root

	// ceiling division (numPatterns + 7) / 8
	bitsetSize := (numPatterns + 7) >> 3
	seen := make([]byte, bitsetSize)
	count := 0

	for i := 0; i < len(word); i++ {
		c := word[i]

		for curr != root && curr.children[c] == nil {
			curr = curr.fail
		}
		if curr.children[c] != nil {
			curr = curr.children[c]
		} else {
			curr = root
		}

		for _, matchIdx := range curr.matches {
			byteIdx := matchIdx >> 3             // matchIdx / 8
			bitMask := byte(1 << (matchIdx & 7)) // matchIdx % 8

			if seen[byteIdx]&bitMask == 0 {
				seen[byteIdx] |= bitMask
				count++
			}
		}

		if count == numPatterns {
			break
		}
	}

	return count
}
