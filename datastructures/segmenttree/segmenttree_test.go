package segmenttree

import "testing"

var ints []int = []int{1, 3, 5, 7, 9, 11}

func TestSum(t *testing.T) {
	tree := New(ints)
	sum := tree.Sum(0, 0, len(ints)-1, 1, 3)
	if sum != 15 {
		t.Errorf("wanted sum: %v", 15)
		t.Errorf("getted sum: %v", sum)
	}
}

func TestUpdate(t *testing.T) {
	tree := New(ints)
	tree.Update(0, 0, len(ints)-1, 4, 6)
	if tree.arr[4] != 6 {
		t.Errorf("wanted idx4: %v", 6)
		t.Errorf("getted idx4: %v", tree.arr[4])
	}
	sum := tree.Sum(0, 0, len(ints)-1, 1, 4)
	if sum != 21 {
		t.Errorf("wanted sum: %v", 21)
		t.Errorf("getted sum: %v", sum)
	}
}
