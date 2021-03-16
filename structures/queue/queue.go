//
// Package queue implements queue data structure.
//
// There is two implementations:
//   * Queue based on circular buffer
//   * LLQueue based on linked list
//
package queue

import (
	"errors"
)

// https://en.wikipedia.org/wiki/Queue_(abstract_data_type)
// https://afteracademy.com/blog/queue-and-its-basic-operations

// Queue structure backed by slice (circular array)
type Queue struct {
	data []interface{}
	head int
	tail int
}

// NewQueue creates new queue of given size
func NewQueue(size int) *Queue {
	return &Queue{
		data: make([]interface{}, size),
		head: -1, // empty list
		tail: -1, // start inserting at position 0
	}
}

// Empty return true if queue is empty
func (q *Queue) Empty() bool {
	return q.tail == -1 && q.head == -1
}

// Size of active elements in queue
func (q *Queue) Size() int {
	if q.Empty() {
		return 0
	}
	return (q.tail + 1) - (q.head % cap(q.data))
}

// Enqueue adds new element to queue
func (q *Queue) Enqueue(v interface{}) error {
	newTail := (q.tail + 1) % cap(q.data)
	if q.data[newTail] != nil {
		return errors.New("queue is full")
	}
	if q.head == -1 {
		q.head = newTail
	}
	q.tail = newTail
	q.data[newTail] = v
	return nil
}

// Dequeue removes and returns oldest element in queue
func (q *Queue) Dequeue() interface{} {
	if q.Empty() {
		return nil
	}
	defer func() {
		if q.tail == q.head {
			q.tail, q.head = -1, -1
		} else {
			q.head = (q.head + 1) % cap(q.data)
		}
	}()
	data := q.data[q.head]
	q.data[q.head] = nil
	return data
}

// Front returns oldest element from queue without removing it
func (q *Queue) Front() interface{} {
	return q.data[q.head]
}

// LLQueue structure backed by linked list
type LLQueue struct {
	cap  int
	len  int
	head *LLQueueNode
	tail *LLQueueNode
}

// LLQueueNode of LLQueue structure
type LLQueueNode struct {
	data interface{}
	next *LLQueueNode
}

// NewLLQueue creates new queue of given size
func NewLLQueue(size int) *LLQueue {
	return &LLQueue{cap: size}
}

// Empty return true if queue is empty
func (q *LLQueue) Empty() bool {
	return q.head == nil && q.tail == nil
}

// Size of active elements in queue
func (q *LLQueue) Size() int {
	var cnt int
	for i := q.head; i != nil; i = i.next {
		cnt++
	}
	return cnt
}

// Enqueue adds new element to queue
func (q *LLQueue) Enqueue(v interface{}) error {
	if q.len+1 > q.cap {
		return errors.New("queue is full")
	}
	n := &LLQueueNode{data: v}
	if q.tail != nil {
		q.tail.next = n
	}
	q.tail = n
	if q.head == nil {
		q.head = q.tail
	}
	q.len++
	return nil
}

// Dequeue removes and returns oldest element in queue
func (q *LLQueue) Dequeue() interface{} {
	if q.head == nil {
		return nil
	}
	defer func() {
		q.head = q.head.next
		if q.head == nil {
			q.tail = nil
		}
	}()
	q.len--
	return q.head.data
}

// Front returns oldest element from queue without removing it
func (q *LLQueue) Front() interface{} {
	if q.head == nil {
		return nil
	}
	return q.head.data
}
