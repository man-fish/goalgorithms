package priorityqueue

import (
	"testing"

	"github.com/man-fish/goalgorithms/datastructures/compare"
)

type cpInt int

func (i cpInt) Equal(c compare.Comparable) bool {
	return i == c.(cpInt)
}

func (i cpInt) CompareTo(c compare.Comparable) int {
	if i-c.(cpInt) > 0 {
		return 1
	} else if i-c.(cpInt) == 0 {
		return 0
	} else {
		return -1
	}
}

var _ compare.Comparable = (*cpInt)(nil)

var cpInts = []cpInt{2, 3, 5, 1, 7}

func TestAdd(t *testing.T) {
	pq := New(10)
	for _, v := range cpInts {
		pq.Add(v)
	}
	if pq.Size() != len(cpInts) {
		t.Errorf("add nums failed expected: %v", len(cpInts))
		t.Errorf("					getted: %v", pq.Size())
	}
}

func TestTop(t *testing.T) {
	pq := New(10)
	for _, v := range cpInts {
		pq.Add(v)
	}
	if pq.Top() != cpInt(7) {
		t.Errorf("get max failed expected: %v", 7)
		t.Errorf("				   getted: %v", pq.Top())
	}
}

func TestPop(t *testing.T) {
	pq := New(10)
	for _, v := range cpInts {
		pq.Add(v)
	}
	t.Logf(" init: %v", pq)
	if m := pq.Pop(); m != cpInt(7) {
		t.Errorf("pop max failed expected: %v", 7)
		t.Errorf("				   getted: %v", m)
	}
	t.Logf("fst pop: %v", pq)
	pq.Pop()
	t.Logf("sec pop: %v", pq)
	if m := pq.Pop(); m != cpInt(3) {
		t.Errorf("pop max failed expected: %v", 3)
		t.Errorf("				   getted: %v", m)
	}
}
