//
// Package stack implements stack data structure.
//
// There is two implementations:
//   * Stack based on slice
//   * LLStack based on linked list
//
// https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
// https://afteracademy.com/blog/stack-and-its-basic-operations
//
package stack

import "errors"

// Stack is stack implementation using slices
type Stack struct {
	data []interface{}
	head int
}

// NewStack creates a stack of given size
func NewStack(size int) *Stack {
	return &Stack{
		data: make([]interface{}, size),
		head: -1, // means stack is empty
	}
}

// Empty returns true if we do not have any data in stack
func (s *Stack) Empty() bool {
	return s.head == -1
}

// Size returns number of available elements in stack
func (s *Stack) Size() int {
	return s.head + 1
}

// Push element in stack
func (s *Stack) Push(i interface{}) error {
	if s.head+1 > cap(s.data)-1 {
		return errors.New("stack overflow")
	}
	s.data[s.head+1] = i
	s.head += 1
	return nil
}

// Peek an element from top of the stack
func (s *Stack) Peek() interface{} {
	if s.head == -1 {
		return nil
	}
	return s.data[s.head]
}

// Pop an element from the top of the stack
// This removes element from stack
func (s *Stack) Pop() interface{} {
	if s.head == -1 {
		return nil
	}
	defer func() {
		s.data[s.head] = nil
		s.head -= 1
	}()
	return s.data[s.head]
}

// LLStack is stack implementation using linked list
type LLStack struct {
	head *LLStackNode
}

type LLStackNode struct {
	data interface{}
	next *LLStackNode
}

// NewLLStack creates a new linked list stack
func NewLLStack() *LLStack {
	return &LLStack{}
}

// Empty returns true if we do not have any data in stack
func (s *LLStack) Empty() bool {
	return s.head == nil
}

// Size returns number of available elements in stack
func (s *LLStack) Size() int {
	var cnt int
	for i := s.head; i != nil; i = i.next {
		cnt++
	}
	return cnt
}

// Push element in stack
func (s *LLStack) Push(i interface{}) {
	s.head = &LLStackNode{
		data: i,
		next: s.head,
	}
}

// Peek an element from top of the stack
func (s *LLStack) Peek() interface{} {
	return s.head.data
}

// Pop an element from the top of the stack
// This removes element from stack
func (s *LLStack) Pop() interface{} {
	defer func() {
		s.head = s.head.next
	}()
	return s.head.data
}
