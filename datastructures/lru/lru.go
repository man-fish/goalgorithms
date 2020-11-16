package lru

import (
	"container/list"
)

// Cache is a LRU http cache which is not safe for concurrent access
type Cache struct {
	// maxBytes is max memory allow to use, 0 for all
	maxBytes int64
	// nBytes is memory which has been used
	nBytes int64
	// ll is LRU List
	ll *list.List
	// cache is a fuck unconcurrent map
	cache map[string]*list.Element
	// OnEvicted is a callback happens when a record is evicted
	OnEvicted EvictHandler
}

type entry struct {
	key   string
	value Value
}

// Value use to count how many bytes this record counts
type Value interface {
	Len() int
}

// EvictHandler func is the callback for evict
type EvictHandler func(key string, value Value)

// New is the construct function for gc
func New(mBytes int64, onEvicted EvictHandler) *Cache {
	return &Cache{
		maxBytes:  mBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

// Get use to get a record from cache
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

// RemoveOldest use to delete a latest recently used record from cache
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nBytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

// Add use to add a record to cache
func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nBytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nBytes += int64(len(key)) + int64(value.Len())
	}
	for c.maxBytes != 0 && c.nBytes > c.maxBytes {
		c.RemoveOldest()
	}
}

// Len returns kv nums
func (c *Cache) Len() int {
	return c.ll.Len()
}
