//
// Package binarytree is implementation of binary tree data structure.
//
// https://en.wikipedia.org/wiki/Binary_search_tree
// https://en.wikipedia.org/wiki/Tree_traversal
// https://afteracademy.com/blog/what-is-a-tree-data-structure
// https://afteracademy.com/blog/binary-search-tree-introduction-operations-and-applications
//
package binarytree

import (
	"errors"
	"fmt"
)

// Less is node comparison function
// It should return true if i < j
// if !Less(i,j) AND !Less(j,i) values are considered equal
type Less func(i, j interface{}) bool

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

// BinaryTree is BST implementation
type BinaryTree struct {
	root   *Node
	lessFn Less
}

// Node of BST
type Node struct {
	Data   interface{}
	parent *Node
	left   *Node
	right  *Node
}

// NewBinaryTree creates new instance of BST
func NewBinaryTree(fn Less) *BinaryTree {
	return &BinaryTree{lessFn: fn}
}

// Validate binary tree integrity
// Returns error should binary tree violate any of following rules:
// * comparison function is nil
// * root node have a parent
// * any right child is less than parent
// * any left child is greater than parent
func (t *BinaryTree) Validate() error {
	if t.lessFn == nil {
		return errors.New("comparisonFn is nil")
	}
	if t.root == nil {
		return nil // empty binary tree
	} else if t.root != nil && t.root.parent != nil {
		return errors.New("root element have parent node")
	}
	return t.validateValues(t.root)
}

// validateValues of a node and it's children recursively
func (t *BinaryTree) validateValues(n *Node) error {
	if n.left != nil {
		if t.lessFn(n.Data, n.left.Data) {
			return fmt.Errorf("value of left child (%v) is greater than parent node (%v)",
				n.left.Data, n.Data)
		}
		if err := t.validateValues(n.left); err != nil {
			return err
		}
	}
	if n.right != nil {
		if t.lessFn(n.right.Data, n.Data) {
			return fmt.Errorf("value of right child (%v) is lesser than parent node (%v)",
				n.left.Data, n.Data)
		}
		if err := t.validateValues(n.right); err != nil {
			return err
		}
	}
	return nil
}

// Len of a binary tree is a number of elements it have
func (t *BinaryTree) Len() int {
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

// Min value of binary tree
// Returns nil if binary tree is empty
func (t *BinaryTree) Min() interface{} {
	if t.root == nil {
		return nil
	}
	n := t.root
	for n.left != nil {
		n = n.left
	}
	return n.Data
}

// Max value of binary tree
// Returns nil if binary tree is empty
func (t *BinaryTree) Max() interface{} {
	if t.root == nil {
		return nil
	}
	n := t.root
	for n.right != nil {
		n = n.right
	}
	return n.Data
}

// Insert new value to binary tree
// Duplicate values are ignored
func (t *BinaryTree) Insert(v interface{}) {
	if t.root == nil { // empty tree, just insert as root node
		t.root = &Node{Data: v}
		return
	}
	t.insertAfter(t.root, v)
}

// insertAfter traverses tree and finds a spot to insert new value
func (t *BinaryTree) insertAfter(n *Node, v interface{}) {
	if t.lessFn(n.Data, v) { // goes to right side
		if n.right == nil {
			n.right = &Node{Data: v, parent: n}
		} else {
			t.insertAfter(n.right, v)
		}
	} else if !t.lessFn(v, n.Data) {
		return // duplicate
	} else { // goes to left side
		if n.left == nil {
			n.left = &Node{Data: v, parent: n}
		} else {
			t.insertAfter(n.left, v)
		}
	}
}

// Delete value from binary tree
func (t *BinaryTree) Delete(v interface{}) bool {
	n := t.Search(v)

	// no such node exists
	if n == nil {
		return false
	}

	// deletion of root node
	if n.parent == nil {
		switch {
		case n.left == nil && n.right == nil: // no children
			t.root = nil // empty binary tree
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
	case n.left == nil && n.right == nil: // no children
		t.removeParentRelation(n)
	case n.right != nil && n.left == nil: // one node, right
		n.right.parent = n.parent // replace parent
		if t.lessFn(n.parent.Data, n.Data) {
			n.parent.right = n.right
		} else {
			n.parent.left = n.right
		}
	case n.left != nil && n.right == nil: // one node, left
		n.left.parent = n.parent // replace parent
		if t.lessFn(n.parent.Data, n.Data) {
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
	leftDeepChild
	rightDeepChild
)

// nodeSuccessor finds a node that should replace deleted node
func (t *BinaryTree) nodeSuccessor(n *Node) (*Node, successorType) {
	// simple case, two scenarios
	// we have n.left node, that does not have right child
	// we have n.right node, that does not have left child
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
		return leftSuccessor, leftDeepChild
	}
	rightSuccessor := n.right
	for rightSuccessor.left != nil {
		rightSuccessor = rightSuccessor.left
	}
	if rightSuccessor.right == nil {
		return rightSuccessor, rightDeepChild
	}
	// by default return left one
	return leftSuccessor, leftDeepChild
}

// replaceNode with its successor
func (t *BinaryTree) replaceNode(n, s *Node, st successorType) {
	var replacingRootNode = n.parent == nil
	// detach successor as child of its parent node
	t.removeParentRelation(s)
	// root node is replaced?
	if replacingRootNode {
		t.root = s
	}
	// replace nodes
	switch st {
	case leftChild, leftDeepChild:
		s.right = n.right
		n.right.parent = s
		if replacingRootNode {
			s.parent = nil
		} else {
			s.parent = n.parent
			t.setAsParent(s, n.parent)
		}
		if st == leftDeepChild {
			tmp := s
			for tmp.left != nil {
				tmp = tmp.left
			}
			tmp.left = n.left
			n.left.parent = tmp
		}
	case rightChild, rightDeepChild:
		s.left = n.left
		n.left.parent = s
		if replacingRootNode {
			s.parent = nil
		} else {
			s.parent = n.parent
			t.setAsParent(s, n.parent)
		}
		if st == rightDeepChild {
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
func (t *BinaryTree) removeParentRelation(n *Node) {
	// if n > parent, n is right child
	if t.lessFn(n.parent.Data, n.Data) {
		n.parent.right = nil
	} else { // otherwise n is left child
		n.parent.left = nil
	}
}

// setAsParent sets n as child of p
func (t *BinaryTree) setAsParent(n, p *Node) {
	if t.lessFn(p.Data, n.Data) {
		p.right = n
	} else {
		p.left = n
	}
}

// Search value
func (t *BinaryTree) Search(v interface{}) *Node {
	if t.root == nil {
		return nil // empty tree
	}
	for i := t.root; ; {
		greater := t.lessFn(i.Data, v)
		if greater {
			if i.right == nil {
				break // no answer
			} else {
				i = i.right // traverse to right
			}
		} else {
			if !t.lessFn(v, i.Data) { // equal?
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

// Traverse binary tree with given traversal order and callback function
// which will be called on each traversed node.
// Callback can return false to abort traversal or true to continue.
func (t *BinaryTree) Traverse(o Order, callback func(*Node) bool) {
	t.traverse(t.root, o, callback)
}

// TraverseFrom is identical to Traverse, but additionally
// accepts a node from which traversal is started.
func (t *BinaryTree) TraverseFrom(n *Node, o Order, callback func(*Node) bool) {
	t.traverse(n, o, callback)
}

func (t *BinaryTree) traverse(n *Node, o Order, callback func(*Node) bool) {
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
