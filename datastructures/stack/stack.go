/*
Package stack implement a stack datasruct:
In computer science, a stack is an abstract data type that serves as a collection of elements, with two main principal operations:
	* push: which adds an element to the collection.
	* pop : which removes the most recently added element that was not yet removed.
WikiPage:
	* https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
*/
package stack

import "github.com/man-fish/goalgorithms/datastructures/doublylinkedlist"

// Stack represented a stack implement by doublylinkedlist
type Stack struct {
	list *doublylinkedlist.DoublyLinkedList
}

// New is the constructor of Stack
func New() *Stack {
	return &Stack{
		list: doublylinkedlist.New(),
	}
}

// IsEmpty return whether the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

// Size return ele nums in stack
func (s *Stack) Size() int {
	return s.list.Len()
}

// Push add a element at the top of the stack
func (s *Stack) Push(v interface{}) {
	s.list.PushBack(v)
}

// Pop remove and return the top ele from stack
func (s *Stack) Pop() interface{} {
	return s.list.RemoveBack()
}

// Peek return the top ele from stack
func (s *Stack) Peek() interface{} {
	return s.list.Back().Value
}
