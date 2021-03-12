package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testValues        = []interface{}{1, 2, 3, 4, 5}
	testValuesInverse = []interface{}{5, 4, 3, 2, 1}
)

func TestList_Len(t *testing.T) {
	l := NewList()
	for _, v := range testValues {
		l.Append(v)
	}
	assert.Equal(t, len(testValues), l.Len())
	l.Append(999)
	assert.Equal(t, len(testValues)+1, l.Len())
}

func TestList_Append(t *testing.T) {
	l := NewList()
	for _, v := range testValues {
		l.Append(v)
	}
	assert.Equal(t, testValues, l.Values())
}

func TestList_Prepend(t *testing.T) {
	l := NewList()
	for _, v := range testValues {
		l.Prepend(v)
	}
	assert.Equal(t, testValuesInverse, l.Values())
}

func TestList_Traverse(t *testing.T) {
	l := NewList()
	for _, v := range testValues {
		l.Append(v)
	}
	values := make([]interface{}, 0)
	l.Traverse(func(n *Node) bool {
		values = append(values, n.Data)
		return true
	})
	assert.Equal(t, testValues, values)
}

func TestList_SearchIdx(t *testing.T) {
	l := NewList()
	for _, v := range testValues {
		l.Append(v)
	}
	assert.Nil(t, NewList().SearchIdx(0))
	assert.Nil(t, l.SearchIdx(-1))
	assert.Nil(t, l.SearchIdx(999))
	found := l.SearchIdx(0)
	assert.NotNil(t, found)
	assert.Equal(t, 1, found.Data)
	found = l.SearchIdx(2)
	assert.NotNil(t, found)
	assert.Equal(t, 3, found.Data)
	found = l.SearchIdx(4)
	assert.NotNil(t, found)
	assert.Equal(t, 5, found.Data)
}

func TestList_SearchValue(t *testing.T) {
	l := NewList()
	for _, v := range testValues {
		l.Append(v)
	}
	assert.Nil(t, NewList().SearchValue(999))
	assert.Nil(t, l.SearchValue(999))
	found := l.SearchValue(1)
	assert.NotNil(t, found)
	assert.Equal(t, 1, found.Data)
	found = l.SearchValue(5)
	assert.NotNil(t, found)
	assert.Equal(t, 5, found.Data)
}

func TestList_DeleteIdx(t *testing.T) {
	l := NewList()
	for _, v := range testValues {
		l.Append(v)
	}
	assert.False(t, NewList().DeleteIdx(0))
	assert.False(t, l.DeleteIdx(-1))
	assert.False(t, l.DeleteIdx(999))
	assert.True(t, l.DeleteIdx(0))
	assert.Equal(t, []interface{}{2, 3, 4, 5}, l.Values())
	assert.True(t, l.DeleteIdx(3))
	assert.Equal(t, []interface{}{2, 3, 4}, l.Values())
	assert.True(t, l.DeleteIdx(1))
	assert.Equal(t, []interface{}{2, 4}, l.Values())
}

func TestList_DeleteValue(t *testing.T) {
	l := NewList()
	for _, v := range testValues {
		l.Append(v)
	}
	assert.False(t, NewList().DeleteValue(0))
	assert.False(t, l.DeleteValue(999))
	assert.True(t, l.DeleteValue(1))
	assert.Equal(t, []interface{}{2, 3, 4, 5}, l.Values())
	assert.True(t, l.DeleteValue(5))
	assert.Equal(t, []interface{}{2, 3, 4}, l.Values())
	assert.True(t, l.DeleteValue(3))
	assert.Equal(t, []interface{}{2, 4}, l.Values())
}

func TestList_UpdateIdx(t *testing.T) {
	l := NewList()
	for _, v := range testValues {
		l.Append(v)
	}
	assert.False(t, l.UpdateIdx(-1, 999))
	assert.False(t, l.UpdateIdx(123, 999))
	assert.True(t, l.UpdateIdx(0, 999))
	assert.Equal(t, []interface{}{999, 2, 3, 4, 5}, l.Values())
}

func TestList_UpdateValue(t *testing.T) {
	l := NewList()
	for _, v := range testValues {
		l.Append(v)
	}
	assert.False(t, l.UpdateValue(123, 999))
	assert.True(t, l.UpdateValue(1, 999))
	assert.Equal(t, []interface{}{999, 2, 3, 4, 5}, l.Values())
}

func TestList_Merge(t *testing.T) {
	l := NewList()
	l.Append(1)
	l.Append(2)
	l2 := NewList()
	l2.Append(3)
	l2.Append(4)
	l.Merge(l2)
	assert.Equal(t, []interface{}{1, 2, 3, 4}, l.Values())
}

func TestList_Sort(t *testing.T) {
	l := NewList()
	for _, v := range []interface{}{9, 8, 7, 1, 4, 3, 2, 5, 6} {
		l.Append(v)
	}
	l.Sort(func(v1, v2 interface{}) bool {
		return v2.(int) < v1.(int)
	})
	assert.Equal(t, []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}, l.Values())
	l.Sort(func(v1, v2 interface{}) bool {
		return v2.(int) > v1.(int)
	})
	assert.Equal(t, []interface{}{9, 8, 7, 6, 5, 4, 3, 2, 1}, l.Values())
}
