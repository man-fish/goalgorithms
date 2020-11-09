/*
Package queue implement a queue datasruct:
	In computer science, a **queue** is a particular kind of abstract data
	type or collection in which the entities in the collection are
	kept in order and the principle (or only) operations on the
	collection are the addition of entities to the rear terminal
	position, known as enqueue, and removal of entities from the
	front terminal position, known as dequeue. This makes the queue
	a First-In-First-Out (FIFO) data structure. In a FIFO data
	structure, the first element added to the queue will be the
	first one to be removed. This is equivalent to the requirement
	that once a new element is added, all elements that were added
	before have to be removed before the new element can be removed.
	Often a peek or front operation is also entered, returning the
	value of the front element without dequeuing it. A queue is an
	example of a linear data structure, or more abstractly a
	sequential collection.
WikiPage:
	* https://en.wikipedia.org/wiki/Queue_(abstract_data_type)
*/
package queue

import "github.com/man-fish/goalgorithms/datastructures/doublylinkedlist"

// Queue represented a queue implement by doublylinkedlist
type Queue struct {
	list *doublylinkedlist.DoublyLinkedList
}

// New is the constructor of Queue
func New() *Queue {
	return &Queue{
		list: doublylinkedlist.New(),
	}
}

// IsEmpty return whether the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}

// Size return ele nums in queue
func (q *Queue) Size() int {
	return q.list.Len()
}

// Add a element to the front of queue
func (q *Queue) Add(v interface{}) {
	q.list.PushBack(v)
}

// Poll remove and return the first ele from queue
func (q *Queue) Poll() interface{} {
	return q.list.RemoveFront()
}

// Peek return the first ele from queue
func (q *Queue) Peek() interface{} {
	return q.list.Front().Value
}
