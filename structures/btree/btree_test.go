package btree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	lessFn = func(i, j interface{}) bool {
		return i.(int) < j.(int)
	}
	testValues = []interface{}{10, 5, 11, 7, 24, 4, 3, 9, 8, 6, 20, 21, 19} // 13
	//                                 10
	//                               /    \
	//                             5        11
	//                           /   \         \
	//                         4      7          24
	//                       /      /   \       /
	//                     3      6       9    20
	//                                  /     /  \
	//                                8      19  21
)

func generateBTree(values []interface{}) *BTree {
	btree := NewBTree(lessFn)
	for _, v := range values {
		btree.Insert(v)
	}
	return btree
}

func TestBTree_Validate(t *testing.T) {
	btree := generateBTree(testValues)
	assert.NoError(t, btree.Validate())

	btree.root.parent = &Node{}
	assert.Error(t, btree.Validate())

	btree.lessFn = nil
	assert.Error(t, btree.Validate())

	btree = generateBTree(testValues)
	btree.root.left.data = 22
	assert.Error(t, btree.Validate())

	btree = generateBTree(testValues)
	btree.root.right.data = 1
	assert.Error(t, btree.Validate())
}

func TestBTree_Len(t *testing.T) {
	btree := generateBTree(nil) // empty btree
	assert.NoError(t, btree.Validate())
	assert.Equal(t, 0, btree.Len())
	btree = generateBTree(testValues)
	assert.NoError(t, btree.Validate())
	assert.Equal(t, 13, btree.Len())
}

func TestBTree_Min(t *testing.T) {
	btree := generateBTree(nil) // empty btree
	assert.NoError(t, btree.Validate())
	assert.Nil(t, btree.Min())
	btree = generateBTree(testValues) // empty btree
	assert.NoError(t, btree.Validate())
	assert.Equal(t, 3, btree.Min())
}

func TestBTree_Max(t *testing.T) {
	btree := generateBTree(nil) // empty btree
	assert.NoError(t, btree.Validate())
	assert.Nil(t, btree.Min())
	btree = generateBTree(testValues) // empty btree
	assert.NoError(t, btree.Validate())
	assert.Equal(t, 24, btree.Max())
}

func TestBTree_Insert(t *testing.T) {
	btree := generateBTree(nil)
	assert.Equal(t, 0, btree.Len())

	btree.Insert(20)
	assert.Equal(t, 1, btree.Len())

	btree.Insert(21)
	btree.Insert(22)
	btree.Insert(23)
	assert.Equal(t, 4, btree.Len())

	btree.Insert(21) // duplicate
	assert.Equal(t, 4, btree.Len())

	assert.NoError(t, btree.Validate())
}

func TestBtree_Delete_Simple(t *testing.T) {
	btree := generateBTree(testValues)
	assert.False(t, btree.Delete(999)) // does not exist
	assert.Equal(t, 13, btree.Len())
	assert.NoError(t, btree.Validate())

	// root node, only right child
	//                                 10 <-
	//                               /    \
	//                             _        11
	//                           /   \         \
	//                         _      _         24
	//                       /      /   \      /
	//                     _      _       _   20
	//                                  /    /  \
	//                                _     19  21
	btree = generateBTree([]interface{}{10, 11, 24, 20, 19, 21})
	assert.True(t, btree.Delete(10))
	assert.Nil(t, btree.Search(10))
	assert.Equal(t, 5, btree.Len())
	assert.NoError(t, btree.Validate())

	// root node, only left child
	//                                 10 <-
	//                               /    \
	//                             5        __
	//                           /   \        \
	//                         4      7        __
	//                       /      /   \       /
	//                     3      6       9    __
	//                                  /     /  \
	//                                8      __  __
	btree = generateBTree([]interface{}{10, 5, 4, 3, 7, 6, 9, 8})
	assert.True(t, btree.Delete(10))
	assert.Nil(t, btree.Search(10))
	assert.Equal(t, 7, btree.Len())
	assert.NoError(t, btree.Validate())

	// leaf, right side
	//                                 10
	//                               /    \
	//                             5        11
	//                           /   \         \
	//                         4      7         24
	//                       /      /   \       /
	//                     3      6       9    20
	//                                  /     /  \
	//                                8      19  21 <-
	btree = generateBTree(testValues)
	assert.True(t, btree.Delete(21))
	assert.Nil(t, btree.Search(21))
	assert.Equal(t, 12, btree.Len())
	assert.NoError(t, btree.Validate())

	// leaf, left side
	//                                 10
	//                               /    \
	//                             5        11
	//                           /   \         \
	//                         4      7         24
	//                       /      /   \       /
	//                     3 <-    6       9    20
	//                                  /     /  \
	//                                8      19  21
	btree = generateBTree(testValues)
	assert.True(t, btree.Delete(3))
	assert.Nil(t, btree.Search(3))
	assert.Equal(t, 12, btree.Len())
	assert.NoError(t, btree.Validate())

	// not root node, only one right child
	//                                 10
	//                               /    \
	//                             5        11 <-
	//                           /   \         \
	//                         4      7         24
	//                       /      /   \       /
	//                     3       6       9    20
	//                                  /     /  \
	//                                8      19  21
	btree = generateBTree(testValues)
	assert.True(t, btree.Delete(11))
	assert.Nil(t, btree.Search(11))
	assert.Equal(t, 12, btree.Len())
	assert.NoError(t, btree.Validate())

	// not root node, only one left child
	//                                 10
	//                               /    \
	//                             5        11
	//                           /   \         \
	//                         4 <-   7         24
	//                       /      /   \       /
	//                     3       6     9     20
	//                                  /     /  \
	//                                8      19  21
	btree = generateBTree(testValues)
	assert.True(t, btree.Delete(4))
	assert.Nil(t, btree.Search(4))
	assert.Equal(t, 12, btree.Len())
	assert.NoError(t, btree.Validate())
}

