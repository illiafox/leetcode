package medium

import "container/list"

type LRUCache struct {
	capacity int
	m        map[int]*list.Element
	l        *list.List
}

// LRUCacheConstructor implements solution for https://leetcode.com/problems/lru-cache/editorial/
func LRUCacheConstructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		m:        make(map[int]*list.Element),
		l:        list.New(),
	}
}

type cacheEntryLRU struct {
	key int
	val int
}

func (c *LRUCache) Get(key int) int {
	elem, exists := c.m[key]
	if !exists {
		return -1
	}

	c.l.MoveToFront(elem)
	return elem.Value.(cacheEntryLRU).val
}

func (c *LRUCache) sizeCheck() {
	for c.l.Len() > c.capacity {
		back := c.l.Back()
		c.l.Remove(back)
		delete(c.m, back.Value.(cacheEntryLRU).key)
	}
}

func (c *LRUCache) Put(key int, value int) {
	elem, exists := c.m[key]
	if exists {
		c.l.Remove(elem)
	}

	elem = c.l.PushFront(cacheEntryLRU{key, value})
	c.m[key] = elem
	c.sizeCheck()
}
