package btree

import (
	"errors"
	"fmt"
)

// https://en.wikipedia.org/wiki/Binary_search_tree
// https://en.wikipedia.org/wiki/Tree_traversal
// https://afteracademy.com/blog/what-is-a-tree-data-structure
// https://afteracademy.com/blog/binary-search-tree-introduction-operations-and-applications

// ComparisonFn used to compare values in BST
// Returns true if v2 is greater(more prioritised) than v1
// If !ComparisonFn(v1,v2) AND !ComparisonFn(v2,v1) we consider
// values to be equal.
type ComparisonFn func(v1, v2 interface{}) bool

// Order of traversal operation
type Order uint8

const (
	// full names
	PreOrder Order = 1 << iota
	InOrder
	ReverseInOrder
	PostOrder
	// short names
	NLR = PreOrder
	LNR = InOrder
	RNL = ReverseInOrder
	LRN = PostOrder
)

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

// Value return value of node
func (n *Node) Value() interface{} {
	return n.data
}

// NewBTree creates new instance of BST
func NewBTree(fn ComparisonFn) *BTree {
	return &BTree{comparisonFn: fn}
}

// Validate btree integrity
// Returns error should btree violate any of following rules:
// * comparison function is nil
// * root node have a parent
// * any right sibling is less than parent
// * any left sibling is greater than parent
func (t *BTree) Validate() error {
	if t.comparisonFn == nil {
		return errors.New("comparisonFn is nil")
	}
	if t.root == nil {
		return nil // empty btree
	} else if t.root != nil && t.root.parent != nil {
		return errors.New("root element have parent node")
	}
	return t.validateValues(t.root)
}

// validateValues of a node and it's siblings recursively
func (t *BTree) validateValues(n *Node) error {
	if n.left != nil {
		if t.comparisonFn(n.data, n.left.data) {
			return fmt.Errorf("value of left sibling (%v) is greater than parent node (%v)",
				n.left.data, n.data)
		}
		if err := t.validateValues(n.left); err != nil {
			return err
		}
	}
	if n.right != nil {
		if t.comparisonFn(n.right.data, n.data) {
			return fmt.Errorf("value of right sibling (%v) is lesser than parent node (%v)",
				n.left.data, n.data)
		}
		if err := t.validateValues(n.right); err != nil {
			return err
		}
	}
	return nil
}

// Len of a btree is a number of elements it have
func (t *BTree) Len() int {
	if t.root == nil {
		return 0
	}
	var cnt int
	t.traverse(t.root, NLR, func(node *Node) bool {
		cnt++
		return true
	})
	return cnt
}

// Min value of btree
// Returns nil if btree is empty
func (t *BTree) Min() interface{} {
	if t.root == nil {
		return nil
	}
	n := t.root
	for n.left != nil {
		n = n.left
	}
	return n.data
}

// Max value of btree
// Returns nil if btree is empty
func (t *BTree) Max() interface{} {
	if t.root == nil {
		return nil
	}
	n := t.root
	for n.right != nil {
		n = n.right
	}
	return n.data
}

// Insert new value to btree
// Duplicate values are ignored
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
	} else if !t.comparisonFn(v, n.data) {
		return // duplicate
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
		case n.left == nil && n.right == nil: // no siblings
			t.root = nil // empty btree
		case n.right != nil && n.left == nil: // one node, right
			n.right.parent = nil
			t.root = n.right
		case n.left != nil && n.right == nil: // one node, left
			n.left.parent = nil
			t.root = n.left
		default: // both children present
			s, st := t.nodeSuccessor(n)
			t.replaceNode(n, s, st)
		}
		return true
	}

	// found node is not root node
	switch {
	case n.left == nil && n.right == nil: // no siblings
		t.removeParentRelation(n)
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
		s, st := t.nodeSuccessor(n)
		t.replaceNode(n, s, st)
	}

	return true
}

type successorType uint8

