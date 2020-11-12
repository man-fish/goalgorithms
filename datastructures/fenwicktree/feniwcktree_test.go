package fenwicktree

import "testing"

var ints = []int{1, 2, 3, 4, 5, 6, 7, 8}
var fenwicktree = []int{0, 1, 3, 3, 10, 5, 11, 7, 36}

func TestIncrease(t *testing.T) {
	tree := New(8)
	for i, v := range ints {
		tree.Increase(i+1, v)
	}
	for i, v := range tree.tree {
		if v != fenwicktree[i] {
			t.Errorf("wanted fenwick[%v]: %v", i, fenwicktree[i])
			t.Errorf("getted fenwick[%v]: %v", i, v)
		}
	}
}

func TestQuery(t *testing.T) {
	tree := New(8)
	for i, v := range ints {
		tree.Increase(i+1, v)
	}
	if sum := tree.Query(8); sum != 36 {
		t.Errorf("wanted [1-8]: %v", 36)
		t.Errorf("getted [1-8]: %v", sum)
	}
}
