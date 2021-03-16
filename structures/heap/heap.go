//
// Package heap implements heap data structure.
//
package heap

import (
	"math"
)

// https://en.wikipedia.org/wiki/Heap_(data_structure)
// https://afteracademy.com/blog/introduction-to-heaps-in-data-structures
// https://afteracademy.com/blog/heap-building-and-heap-sort
// https://afteracademy.com/blog/operations-on-heaps

// Heap math:
// * For any node left and right child is Arr[2*i + 1] and Arr[2*i+2]
// * For any node except root parent is Arr[(i-1)/2]

// Property of heap structure
type Property uint8

const (
	// any parent node can be only greater or equal to child node
	MaxHeap Property = 1 << iota
	// any parent node can be only less or equal to child node
	MinHeap
)

// Less is node comparison function
// It should return true if i < j
// if !Less(i,j) AND !Less(j,i) values are considered equal
type Less func(i, j interface{}) bool

// Heap is heap represented as array
type Heap struct {
	data   []interface{}
	prop   Property
	lessFn Less
}

// NewHeap creates new instance of heap tree
func NewHeap(p Property, l Less) *Heap {
	return &Heap{data: make([]interface{}, 0), prop: p, lessFn: l}
}

// Len of heap
func (h *Heap) Len() int {
	return len(h.data)
}

// Min value in heap
func (h *Heap) Min() interface{} {
	if h.Len() == 0 {
		return nil
	}
	switch h.prop {
	case MinHeap:
		return h.data[0]
	case MaxHeap:
		leavesStartPos := h.numOfLeaves() - 1
		min := h.data[leavesStartPos]
		leavesStartPos++
		for i := leavesStartPos; i < h.Len(); i++ {
			if h.lessFn(h.data[i], min) {
				min = h.data[i]
			}
		}
		return min
	}
	return nil
}

// Max value of a node
func (h *Heap) Max() interface{} {
	if h.Len() == 0 {
		return nil
	}
	switch h.prop {
	case MinHeap:
		leavesStartPos := h.numOfLeaves() - 1
		max := h.data[leavesStartPos]
		leavesStartPos++
		for i := leavesStartPos; i < h.Len(); i++ {
			if h.lessFn(max, h.data[i]) {
				max = h.data[i]
			}
		}
		return max
	case MaxHeap:
		return h.data[0]
	}
	return nil
}

// Peek a value of root node
func (h *Heap) Peek() interface{} {
	switch h.prop {
	case MinHeap:
		return h.Min()
	case MaxHeap:
		return h.Max()
	}
	return nil
}

// Pop removes and returns a root node value
func (h *Heap) Pop() interface{} {
	if h.Len() == 0 {
		return nil
	}

	var val interface{}

	switch h.prop {
	case MinHeap:
		val = h.Min()
		h.delete(0) // delete root
	case MaxHeap:
		val = h.Max()
		h.delete(0) // delete root
	}
	return val
}

// Insert value to heap
func (h *Heap) Insert(values ...interface{}) {
	for _, v := range values {
		h.insert(v)
	}
}

// insert single value into heap
func (h *Heap) insert(v interface{}) {
	// always insert new value at the end of heap
	h.data = append(h.data, v)
	// re-balance heap
	h.upHeapify(h.Len() - 1)
}

// delete value at position pos from heap
func (h *Heap) delete(pos int) {
	h.swap(pos, h.Len()-1)      // swap with last node
	h.data = h.data[:h.Len()-1] // cut last node from heap
	h.downHeapify(pos)          // re-balance heap from pos
}

// upHeapify balances heap from pos upwards
func (h *Heap) upHeapify(pos int) {
	if h.Len() < 2 || pos == 0 {
		return // no balance needed
	}
	parentPos := h.parent(pos)
	switch h.prop {
	case MinHeap:
		if h.lessFn(h.data[pos], h.data[parentPos]) {
			h.swap(pos, parentPos)
			h.upHeapify(parentPos)
		}
	case MaxHeap:
		if h.lessFn(h.data[parentPos], h.data[pos]) {
			h.swap(pos, parentPos)
			h.upHeapify(parentPos)
		}
	}
}

// downHeapify balances heap from pos downwards
func (h *Heap) downHeapify(pos int) {
	if h.Len() < 2 {
		return // no balance needed
	}
	var (
		left  = h.leftChild(pos)
		right = h.rightChild(pos)
	)
	switch h.prop {
	case MinHeap:
		switch {
		case h.posExists(left) && h.posExists(right):
			// if both children present, we need to choose
			// one with smallest value in this case
			swapWith := right
			if h.lessFn(h.data[left], h.data[right]) {
				swapWith = left
			}
			// check if we need to swap values
			if h.lessFn(h.data[swapWith], h.data[pos]) {
				h.swap(swapWith, pos)
				h.downHeapify(swapWith)
			}
		case h.posExists(left):
			if h.lessFn(h.data[left], h.data[pos]) {
				h.swap(left, pos)
				h.downHeapify(left)
			}
		case h.posExists(right):
			if h.lessFn(h.data[right], h.data[pos]) {
				h.swap(right, pos)
				h.downHeapify(right)
			}
		}
	case MaxHeap:
		switch {
		case h.posExists(left) && h.posExists(right):
			// if both children present, we need to choose
			// one with smallest value in this case
			swapWith := left
			if h.lessFn(h.data[left], h.data[right]) {
				swapWith = right
			}
			// check if we need to swap values
			if h.lessFn(h.data[pos], h.data[swapWith]) {
				h.swap(swapWith, pos)
				h.downHeapify(swapWith)
			}
		case h.posExists(left):
			if h.lessFn(h.data[pos], h.data[left]) {
				h.swap(left, pos)
				h.downHeapify(left)
			}
		case h.posExists(right):
			if h.lessFn(h.data[pos], h.data[right]) {
				h.swap(right, pos)
				h.downHeapify(right)
			}
		}
	}
}

// parent node position in slice relative to pos
func (h *Heap) parent(pos int) int { return (pos - 1) / 2 }

// leftChild location relative to pos
func (h *Heap) leftChild(pos int) int { return 2*pos + 1 }

// rightChild location relative to pos
func (h *Heap) rightChild(pos int) int { return 2*pos + 2 }

// posExists check if position in heap is existing
func (h *Heap) posExists(pos int) bool { return h.Len()-1 >= pos }

// numOfLeaves in heap
func (h *Heap) numOfLeaves() int {
	return int(math.Ceil(float64(h.Len()) / 2.0))
}

// swap two position in heap
func (h *Heap) swap(pos1, pos2 int) {
	h.data[pos1], h.data[pos2] = h.data[pos2], h.data[pos1]
}
