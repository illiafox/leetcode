package medium

type SearchTrie struct {
	nodes [26]*SearchTrie
	word  *string
}

func (t *SearchTrie) Insert(word string) {
	cur := t
	for i := range word {
		l := word[i] - 'a'

		entry := cur.nodes[l]
		if entry == nil {
			entry = new(SearchTrie)
			cur.nodes[l] = entry
		}

		cur = entry
	}

	cur.word = &word
}

func suggestedProducts(products []string, searchWord string) [][]string {
	var trie SearchTrie
	for _, p := range products {
		trie.Insert(p)
	}

	out := make([][]string, len(searchWord))

	start := &trie
	for i := 0; i < len(searchWord); i++ {
		start = start.nodes[searchWord[i]-'a']
		if start == nil {
			return out
		}

		var res []string

		var traverse func(node SearchTrie) bool
		traverse = func(node SearchTrie) bool {
			if node.word != nil {
				res = append(res, *node.word)
				if len(res) == 3 {
					return true
				}
			}
			for _, n := range node.nodes {
				if n != nil && traverse(*n) {
					return true
				}
			}

			return false
		}

		traverse(*start)
		out[i] = res
	}

	return out
}
