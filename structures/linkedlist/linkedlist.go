//
// Package linkedlist implements doubly-linked list.
//
package linkedlist

// https://en.wikipedia.org/wiki/Linked_list
// https://afteracademy.com/blog/types-of-linked-list-and-operation-on-linked-list

// List is linked list
type List struct {
	head *Node
	tail *Node
}

// Node of a list
type Node struct {
	Data interface{}
	prev *Node
	next *Node
}

// NewList creates new list of given type
func NewList() *List {
	return &List{}
}

// Len of a list
func (l *List) Len() int {
	var ret int
	for i := l.tail; i != nil; i = i.next {
		ret++
	}
	return ret
}

// Values returns all values is the same order they are in list
func (l *List) Values() []interface{} {
	ret := make([]interface{}, 0)
	for i := l.tail; i != nil; {
		ret = append(ret, i.Data)
		i = i.next
	}
	return ret
}

// Append value to list
func (l *List) Append(v interface{}) {
	n := &Node{Data: v}
	// insert node in empty list
	if l.head == nil {
		l.head = n
		l.tail = n
	} else {
		n.prev = l.head
		l.head.next = n
		l.head = n
	}
}

// Prepend value to list
func (l *List) Prepend(v interface{}) {
	n := &Node{Data: v}
	if l.tail == nil {
		l.head = n
		l.tail = n
	} else {
		n.next = l.tail
		l.tail.prev = n
		l.tail = n
	}
}

// Traverse list with given callback
func (l *List) Traverse(f func(n *Node) bool) {
	for i := l.tail; i != nil; i = i.next {
		if ok := f(i); !ok {
			break
		}
	}
}

// SearchIdx returns node by given index or nil
func (l *List) SearchIdx(idx int) *Node {
	if l.tail == nil || idx < 0 {
		return nil
	}
	n := l.tail
	for i := 0; i < idx; i++ {
		if n.next == nil {
			return nil // index out of range
		}
		n = n.next
	}
	return n
}

// SearchValue returns first found node by given value or nil
func (l *List) SearchValue(v interface{}) *Node {
	if l.tail == nil {
		return nil
	}
	for n := l.tail; n != nil; n = n.next {
		if n.Data == v {
			return n
		}
	}
	return nil
}

// DeleteIdx removes node by its index
func (l *List) DeleteIdx(idx int) bool {
	if l.tail == nil || idx < 0 {
		return false
	}
	return l.deleteNode(l.SearchIdx(idx))
}

// DeleteValue removes first occurrence of a node with given value
func (l *List) DeleteValue(v interface{}) bool {
	if l.head == nil {
		return false
	}
	return l.deleteNode(l.SearchValue(v))
}

// deleteNode from a list
func (l *List) deleteNode(n *Node) bool {
	if n == nil {
		return false
	}
	// only node in a list
	if n.prev == nil && n.next == nil {
		l.tail, l.head = nil, nil
		return true
	}
	switch {
	case n.prev == nil: // tail
		l.tail = n.next
		l.tail.prev = nil
		n.next = nil
	case n.next == nil: // head
		l.head = n.prev
		l.head.next = nil
		n.prev = nil
	default: // middle
		n.prev.next = n.next
		n.next.prev = n.prev
		n.prev, n.next = nil, nil
	}
	return true
}

// UpdateIdx updates value of node with given index
func (l *List) UpdateIdx(idx int, v interface{}) bool {
	if n := l.SearchIdx(idx); n != nil {
		n.Data = v
		return true
	}
	return false
}

// UpdateValue updates value of a first found node with given value
func (l *List) UpdateValue(s, v interface{}) bool {
	if n := l.SearchValue(s); n != nil {
		n.Data = v
		return true
	}
	return false
}

// Merge two linked lists
func (l *List) Merge(with *List) {
	if with.tail == nil {
		return
	}
	l.head.next = with.tail
}

// Sort linked list with provided sorting function
// Very slow, unoptimized sorting based on bubble sort
// O(n^2+n)
func (l *List) Sort(fn func(v1, v2 interface{}) bool) {
	if l.tail == nil {
		return
	}
	for {
		var (
			prev   *Node
			dryRun = true // indicates no changes were made during single loop
		)
		// loop over all nodes in list starting with HEAD
		for i := l.tail; i.next != nil; {
			// check if we need to swap current with next node
			if fn(i.Data, i.next.Data) {
				// create temporary nodes
				var (
					tmpCurr = &Node{}
					tmpNext = &Node{}
				)

				tmpCurr.Data = i.next.Data
				tmpCurr.next = tmpNext
				tmpNext.Data = i.Data
				tmpNext.next = i.next.next

				if prev == nil { // means we are at pos 0 (head)
					l.tail = tmpCurr
				} else {
					prev.next = tmpCurr
				}

				prev = tmpCurr
				i = tmpNext

				dryRun = false // we will loop over all nodes again

				continue
			}

			prev = i
			i = i.next
		}

		// no changes were made during this loop
		// this means we sorted everything we could
		if dryRun {
			// restore n.prev pointers
			for n := l.tail; n != nil; n = n.next {
				if n.next != nil {
					n.next.prev = n
				}
			}
			break
		}
	}
}