func TestBtree_Delete_WithSuccessor(t *testing.T) {
	// root node, right child successor
	//                                 10 <- delete
	//                               /    \
	//                             5        11 <- successor
	//                           /   \         \
	//                         4      7         24
	//                       /      /   \       /
	//                     3      6       9    20
	//                                  /     /  \
	//                                8      19  21
	btree := generateBTree(testValues)
	assert.True(t, btree.Delete(10))
	assert.Nil(t, btree.Search(10))
	assert.Equal(t, 12, btree.Len())
	assert.NoError(t, btree.Validate())

	// root node, left child successor
	//                                 10 <- delete
	//                               /    \
	//                successor -> 5        15
	//                           /        /    \
	//                         4         12     24
	//                       /                 /
	//                     3                  20
	//                                       /  \
	//                                      19  21
	btree = generateBTree([]interface{}{10, 5, 4, 3, 15, 24, 20, 21, 19, 12})
	assert.True(t, btree.Delete(10))
	assert.Nil(t, btree.Search(10))
	assert.Equal(t, 9, btree.Len())
	assert.NoError(t, btree.Validate())

	// root node, deep-left successor
	//                                 10 <- delete
	//                               /    \
	//                             5        15
	//                               \     /   \
	//                   successor -> 7   12    24
	//                                   /  \
	//                                  11  13
	//                                        \
	//                                        14
	btree = generateBTree([]interface{}{10, 5, 7, 15, 24, 12, 11, 13, 14})
	assert.True(t, btree.Delete(10))
	assert.Nil(t, btree.Search(10))
	assert.Equal(t, 8, btree.Len())
	assert.NoError(t, btree.Validate())

	// root node, deep-right successor
	//                                 10 <- delete
	//                               /    \
	//                             5        15
	//                               \     /   \
	//                                7   12    24
	//                               /   /
	//                              6   11 <- successor
	btree = generateBTree([]interface{}{10, 5, 7, 6, 15, 24, 12, 11})
	assert.True(t, btree.Delete(10))
	assert.Nil(t, btree.Search(10))
	assert.Equal(t, 7, btree.Len())
	assert.NoError(t, btree.Validate())

	// root node, deep successor with child nodes
	//                                 22 <- delete
	//                               /    \
	//                             8        60
	//                               \     /  \
	//                  successor -> 12   30  75
	//                               /     \
	//                              10     35
	btree = generateBTree([]interface{}{22, 60, 30, 35, 75, 8, 12, 10})
	assert.True(t, btree.Delete(22))
	assert.Nil(t, btree.Search(22))
	assert.Equal(t, 7, btree.Len())
	assert.NoError(t, btree.Validate())

	// non-root node
	//                                 10
	//                               /    \
	//                         ->  5        11
	//                           /   \         \
	//                         4      7         24
	//                       /      /   \       /
	//                     3      6       9    20
	//                                  /     /  \
	//                                8      19  21
	btree = generateBTree(testValues)
	assert.True(t, btree.Delete(5))
	assert.Nil(t, btree.Search(5))
	assert.Equal(t, 12, btree.Len())
	assert.NoError(t, btree.Validate())

	// non-root node
	//                                 10
	//                               /    \
	//                             5        11
	//                           /   \         \
	//                         4      7 <-      24
	//                       /      /   \       /
	//                     3      6       9    20
	//                                  /     /  \
	//                                8      19  21
	btree = generateBTree(testValues)
	assert.True(t, btree.Delete(5))
	assert.Nil(t, btree.Search(5))
	assert.Equal(t, 12, btree.Len())
	assert.NoError(t, btree.Validate())

	// non-root node
	//                                 10
	//                               /    \
	//                             5        11
	//                           /   \         \
	//                         4      7 <-      24
	//                       /      /   \       /
	//                     3      6       9    20
	//                                  /     /  \
	//                                8      19  21
	btree = generateBTree(testValues)
	assert.True(t, btree.Delete(7))
	assert.Nil(t, btree.Search(7))
	assert.Equal(t, 12, btree.Len())
	assert.NoError(t, btree.Validate())

	// non-root node
	//                                 10
	//                               /    \
	//                             5        11
	//                           /   \         \
	//                         4      7         24
	//                       /      /   \       /
	//                     3      6       9    20 <-
	//                                  /     /  \
	//                                8      19  21
	btree = generateBTree(testValues)
	assert.True(t, btree.Delete(20))
	assert.Nil(t, btree.Search(20))
	assert.Equal(t, 12, btree.Len())
	assert.NoError(t, btree.Validate())
}

