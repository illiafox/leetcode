package hard

type CountNode struct {
	Count int
	Keys  map[string]struct{}

	Prev *CountNode
	Next *CountNode
}

type AllOne struct {
	keyToNode map[string]*CountNode
	head      *CountNode // min count
	tail      *CountNode // max count
}

func Constructor() AllOne {
	return AllOne{
		keyToNode: make(map[string]*CountNode),
		tail:      nil,
	}
}

func (a *AllOne) Inc(key string) {
	node, exists := a.keyToNode[key]
	if !exists {
		// New key
		if a.head != nil && a.head.Count == 1 {
			a.head.Keys[key] = struct{}{}
			a.keyToNode[key] = a.head
			return
		}
		// Create new head node with count = 1
		newNode := &CountNode{
			Count: 1,
			Keys:  map[string]struct{}{key: {}},
			Next:  a.head,
		}
		if a.head != nil {
			a.head.Prev = newNode
		}
		a.head = newNode
		if a.tail == nil {
			a.tail = newNode
		}
		a.keyToNode[key] = newNode
		return
	}

	// Key already exists
	next := node.Next
	var newNode *CountNode
	if next != nil && next.Count == node.Count+1 {
		// Move to existing next node
		newNode = next
	} else {
		// Create new node with count +1
		newNode = &CountNode{
			Count: node.Count + 1,
			Keys:  make(map[string]struct{}),
			Prev:  node,
			Next:  node.Next,
		}
		if node.Next != nil {
			node.Next.Prev = newNode
		}
		node.Next = newNode
		if a.tail == node {
			a.tail = newNode
		}
	}
	newNode.Keys[key] = struct{}{}
	a.keyToNode[key] = newNode

	// Remove key from old node
	delete(node.Keys, key)
	if len(node.Keys) == 0 {
		// Remove node from list
		if node.Prev != nil {
			node.Prev.Next = node.Next
		} else {
			a.head = node.Next
		}
		if node.Next != nil {
			node.Next.Prev = node.Prev
		} else {
			a.tail = node.Prev
		}
	}
}

func (a *AllOne) Dec(key string) {
	node, exists := a.keyToNode[key]
	if !exists {
		return // shouldn't happen per problem guarantees
	}

	if node.Count == 1 {
		// Remove key entirely
		delete(node.Keys, key)
		delete(a.keyToNode, key)

		if len(node.Keys) == 0 {
			// Remove node from list
			if node.Prev != nil {
				node.Prev.Next = node.Next
			} else {
				a.head = node.Next
			}
			if node.Next != nil {
				node.Next.Prev = node.Prev
			} else {
				a.tail = node.Prev
			}
		}
		return
	}

	// Move key to previous node (count - 1)
	prev := node.Prev
	var newNode *CountNode
	if prev != nil && prev.Count == node.Count-1 {
		newNode = prev
	} else {
		newNode = &CountNode{
			Count: node.Count - 1,
			Keys:  make(map[string]struct{}),
			Prev:  node.Prev,
			Next:  node,
		}
		if node.Prev != nil {
			node.Prev.Next = newNode
		} else {
			a.head = newNode
		}
		node.Prev = newNode
	}

	newNode.Keys[key] = struct{}{}
	a.keyToNode[key] = newNode

	// Remove key from current node
	delete(node.Keys, key)
	if len(node.Keys) == 0 {
		// Remove node from list
		if node.Prev != nil {
			node.Prev.Next = node.Next
		} else {
			a.head = node.Next
		}
		if node.Next != nil {
			node.Next.Prev = node.Prev
		} else {
			a.tail = node.Prev
		}
	}
}

func (a *AllOne) GetMaxKey() string {
	if a.tail != nil {
		for k := range a.tail.Keys {
			return k
		}
	}
	return ""
}

func (a *AllOne) GetMinKey() string {
	if a.head != nil {
		for k := range a.head.Keys {
			return k
		}
	}
	return ""
}