const (
	leftChild successorType = 1 << iota
	rightChild
	leftDeepSibling
	rightDeepSibling
)

// nodeSuccessor finds a node that should replace deleted node
func (t *BTree) nodeSuccessor(n *Node) (*Node, successorType) {
	// simple case, two scenarios
	// we have n.left node, that does not have right sibling
	// we have n.right node, that does not have left sibling
	switch {
	case n.left.right == nil:
		return n.left, leftChild
	case n.right.left == nil:
		return n.right, rightChild
	}
	// otherwise we need to find successor on deeper level
	leftSuccessor := n.left
	for leftSuccessor.right != nil {
		leftSuccessor = leftSuccessor.right
	}
	if leftSuccessor.left == nil {
		return leftSuccessor, leftDeepSibling
	}
	rightSuccessor := n.right
	for rightSuccessor.left != nil {
		rightSuccessor = rightSuccessor.left
	}
	if rightSuccessor.right == nil {
		return rightSuccessor, rightDeepSibling
	}
	// by default return left one
	return leftSuccessor, leftDeepSibling
}

// replaceNode with its successor
func (t *BTree) replaceNode(n, s *Node, st successorType) {
	var replacingRootNode = n.parent == nil
	// detach successor as child of its parent node
	t.removeParentRelation(s)
	// root node is replaced?
	if replacingRootNode {
		t.root = s
	}
	// replace nodes
	switch st {
	case leftChild, leftDeepSibling:
		s.right = n.right
		n.right.parent = s
		if replacingRootNode {
			s.parent = nil
		} else {
			s.parent = n.parent
			t.setAsParent(s, n.parent)
		}
		if st == leftDeepSibling {
			tmp := s
			for tmp.left != nil {
				tmp = tmp.left
			}
			tmp.left = n.left
			n.left.parent = tmp
		}
	case rightChild, rightDeepSibling:
		s.left = n.left
		n.left.parent = s
		if replacingRootNode {
			s.parent = nil
		} else {
			s.parent = n.parent
			t.setAsParent(s, n.parent)
		}
		if st == rightDeepSibling {
			tmp := s
			for tmp.right != nil {
				tmp = tmp.right
			}
			tmp.right = n.right
			n.right.parent = tmp
		}
	}
}

// removeParentRelation removes n as n.parent child
func (t *BTree) removeParentRelation(n *Node) {
	// if n > parent, n is right child
	if t.comparisonFn(n.parent.data, n.data) {
		n.parent.right = nil
	} else { // otherwise n is left child
		n.parent.left = nil
	}
}

// setAsParent sets n as child of p
func (t *BTree) setAsParent(n, p *Node) {
	if t.comparisonFn(p.data, n.data) {
		p.right = n
	} else {
		p.left = n
	}
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

// Traverse btree with given traversal order and callback function
// which will be called on each traversed node.
// Callback can return false to abort traversal or true to continue.
func (t *BTree) Traverse(o Order, callback func(*Node) bool) {
	t.traverse(t.root, o, callback)
}

// TraverseFrom is identical to Traverse, but additionally
// accepts a node from which traversal is started.
func (t *BTree) TraverseFrom(n *Node, o Order, callback func(*Node) bool) {
	t.traverse(n, o, callback)
}

func (t *BTree) traverse(n *Node, o Order, callback func(*Node) bool) {
	if n == nil || callback == nil {
		return
	}
	switch o {
	case NLR:
		callback(n)
		t.traverse(n.left, o, callback)
		t.traverse(n.right, o, callback)
	case LNR:
		t.traverse(n.left, o, callback)
		callback(n)
		t.traverse(n.right, o, callback)
	case RNL:
		t.traverse(n.right, o, callback)
		callback(n)
		t.traverse(n.left, o, callback)
	case LRN:
		t.traverse(n.left, o, callback)
		t.traverse(n.right, o, callback)
		callback(n)
	}
}
