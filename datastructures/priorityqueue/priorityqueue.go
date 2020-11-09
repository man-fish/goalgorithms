/*
Package priorityqueue implement a priority queue heap datasruct:
	In computer science, a **priority queue** is an abstract data type
	which is like a regular queue or stack data structure, but where
	additionally each element has a "priority" associated with it.
	In a priority queue, an element with high priority is served before
	an element with low priority. If two elements have the same
	priority, they are served according to their order in the queue.

	While priority queues are often implemented with heaps, they are
	conceptually distinct from heaps. A priority queue is an abstract
	concept like "a list" or "a map"; just as a list can be implemented
	with a linked list or an array, a priority queue can be implemented
	with a heap or a variety of other methods such as an unordered
	array.
WikiPage:
	* https://en.wikipedia.org/wiki/Priority_queue)
*/
package priorityqueue

import "github.com/man-fish/goalgorithms/datastructures/compare"

// PqHeap is a priority queue
type PqHeap struct {
	tree []compare.Comparable
	// element in tree
	size int
}

// New is a constructor
func New(n int) *PqHeap {
	return &PqHeap{
		tree: make([]compare.Comparable, n+1),
	}
}

// Size return q.size
func (q *PqHeap) Size() int {
	return q.size
}

// IsEmpty returns whether the tree is empty
func (q *PqHeap) IsEmpty() bool {
	return q.size == 0
}

// Top returns the top ele of the queue
func (q *PqHeap) Top() compare.Comparable {
	return q.tree[1]
}

// Add add a ele to queue
func (q *PqHeap) Add(c compare.Comparable) {
	q.size++
	q.tree[q.size] = c
	q.swim(q.size)
}

// Pop removes the top ele from queue
func (q *PqHeap) Pop() compare.Comparable {
	top := q.tree[1]
	q.tree[1], q.tree[q.size] = q.tree[q.size], q.tree[1]
	q.size--
	q.sink(1)
	q.tree[q.size+1] = nil
	return top
}

func (q *PqHeap) swim(i int) {
	for i > 1 && q.tree[i].CompareTo(q.tree[i/2]) == 1 {
		q.tree[i], q.tree[i/2] = q.tree[i/2], q.tree[i]
		i = i / 2
	}
}

func (q *PqHeap) sink(i int) {
	for i*2 <= q.size {
		k := i * 2
		if q.tree[k].CompareTo(q.tree[k+1]) == -1 {
			k++
		}
		if q.tree[k].CompareTo(q.tree[i]) == -1 {
			break
		}
		q.tree[i], q.tree[k] = q.tree[k], q.tree[i]
		i = k
	}
}
