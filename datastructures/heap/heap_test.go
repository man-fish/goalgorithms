package heap

import (
	"testing"
)

var ints = []int{2, -3, 5, 4, 7, 1}

func TestMaxHeap(t *testing.T) {
	h := New(ints, false)
	if !h.isHeapified() {
		t.Errorf("not a heap:\n%v", h)
	}
}

func TestInsert(t *testing.T) {
	h := New(ints, false)
	h.Insert(10)
	if !h.isHeapified() {
		t.Errorf("not a heap after insert:\n%v", h)
	}
	if h.size != len(ints)+1 {
		t.Errorf("err size after insert:\n%v", h)
	}
}
