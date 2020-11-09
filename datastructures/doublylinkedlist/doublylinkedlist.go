/*
Package doublylinkedlist implements a doubly linked list:
	* this package is just copied from container/list in standard library.
	* it avoid one list access the other one on insert or other operate.

	To iterate over a list (where l is a *List):
	```go
		for e := l.Front(); e != nil; e = e.Next() {
			// do something with e.Value
		}
	```
Doubly LinkedList on Wiki:
	In computer science, a doubly linked list is a linked data structure
	that consists of a set of sequentially linked records called nodes.
	Each node contains three fields: two link fields (references to the
	previous and to the next node in the sequence of nodes) and one data
	field. The beginning and ending nodes' previous and next links.
*/
package doublylinkedlist

// Element is an element of linked list
type Element struct {
	next, prev *Element
	// The list to which this element belongs.
	list  *DoublyLinkedList
	Value interface{}
}

// Next returns the next list element or nil
func (e *Element) Next() *Element {
	if p := e.next; e.list != nil && p != &e.list.root {
		// avoid single node
		return p
	}
	return nil
}

// Prev returns the prev list element or nil
func (e *Element) Prev() *Element {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// DoublyLinkedList represents a doubly linked list.
type DoublyLinkedList struct {
	root Element // root is just a head node contains no value
	len  int     // element nums without root
}

// Init initializes or clear list l
func (l *DoublyLinkedList) Init() *DoublyLinkedList {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

// lazyInit lazily initializes a zero List value.
func (l *DoublyLinkedList) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

// New returns an initialized list.
func New() *DoublyLinkedList {
	return new(DoublyLinkedList).Init()
}

// Len returns len of l
func (l *DoublyLinkedList) Len() int {
	return l.len
}

// Front returns the first element of list l or nil if the list is empty.
func (l *DoublyLinkedList) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// Back returns the last element of list l or nil if the list is empty.
func (l *DoublyLinkedList) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

func (l *DoublyLinkedList) insert(e, at *Element) *Element {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (l *DoublyLinkedList) insertValue(v interface{}, at *Element) *Element {
	return l.insert(&Element{Value: v}, at)
}

// avoid memory leaks
func (l *DoublyLinkedList) remove(e *Element) *Element {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil // avoid memory leaks
	e.prev = nil // avoid memory leaks
	e.list = nil
	l.len--
	return e
}

// move moves e to next to at and returns e.
func (l *DoublyLinkedList) move(e, at *Element) *Element {
	if e == at {
		return e
	}
	// remove connect of e from its origin space
	e.prev.next = e.next
	e.next.prev = e.prev

	e.prev = at
	e.next = at.next

	e.prev.next = e
	e.next.prev = e

	return e
}

// Remove removes e from l if e is an element of list l.
// It returns the element value e.Value.
// The element must not be nil.
func (l *DoublyLinkedList) Remove(e, at *Element) interface{} {
	if e.list == l {
		// if e.list == l, l must have been initialized when e was inserted
		// in l or l == nil (e is a zero Element) and l.remove will crash
		l.remove(e)
	}
	return e.Value
}

// RemoveFront removes the first ele from list
// It returns the element value e.Value.
func (l *DoublyLinkedList) RemoveFront() interface{} {
	return l.remove(l.root.next).Value
}

// RemoveBack removes the first ele from list
// It returns the element value e.Value.
func (l *DoublyLinkedList) RemoveBack() interface{} {
	return l.remove(l.root.prev).Value
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (l *DoublyLinkedList) PushFront(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *DoublyLinkedList) PushBack(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}

// InsertBefore inserts a new element e with value v immediately before mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *DoublyLinkedList) InsertBefore(v interface{}, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark.prev)
}

// InsertAfter inserts a new element e with value v immediately after mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *DoublyLinkedList) InsertAfter(v interface{}, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	// see comment in List.Remove about initialization of l
	return l.insertValue(v, mark)
}

// MoveToFront moves element e to the front of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *DoublyLinkedList) MoveToFront(e *Element) {
	if e.list != l || l.root.next == e {
		return
	}
	l.move(e, &l.root)
}

// MoveToBack moves element e to the back of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *DoublyLinkedList) MoveToBack(e *Element) {
	if e.list != l || l.root.prev == e {
		return
	}
	// see comment in List.Remove about initialization of l
	l.move(e, l.root.prev)
}