func TestBtree_Delete_Multi(t *testing.T) {
	// delete multiple nodes
	//                                 10
	//                               /    \
	//                             5        11
	//                           /   \         \
	//                         4      7         24
	//                       /      /   \       /
	//                     3      6       9    20
	//                                  /     /  \
	//                                8      19  21
	btree := generateBTree(testValues)
	assert.True(t, btree.Delete(10))
	assert.Nil(t, btree.Search(10))
	assert.True(t, btree.Delete(20))
	assert.Nil(t, btree.Search(20))
	assert.True(t, btree.Delete(7))
	assert.Nil(t, btree.Search(7))
	assert.True(t, btree.Delete(4))
	assert.Nil(t, btree.Search(4))
	assert.Equal(t, 9, btree.Len())
	assert.NoError(t, btree.Validate())
}

func TestBTree_Search(t *testing.T) {
	btree := generateBTree(testValues)
	searchRes := btree.Search(24)
	assert.NotNil(t, searchRes)
	assert.Equal(t, 24, searchRes.data)
	searchRes = btree.Search(10) // root node
	assert.NotNil(t, searchRes)
	assert.Equal(t, 10, searchRes.data)
}

func TestBtree_Traverse_NLR(t *testing.T) {
	var (
		traversedPath = make([]interface{}, 0)
		btree         = generateBTree(testValues)
	)
	btree.Traverse(NLR, func(node *Node) bool {
		traversedPath = append(traversedPath, node.data)
		return true
	})
	assert.Equal(t, []interface{}{10, 5, 4, 3, 7, 6, 9, 8, 11, 24, 20, 19, 21}, traversedPath)
}

func TestBtree_Traverse_LNR(t *testing.T) {
	var (
		traversedPath = make([]interface{}, 0)
		btree         = generateBTree(testValues)
	)
	btree.Traverse(LNR, func(node *Node) bool {
		traversedPath = append(traversedPath, node.data)
		return true
	})
	assert.Equal(t, []interface{}{3, 4, 5, 6, 7, 8, 9, 10, 11, 19, 20, 21, 24}, traversedPath)
}

func TestBtree_Traverse_RNL(t *testing.T) {
	var (
		traversedPath = make([]interface{}, 0)
		btree         = generateBTree(testValues)
	)
	btree.Traverse(RNL, func(node *Node) bool {
		traversedPath = append(traversedPath, node.data)
		return true
	})
	assert.Equal(t, []interface{}{24, 21, 20, 19, 11, 10, 9, 8, 7, 6, 5, 4, 3}, traversedPath)
}

func TestBtree_Traverse_LRN(t *testing.T) {
	var (
		traversedPath = make([]interface{}, 0)
		btree         = generateBTree(testValues)
	)
	btree.Traverse(LRN, func(node *Node) bool {
		traversedPath = append(traversedPath, node.data)
		return true
	})
	assert.Equal(t, []interface{}{3, 4, 6, 8, 9, 7, 5, 19, 21, 20, 24, 11, 10}, traversedPath)
}
