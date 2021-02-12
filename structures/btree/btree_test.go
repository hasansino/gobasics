package btree

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func dummyBtree() *BTree {
	btree := NewBTree(func(v1, v2 interface{}) bool {
		v1Int, ok := v1.(int)
		v2Int, ok2 := v2.(int)
		if ok && ok2 {
			return v2Int > v1Int
		}
		return false
	})
	btree.Insert(10)
	btree.Insert(5)
	btree.Insert(11)
	btree.Insert(8)
	btree.Insert(24)
	btree.Insert(4)
	btree.Insert(6)
	btree.Insert(9)
	btree.Insert(3)
	return btree
}

func TestBtreeSearch(t *testing.T) {
	btree := dummyBtree()
	searchRes := btree.Search(24)
	assert.NotNil(t, searchRes)
	assert.Equal(t, 24, searchRes.data)
	searchRes = btree.Search(10) // root node
	assert.NotNil(t, searchRes)
	assert.Equal(t, 10, searchRes.data)
}

func TestBtree_Delete(t *testing.T) {
	var btree *BTree

	btree = dummyBtree()
	assert.False(t, btree.Delete(999)) // does not exist
	assert.True(t, btree.Delete(24))   // leaf, right side
	assert.Nil(t, btree.Search(24))
	btree = dummyBtree()
	assert.True(t, btree.Delete(11)) // only one child
	assert.Nil(t, btree.Search(11))
	btree = dummyBtree()
	assert.True(t, btree.Delete(4)) // leaf, left side
	assert.Nil(t, btree.Search(4))
	btree = dummyBtree()
	assert.True(t, btree.Delete(5)) // two children
	assert.Nil(t, btree.Search(5))
	spew.Config.DisablePointerAddresses = true
	spew.Dump(btree)
}
