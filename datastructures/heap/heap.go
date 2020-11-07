package heap

import (
	"fmt"
	"math"
)

// Heap is a binary tree implemented with an array
type Heap struct {
	tree []int
	size int
	// isMaxium determines whether the heap is a maximum heap or a minimum heap
	isMaximum bool
}

// New is the heap constructor
// isMaximum determines the heap direction.
// true for maximum, false for minimum.
func New(tree []int, isMaximum bool) *Heap {
	h := &Heap{
		tree:      tree,
		size:      len(tree),
		isMaximum: isMaximum,
	}
	h.buildHeap()
	return h
}

func (h *Heap) buildHeap() {
	// get the parent of the last leaf node
	lst := h.size - 1
	plst := (lst - 1) / 2
	// heapify down each node
	for i := plst; i >= 0; i-- {
		h.heapifyDown(i, h.size)
	}
}

func (h *Heap) isHeapified() bool {
	for i := 0; i < h.size; i++ {
		l := 2*i + 1
		r := 2*i + 2
		if l < h.size && ((h.tree[l] < h.tree[i]) && true) != h.isMaximum {
			return false
		}
		if r < h.size && ((h.tree[r] < h.tree[i]) && true) != h.isMaximum {
			return false
		}
	}
	return true
}

func (h *Heap) heapifyDown(i, n int) {
	// get left and right node
	l := 2*i + 1
	r := 2*i + 2
	max := i

	// get the should swap one among l, r and i
	if l < n && ((h.tree[l] < h.tree[max]) && true) != h.isMaximum {
		max = l
	}
	if r < n && ((h.tree[r] < h.tree[max]) && true) != h.isMaximum {
		max = r
	}

	if max != i {
		h.tree[max], h.tree[i] = h.tree[i], h.tree[max]
		// heapify down
		h.heapifyDown(max, n)
	}
}

func (h *Heap) heapifyUp(i int) {
	// get parent
	p := (i - 1) / 2
	if p > 0 && ((h.tree[i] < h.tree[p]) && true) != h.isMaximum {
		h.tree[p], h.tree[i] = h.tree[i], h.tree[p]
		// heapify up
		h.heapifyUp(p)
	}
}

// Size returns size of heap
func (h *Heap) Size() int {
	return h.size
}

// Insert add node to heap and keep the heap heapified
func (h *Heap) Insert(node int) {
	h.tree = append(h.tree, node)
	h.size++
	h.heapifyUp(h.size - 1)
}

func (h *Heap) String() string {
	heap := "\n"
	idx := 0
	for idx < h.size {
		n := int(math.Pow(2, float64(idx)))
		for i := 0; i < n && (i+idx) < h.size; i++ {
			heap += fmt.Sprintf("%v\t", h.tree[idx+i])
		}
		idx += n
		heap += "\n"
	}
	return heap
}

// Sort is a sort Algorithm
func Sort(arr []int) []int {
	h := New(arr, true)
	for i := h.Size() - 1; i > 0; i-- {
		h.heapifyDown(0, i)
		h.tree[i], h.tree[0] = h.tree[0], h.tree[i]
	}
	return h.tree
}
