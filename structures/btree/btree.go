package btree

// https://en.wikipedia.org/wiki/Binary_search_tree
// https://afteracademy.com/blog/what-is-a-tree-data-structure
// https://afteracademy.com/blog/binary-search-tree-introduction-operations-and-applications

// ComparisonFn used to compare values in BST
// Returns true if v2 is greater(more prioritised) than v1
type ComparisonFn func(v1, v2 interface{}) bool

// BTree is BST implementation
type BTree struct {
	root         *Node
	comparisonFn ComparisonFn
}

// Node of BST
type Node struct {
	data   interface{}
	parent *Node
	left   *Node
	right  *Node
}

// NewBTree creates new instance of BST
func NewBTree(fn ComparisonFn) *BTree {
	return &BTree{comparisonFn: fn}
}

// Insert new value to btree
func (t *BTree) Insert(v interface{}) {
	if t.root == nil { // empty tree, just insert as root node
		t.root = &Node{data: v}
		return
	}
	t.insertAfter(t.root, v)
}

// insertAfter traverses tree and finds a spot to insert new value
func (t *BTree) insertAfter(n *Node, v interface{}) {
	if t.comparisonFn(n.data, v) { // goes to right side
		if n.right == nil {
			n.right = &Node{data: v, parent: n}
		} else {
			t.insertAfter(n.right, v)
		}
	} else { // goes to left side
		if n.left == nil {
			n.left = &Node{data: v, parent: n}
		} else {
			t.insertAfter(n.left, v)
		}
	}
}

// Delete value from btree
func (t *BTree) Delete(v interface{}) bool {
	n := t.Search(v)

	// no such node exists
	if n == nil {
		return false
	}

	// deletion of root node
	if n.parent == nil {
		switch {
		case n.left == nil && n.right == nil:
			t.root = nil // empty the tree
		case n.left != nil:
			n.left.parent = nil
			t.root = n.left
		case n.right != nil:
			n.right.parent = nil
			t.root = n.right
		default:
			// @TODO
		}
		return true
	}

	// found node is not root node
	switch {
	case n.left == nil && n.right == nil:
		if t.comparisonFn(n.parent.data, n.data) {
			n.parent.right = nil
		} else {
			n.parent.left = nil
		}
	case n.right != nil && n.left == nil: // one node, right
		n.right.parent = n.parent // replace parent
		if t.comparisonFn(n.parent.data, n.data) {
			n.parent.right = n.right
		} else {
			n.parent.left = n.right
		}
	case n.left != nil && n.right == nil: // one node, left
		n.left.parent = n.parent // replace parent
		if t.comparisonFn(n.parent.data, n.data) {
			n.parent.right = n.left
		} else {
			n.parent.left = n.left
		}
	default: // both children present
		s := t.inOrderSuccessor(n)
		// remove parent relation
		if t.comparisonFn(s.parent.data, s.data) {
			s.parent.right = nil
		} else {
			s.parent.left = nil
		}
		// replace nodes
		s.left, s.right = n.left, n.right
		if t.comparisonFn(n.parent.data, n.data) {
			n.parent.right = s
		} else {
			n.parent.left = s
		}
		s.parent = n.parent
	}

	return true
}

// inOrderSuccessor of a node
func (t *BTree) inOrderSuccessor(n *Node) *Node {
	s := n.right
	for s.left != nil {
		s = s.left
	}
	return s
}

// Search value
func (t *BTree) Search(v interface{}) *Node {
	if t.root == nil {
		return nil // empty tree
	}
	for i := t.root; ; {
		greater := t.comparisonFn(i.data, v)
		if greater {
			if i.right == nil {
				break // no answer
			} else {
				i = i.right // traverse to right
			}
		} else {
			if !t.comparisonFn(v, i.data) { // equal?
				return i // found it!
			} else {
				if i.left == nil {
					break // no answer
				} else {
					i = i.left // traverse to left
				}
			}
		}
	}
	return nil
}
