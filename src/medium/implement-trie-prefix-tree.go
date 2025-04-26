package medium

type Trie struct {
	count int
	nodes [26]*Trie
}

func TrieConstructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	cur := t
	for i := range word {
		l := word[i] - 'a'

		entry := cur.nodes[l]
		if entry == nil {
			entry = new(Trie)
			cur.nodes[l] = entry
		}

		cur = entry
	}

	cur.count++
}

func (t *Trie) traverse(word string) *Trie {
	cur := t
	for i := 0; i < len(word) && cur != nil; i++ {
		l := word[i] - 'a'
		cur = cur.nodes[l]
	}

	return cur
}

func (t *Trie) Search(word string) bool {
	cur := t.traverse(word)
	return cur != nil && cur.count > 0
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.traverse(prefix) != nil
}
