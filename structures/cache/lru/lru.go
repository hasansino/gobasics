//
// Package lru is a very simple implementation of LRU Cache structure.
//
package lru

// Cache is LRU cache implementation
type Cache struct {
	size  int
	queue *queue
	data  map[string]*cacheEntry
}

type cacheEntry struct {
	value interface{}
	qItem *queueItem
}

// NewCache creates new instance of LRU cache
// with pre-initialized data structures to `size`
func NewCache(size int) *Cache {
	return &Cache{
		size:  size,
		queue: &queue{},
		data:  make(map[string]*cacheEntry, size),
	}
}

// Len returns size of cached data
func (c *Cache) Len() int {
	return len(c.data)
}

// Put a key-value pair into cache, it will update entry
// if it already exists
// This operation will make entry most recently used
func (c *Cache) Put(key string, value interface{}) {
	if _, exists := c.data[key]; exists {
		// update existing node and move qNode in front
		c.data[key].value = value
		c.queue.upfront(c.data[key].qItem)
	} else {
		if len(c.data) == c.size {
			// evict least used node from cache
			qItem := c.queue.tail
			delete(c.data, qItem.key)
			c.queue.evict(qItem)
		}
		// write new cache entry
		c.data[key] = &cacheEntry{
			value: value,
			qItem: c.queue.add(key),
		}
	}
}

// Get a value by key, return nil if value not found
// This operation will make entry most recently used
func (c *Cache) Get(key string) interface{} {
	if v, ok := c.data[key]; ok {
		c.queue.upfront(v.qItem)
		return v.value
	}
	return nil
}

type queue struct {
	tail *queueItem
	head *queueItem
}

type queueItem struct {
	next *queueItem
	prev *queueItem
	key  string
}

// add entry in front of queue
func (q *queue) add(key string) *queueItem {
	newNode := &queueItem{key: key}
	if q.head == nil {
		// first entry
		q.head, q.tail = newNode, newNode
	} else {
		// new head
		newNode.prev = q.head
		q.head.next = newNode
		q.head = newNode
	}
	return newNode
}

// upfront moves queueItem in front of queue
func (q *queue) upfront(n *queueItem) {
	if n.next == nil { // already in front
		return
	}
	if n.prev == nil { // this is tail
		n.next.prev = nil
		q.tail = n.next
		n.next = nil
		q.head.next = n
		n.prev = q.head
		q.head = n
		return
	}
	// somewhere in the middle
	n.prev.next = n.next
	n.next.prev = n.prev
	n.prev, n.next = q.head, nil
	q.head.next = n
	q.head = n
}

// evict deletes node from queue
func (q *queue) evict(n *queueItem) {
	if n.prev == nil {
		if n.key == q.head.key {
			n.prev, n.next = nil, nil
			q.tail, q.head = nil, nil
			return
		}
		n.next.prev = nil
		q.tail = n.next
		n.prev, n.next = nil, nil
	}
}
