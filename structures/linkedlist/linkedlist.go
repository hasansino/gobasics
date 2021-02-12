package linkedlist

// https://en.wikipedia.org/wiki/Linked_list

// LinkType represents type of linking used in list
type LinkType int

const (
	TypeSinglyLinked LinkType = iota + 1
	TypeDoubleLinked
	TypeCircular
)

// List is linked list
type List struct {
	head  *Node
	lType LinkType
}

// Node of a list
type Node struct {
	data interface{}
	prev *Node
	next *Node
}

// NewList creates new list of given type
func NewList(t LinkType) *List {
	return &List{lType: t}
}

// Len of a list
func (l *List) Len() int {
	var ret int

	switch l.lType {
	case TypeSinglyLinked:
		for i := l.head; i != nil; i = i.next {
			ret++
		}
	case TypeDoubleLinked:
	case TypeCircular:
	}

	return ret
}

// Values returns all values is the same order they are in list
func (l *List) Values() []interface{} {
	ret := make([]interface{}, 0)

	switch l.lType {
	case TypeSinglyLinked:
		for i := l.head; i != nil; {
			ret = append(ret, i.data)
			i = i.next
		}
	case TypeDoubleLinked:
	case TypeCircular:
	}
	return ret
}

// Append value to list
func (l *List) Append(v interface{}) {
	n := &Node{data: v}
	switch l.lType {
	case TypeSinglyLinked:
		// insert node in empty list
		if l.head == nil {
			l.head = n
		} else {
			// find last node in list
			last := l.head
			for last.next != nil {
				last = last.next
			}
			last.next = n
		}
	case TypeDoubleLinked:
	case TypeCircular:
	}
}

// Prepend value to list
func (l *List) Prepend(v interface{}) {
	n := &Node{data: v}
	switch l.lType {
	case TypeSinglyLinked:
		n.next = l.head
		l.head = n
	case TypeDoubleLinked:
	case TypeCircular:
	}
}

// Traverse list with given callback
func (l *List) Traverse(f func(n *Node) bool) {
	switch l.lType {
	case TypeSinglyLinked:
		for i := l.head; i != nil; i = i.next {
			if ok := f(i); !ok {
				break
			}
		}
	case TypeDoubleLinked:
	case TypeCircular:
	}
}

// SearchIdx returns node by given index or nil
func (l *List) SearchIdx(idx int) *Node {
	if l.head == nil || idx < 0 {
		return nil
	}

	var n *Node

	switch l.lType {
	case TypeSinglyLinked:
		n = l.head
		for i := 0; i < idx; i++ {
			if n.next == nil {
				return nil // index out of range
			}
			n = n.next
		}
	case TypeDoubleLinked:
	case TypeCircular:
	}

	return n
}

// SearchValue returns first found node by given value or nil
func (l *List) SearchValue(v interface{}) *Node {
	if l.head == nil {
		return nil
	}

	switch l.lType {
	case TypeSinglyLinked:
		for n := l.head; n != nil; n = n.next {
			if n.data == v {
				return n
			}
		}
	case TypeDoubleLinked:
	case TypeCircular:
	}

	return nil
}

// DeleteIdx removes node by its index
func (l *List) DeleteIdx(idx int) bool {
	if l.head == nil || idx < 0 {
		return false
	}

	switch l.lType {
	case TypeSinglyLinked:
		// special case if we are deleting head
		if idx == 0 {
			l.head = l.head.next
			return true
		}
		// find previous node
		prev := l.SearchIdx(idx - 1)
		if prev == nil {
			return false
		}
		prev.next = prev.next.next
		return true
	case TypeDoubleLinked:
	case TypeCircular:
	}

	return false
}

// DeleteValue removes firs occurrence of a node with given value
func (l *List) DeleteValue(v interface{}) bool {
	if l.head == nil {
		return false
	}

	switch l.lType {
	case TypeSinglyLinked:
		var prev *Node
		for n := l.head; n != nil; n = n.next {
			if n.data == v { // found
				if prev == nil {
					l.head = l.head.next
				} else {
					prev.next = n.next
				}
				return true
			}
			prev = n
		}
	case TypeDoubleLinked:
	case TypeCircular:
	}

	return false
}

// UpdateIdx updates value of node with given index
func (l *List) UpdateIdx(idx int, v interface{}) bool {
	if n := l.SearchIdx(idx); n != nil {
		n.data = v
		return true
	}
	return false
}

// UpdateValue updates value of a first found node with given value
func (l *List) UpdateValue(s, v interface{}) bool {
	if n := l.SearchValue(s); n != nil {
		n.data = v
		return true
	}
	return false
}

// Merge two linked lists
func (l *List) Merge(with *List) {
	if with.head == nil {
		return
	}

	switch l.lType {
	case TypeSinglyLinked:
		last := l.head
		for last.next != nil {
			last = last.next
		}
		last.next = with.head
	case TypeDoubleLinked:
	case TypeCircular:
	}
}

// Sort linked list with provided sorting function
func (l *List) Sort(fn func(v1, v2 interface{}) bool) {
	if l.head == nil {
		return
	}

	switch l.lType {
	case TypeSinglyLinked:
		for {
			var (
				prev   *Node
				dryRun = true
			)
			for i := l.head; i.next != nil; {
				// check if we need to swap them
				if fn(i.data, i.next.data) {
					// create temporary nodes
					var (
						tmpCurr = &Node{}
						tmpNext = &Node{}
					)

					tmpCurr.data = i.next.data
					tmpCurr.next = tmpNext
					tmpNext.data = i.data
					tmpNext.next = i.next.next

					if prev == nil { // means we are at pos 0 (head)
						l.head = tmpCurr // new head
					} else {
						prev.next = tmpCurr
					}

					prev = tmpCurr
					i = tmpNext

					dryRun = false

					continue
				}

				prev = i
				i = i.next
			}

			if dryRun {
				break
			}
		}
	case TypeDoubleLinked:
	case TypeCircular:
	}
}
