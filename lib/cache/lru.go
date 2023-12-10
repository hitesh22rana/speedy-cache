package cache

import (
	"container/list"
	"sync"
)

type Pair struct {
	key   string
	value interface{}
}

type LRUCache struct {
	capacity int
	list     *list.List
	exists   map[string]*list.Element
	mu       sync.Mutex
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		list:     list.New(),
		exists:   make(map[string]*list.Element),
		mu:       sync.Mutex{},
	}
}

func (c *LRUCache) Get(key string) (interface{}, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if node, ok := c.exists[key]; ok {
		val := node.Value.(*list.Element).Value.(Pair).value
		c.list.MoveToFront(node)
		return val, nil
	}

	return nil, ErrKeyNotFound
}

func (c *LRUCache) Set(key string, value interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if node, ok := c.exists[key]; ok {
		c.list.MoveToFront(node)
		node.Value.(*list.Element).Value = Pair{key: key, value: value}
		return nil
	}

	if c.list.Len() == c.capacity {
		idx := c.list.Back().Value.(*list.Element).Value.(Pair).key
		delete(c.exists, idx)
		c.list.Remove(c.list.Back())
	}

	node := &list.Element{
		Value: Pair{
			key:   key,
			value: value,
		},
	}

	ptr := c.list.PushFront(node)
	c.exists[key] = ptr

	return nil
}

func (c *LRUCache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if node, ok := c.exists[key]; ok {
		delete(c.exists, key)
		c.list.Remove(node)
		return nil
	}

	return ErrKeyNotFound
}
