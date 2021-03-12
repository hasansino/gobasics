package lru

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	c := NewCache(5)
	c.Put("test", 1)
	c.Put("test2", 1)
	c.Put("test3", 1)
	c.Put("test4", 1)
	c.Put("test5", 1)

	assert.Equal(t, 5, c.Len())
	assert.Equal(t, []string{"test", "test2", "test3", "test4", "test5"}, queueList(c))

	c.Put("foo", 9)
	assert.Equal(t, 5, c.Len())
	assert.Equal(t, []string{"test2", "test3", "test4", "test5", "foo"}, queueList(c))

	assert.Equal(t, interface{}(9), c.Get("foo"))
	assert.Equal(t, interface{}(1), c.Get("test4"))
	assert.Equal(t, interface{}(1), c.Get("test3"))
	assert.Equal(t, []string{"test2", "test5", "foo", "test4", "test3"}, queueList(c))

	assert.Nil(t, c.Get("buzz"))
	assert.Nil(t, c.Get("bar"))

	c.Put("buzz", 2)
	assert.Nil(t, c.Get("test2"))
	assert.Equal(t, []string{"test5", "foo", "test4", "test3", "buzz"}, queueList(c))

	c.Put("go", "dam")
	assert.Equal(t, []string{"foo", "test4", "test3", "buzz", "go"}, queueList(c))

	assert.Equal(t, interface{}(9), c.Get("foo"))
	assert.Equal(t, []string{"test4", "test3", "buzz", "go", "foo"}, queueList(c))

	assert.Nil(t, c.queue.tail.prev)
	assert.Nil(t, c.queue.head.next)
}

func TestLRUCache_SmallCache(t *testing.T) {
	c := NewCache(1)
	c.Put("test", 1)
	c.Put("test2", 1)
	assert.Equal(t, 1, c.Len())
	assert.Equal(t, []string{"test2"}, queueList(c))
	assert.Equal(t, interface{}(1), c.Get("test2"))
	assert.Equal(t, []string{"test2"}, queueList(c))

	assert.Nil(t, c.queue.tail.prev)
	assert.Nil(t, c.queue.head.next)

	c2 := NewCache(2)
	c2.Put("test", 1)
	c2.Put("test2", 1)
	assert.Equal(t, 2, c2.Len())
	assert.Equal(t, []string{"test", "test2"}, queueList(c2))
	c2.Get("test")
	assert.Equal(t, []string{"test2", "test"}, queueList(c2))

	assert.Nil(t, c2.queue.tail.prev)
	assert.Nil(t, c2.queue.head.next)
}

func queueList(c *Cache) []string {
	ret := make([]string, 0, c.size)
	for n := c.queue.tail; n != nil; n = n.next {
		ret = append(ret, n.key)
	}
	return ret
}
