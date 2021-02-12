package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := NewStack(10)
	assert.True(t, s.Empty())

	assert.Nil(t, s.Push(7))
	assert.False(t, s.Empty())
	assert.Equal(t, 1, s.Size())

	assert.Equal(t, 7, s.Peek())
	assert.False(t, s.Empty())

	assert.Equal(t, 7, s.Pop())
	assert.True(t, s.Empty())
	assert.Equal(t, 0, s.Size())
}

func TestStack_Overflow(t *testing.T) {
	s := NewStack(5)
	assert.Nil(t, s.Push(1))
	assert.Nil(t, s.Push(2))
	assert.Nil(t, s.Push(3))
	assert.Nil(t, s.Push(4))
	assert.Nil(t, s.Push(5))
	assert.Equal(t, 5, s.Size())
	assert.Error(t, s.Push(6))
	assert.Equal(t, 5, s.Size())
}

func TestLLStack(t *testing.T) {
	s := NewLLStack()
	assert.True(t, s.Empty())

	s.Push(7)
	assert.False(t, s.Empty())

	assert.Equal(t, 7, s.Peek())
	assert.False(t, s.Empty())

	assert.Equal(t, 7, s.Pop())
	assert.True(t, s.Empty())

	for i := 0; i < 1000; i++ {
		s.Push(i)
	}

	assert.Equal(t, 1000, s.Size())
}
