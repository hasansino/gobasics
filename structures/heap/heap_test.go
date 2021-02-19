package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	lessFn     = func(i, j interface{}) bool { return i.(int) <= j.(int) }
	testValues = []interface{}{1, 5, 6, 4, 9, 3, 7}
)

func TestHeap_MinHeap_Build(t *testing.T) {
	h := NewHeap(MinHeap, lessFn)
	h.Insert(testValues...)
	assert.Equal(t, []interface{}{1, 4, 3, 5, 9, 6, 7}, h.data)
	assert.Equal(t, 7, h.Len())
}

func TestHeap_MinHeap_MinMax(t *testing.T) {
	h := NewHeap(MinHeap, lessFn)
	h.Insert(testValues...)
	assert.Equal(t, 1, h.Min())
	assert.Equal(t, 9, h.Max())
}

func TestHeap_MinHeap_Peek(t *testing.T) {
	h := NewHeap(MinHeap, lessFn)
	h.Insert(testValues...)
	assert.Equal(t, 1, h.Peek())
}

func TestHeap_MinHeap_Pop(t *testing.T) {
	h := NewHeap(MinHeap, lessFn)
	h.Insert(testValues...)
	assert.Equal(t, 1, h.Pop())
	assert.Equal(t, 6, h.Len())
	assert.Equal(t, []interface{}{3, 4, 6, 5, 9, 7}, h.data)
	assert.Equal(t, 3, h.Pop())
	assert.Equal(t, 5, h.Len())
	assert.Equal(t, []interface{}{4, 5, 6, 7, 9}, h.data)
	assert.Equal(t, 4, h.Pop())
	assert.Equal(t, 4, h.Len())
	assert.Equal(t, []interface{}{5, 7, 6, 9}, h.data)

	assert.Equal(t, 5, h.Pop())
	assert.Equal(t, 3, h.Len())
	assert.Equal(t, []interface{}{6, 7, 9}, h.data)

	assert.Equal(t, 6, h.Pop())
	assert.Equal(t, 2, h.Len())
	assert.Equal(t, []interface{}{7, 9}, h.data)

	assert.Equal(t, 7, h.Pop())
	assert.Equal(t, 1, h.Len())
	assert.Equal(t, []interface{}{9}, h.data)

	assert.Equal(t, 9, h.Pop())
	assert.Equal(t, 0, h.Len())
	assert.Equal(t, []interface{}{}, h.data)

	assert.Nil(t, h.Pop())
}

func TestHeap_MaxHeap_Build(t *testing.T) {
	h := NewHeap(MaxHeap, lessFn)
	h.Insert(testValues...)
	assert.Equal(t, []interface{}{9, 6, 7, 1, 4, 3, 5}, h.data)
	assert.Equal(t, 7, h.Len())
}

func TestHeap_MaxHeap_MinMax(t *testing.T) {
	h := NewHeap(MaxHeap, lessFn)
	h.Insert(testValues...)
	assert.Equal(t, 1, h.Min())
	assert.Equal(t, 9, h.Max())
}

func TestHeap_MaxHeap_Peek(t *testing.T) {
	h := NewHeap(MaxHeap, lessFn)
	h.Insert(testValues...)
	assert.Equal(t, 9, h.Peek())
}

func TestHeap_MaxHeap_Pop(t *testing.T) {
	h := NewHeap(MaxHeap, lessFn)
	h.Insert(testValues...)
	assert.Equal(t, 9, h.Pop())
	assert.Equal(t, 6, h.Len())
	assert.Equal(t, []interface{}{7, 6, 5, 1, 4, 3}, h.data)
	assert.Equal(t, 7, h.Pop())
	assert.Equal(t, 5, h.Len())
	assert.Equal(t, []interface{}{6, 4, 5, 1, 3}, h.data)
	assert.Equal(t, 6, h.Pop())
	assert.Equal(t, 4, h.Len())
	assert.Equal(t, []interface{}{5, 4, 3, 1}, h.data)
}
