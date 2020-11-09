/*
Package linkedlist implements a doubly linked list:
	In computer science, a **linked list** is a linear collection
	of data elements, in which linear order is not given by
	their physical placement in memory. Instead, each
	element points to the next. It is a data structure
	consisting of a group of nodes which together represent
	a sequence. Under the simplest form, each node is
	composed of data and a reference (in other words,
	a link) to the next node in the sequence. This structure
	allows for efficient insertion or removal of elements
	from any position in the sequence during iteration.
	More complex variants add additional links, allowing
	efficient insertion or removal from arbitrary element
	references. A drawback of linked lists is that access
	time is linear (and difficult to pipeline). Faster
	access, such as random access, is not feasible. Arrays
	have better cache locality as compared to linked lists.
WikiPage:
	* https://en.wikipedia.org/wiki/Linked_list
*/
package linkedlist

import (
	"reflect"

	"github.com/man-fish/goalgorithms/datastructures/compare"
)

// Element is an element of linked list
type Element struct {
	Next  *Element
	Value interface{}
}

// LinkedList represents a linked list.
type LinkedList struct {
	root Element
	len  int
}

// New is the constructor of a linked list
func New() *LinkedList {
	return new(LinkedList)
}

// Get indexes a element from linked list
// if i is larger than the list size it
// will return nil
func (l *LinkedList) Get(i int) interface{} {
	p := l.root.Next
	for j := 0; p != nil && j < i; j++ {
		p = p.Next
	}
	if i >= 0 && p != nil {
		return p.Value
	}
	return nil
}

// Set indexes a ele and set its value
func (l *LinkedList) Set(i int, v interface{}) {
	p := l.root.Next
	for j := 0; p != nil && j < i; j++ {
		p = p.Next
	}
	if p != nil {
		p.Value = v
	}
}

// Insert add a ele after the ele on index i
func (l *LinkedList) Insert(i int, v interface{}) *Element {
	p := &l.root
	for j := 0; p != nil && j < i; j++ {
		p = p.Next
	}
	p.Next = &Element{p.Next, v}
	return p.Next
}

// InsertDifferent add a unique value to the list
func (l *LinkedList) InsertDifferent(v interface{}) {
	p := &l.root
	for p.Next != nil {
		p = p.Next
		if p.Value == v {
			return
		}
	}
	p.Next = &Element{p.Next, v}
}

// InsertAtLast add a value at the last of the linkedlist
func (l *LinkedList) InsertAtLast(v interface{}) *Element {
	p := &l.root
	for p.Next != nil {
		p = p.Next
	}
	p.Next = &Element{nil, v}
	return p.Next
}

// Delete the ele on index i
func (l *LinkedList) Delete(i int) *Element {
	p := &l.root
	for j := 0; p != nil && j < i; j++ {
		p = p.Next
	}
	tmp := p.Next
	p.Next = p.Next.Next
	return tmp
}

// Remove remove an element from list
func (l *LinkedList) Remove(x compare.Comparable) *Element {
	p := &l.root
	for p.Next != nil {
		if x.Equal(p.Next.Value.(compare.Comparable)) {
			temp := p.Next
			p.Next = p.Next.Next
			return temp
		}
	}
	return nil
}

// Search search an element from list
func (l *LinkedList) Search(x compare.Comparable) *Element {
	p := l.root.Next
	for p != nil {
		d := p.Value
		if reflect.TypeOf(d).Kind() == reflect.Ptr {
			d = &(p.Value)
		}
		if x.Equal(d.(compare.Comparable)) {
			return p
		}
		p = p.Next
	}
	return nil
}
