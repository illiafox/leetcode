package medium

type magickDictTrie struct {
	isWord bool
	nodes  [26]*magickDictTrie
}

type MagicDictionary struct {
	trie *magickDictTrie
}

func (dict *MagicDictionary) BuildDict(dictionary []string) {
	for _, word := range dictionary {
		cur := dict.trie
		for i := range word {
			l := word[i] - 'a'

			entry := cur.nodes[l]
			if entry == nil {
				entry = new(magickDictTrie)
				cur.nodes[l] = entry
			}

			cur = entry
		}

		cur.isWord = true
	}
}

func (dict *MagicDictionary) traverse(t *magickDictTrie, word string, skipped bool) bool {
	if len(word) == 0 {
		return skipped && t.isWord
	}

	l := word[0] - 'a'
	foundNode := t.nodes[l]
	word = word[1:]

	if !skipped {
		for i, node := range t.nodes {
			if byte(i) == l {
				continue
			}
			if node != nil && dict.traverse(node, word, true) {
				return true
			}
		}
	}
	if foundNode == nil {
		return false
	}

	return dict.traverse(foundNode, word, skipped)
}

func MagicDictionaryConstructor() MagicDictionary {
	return MagicDictionary{new(magickDictTrie)}
}

// https://leetcode.com/problems/implement-magic-dictionary/description/
func (dict *MagicDictionary) Search(searchWord string) bool {
	return dict.traverse(dict.trie, searchWord, false)
}
