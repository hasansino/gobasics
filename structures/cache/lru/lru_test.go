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

	c.Put("foo", 9)
	assert.Equal(t, 5, c.Len())

	assert.Equal(t, interface{}(9), c.Get("foo"))
	assert.Nil(t, c.Get("buzz"))
}
