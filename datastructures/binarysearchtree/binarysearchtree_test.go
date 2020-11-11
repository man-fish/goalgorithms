package binarysearchtree

import (
	"testing"
)

var ints = []int{67, 87, 43, 23, 45, 37, 90}

/*
		67
	43		87
23		45
	37			90
*/

func TestInsert(t *testing.T) {
	bst := New(ints)
	bst.Insert(45)
	if !bst.Search(45) {
		t.Errorf("wanted 45 but nil")
	}
}

func TestDelete(t *testing.T) {
	bst := New(ints)
	bst.Delete(43)
	if bst.Search(43) {
		t.Errorf("doesn`t want 43 but has.")
	}
	bst.Delete(90)
	if bst.Search(90) {
		t.Errorf("doesn`t want 90 but has.")
	}
}
