package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q := NewQueue(5)
	assert.True(t, q.Empty())

	assert.NoError(t, q.Enqueue(1))
	assert.NoError(t, q.Enqueue(2))
	assert.NoError(t, q.Enqueue(3))
	assert.Equal(t, 3, q.Size())
	assert.NoError(t, q.Enqueue(4))
	assert.NoError(t, q.Enqueue(5))
	assert.Error(t, q.Enqueue(6)) // queue is full

	assert.Equal(t, 1, q.Dequeue())
	assert.Equal(t, 4, q.Size())
	assert.Equal(t, 2, q.Dequeue())
	assert.Equal(t, 3, q.Size())
	assert.Equal(t, 3, q.Dequeue())
	assert.Equal(t, 4, q.Dequeue())
	assert.Equal(t, 5, q.Dequeue())
	assert.Equal(t, 0, q.Size())
	assert.True(t, q.Empty())

	assert.NoError(t, q.Enqueue(999))
	assert.Equal(t, 1, q.Size())
	assert.Equal(t, 999, q.Front())
	assert.Equal(t, 1, q.Size())
}

func TestLLQueue(t *testing.T) {
	q := NewLLQueue(5)
	assert.True(t, q.Empty())

	assert.NoError(t, q.Enqueue(1))
	assert.NoError(t, q.Enqueue(2))
	assert.NoError(t, q.Enqueue(3))
	assert.Equal(t, 3, q.Size())
	assert.NoError(t, q.Enqueue(4))
	assert.NoError(t, q.Enqueue(5))
	assert.Error(t, q.Enqueue(6)) // queue is full

	assert.Equal(t, 1, q.Dequeue())
	assert.Equal(t, 4, q.Size())
	assert.Equal(t, 2, q.Dequeue())
	assert.Equal(t, 3, q.Size())
	assert.Equal(t, 3, q.Dequeue())
	assert.Equal(t, 4, q.Dequeue())
	assert.Equal(t, 5, q.Dequeue())
	assert.Equal(t, 0, q.Size())
	assert.True(t, q.Empty())

	assert.NoError(t, q.Enqueue(999))
	assert.Equal(t, 1, q.Size())
	assert.Equal(t, 999, q.Front())
	assert.Equal(t, 1, q.Size())
}
